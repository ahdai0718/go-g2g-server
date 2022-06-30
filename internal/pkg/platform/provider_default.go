package platform

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"ohdada/g2gserver/internal/pkg/constant"
	"ohdada/g2gserver/internal/pkg/encrypt/aes"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
	"path"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

const (
	defaultAESKeyByteLength   = 16
	defaultAESIVByteLength    = 16
	defaultAccessTokenTimeout = 60 * time.Minute
)

var (
	defaultClientAccessTokenMap = make(map[string]*ClientAccessToken)
	defaultAPIPath              = pb.PlatformProviderAPIPath{
		OauthAccessToken:        "oauth/access_token",
		PlayerTokenValidate:     "player/token/validate",
		PlayerBalance:           "player/balance",
		PlayerBetPlace:          "player/bet/place",
		PlayerBetCancel:         "player/bet/cancel",
		PlayerBetSettle:         "player/bet/settle",
		PlayerTransactionLock:   "player/transaction/lock",
		PlayerTransactionUnlock: "player/transaction/unlock",
		PlayerTransactionCancel: "player/transaction/cancel",
		PlayerTransactionStatus: "player/transaction/status",
	}
)

// OAuthAccessTokenResponseDefault .
type OAuthAccessTokenResponseDefault struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	Error       struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
	ExpiresIn int64 `json:"expires_in"`
}

// ProviderDefault .
type ProviderDefault struct {
	lastRequestLockTime time.Time
	player              *pb.PlatformPlayer
	client              *resty.Client
	pb.PlatformProvider
	isRequestLock   bool
	isRequestUnlock bool
	isRequestBet    bool
	isRequestSettle bool
}

// Init .
func (provider *ProviderDefault) Init(platformProvider *pb.PlatformProvider) error {
	copier.Copy(provider, platformProvider)
	provider.player = &pb.PlatformPlayer{}

	provider.client = resty.New()

	if provider.shouldUseClientAccessToken() {
		provider.requestOAuthAccessToken()
	}

	return nil
}

// SetPublicIPAddress .
func (provider *ProviderDefault) SetPublicIPAddress(publicIPAddress string) {
	provider.PublicIpAddress = publicIPAddress
}

// SetRunMode .
func (provider *ProviderDefault) SetRunMode(runMode string) {
	provider.RunMode = runMode
}

// SetLanguage .
func (provider *ProviderDefault) SetLanguage(language string) {
	provider.player.Language = language
}

// ProviderName .
func (provider *ProviderDefault) ProviderName() string {
	return provider.PlatformProvider.Name
}

// PlatformPlayer .
func (provider *ProviderDefault) PlatformPlayer() *pb.PlatformPlayer {
	return provider.player
}

// SetPlayerAccessToken .
func (provider *ProviderDefault) SetPlayerAccessToken(token string) {
	provider.player.AccessToken = token
}

// SetLockCredit .
func (provider *ProviderDefault) SetLockCredit(credit int64) {
	provider.player.LockCredit = credit
}

// LockCredit .
func (provider *ProviderDefault) LockCredit() int64 {
	return provider.player.LockCredit
}

// AESKey .
func (provider *ProviderDefault) AESKey() string {
	return provider.PlatformProvider.AesKey
}

// AESIV .
func (provider *ProviderDefault) AESIV() string {
	return provider.PlatformProvider.AesIv
}

// Tick .
func (provider *ProviderDefault) Tick() {

	if err := provider.checkThenRenewAccessToken(); err != nil {
		glog.Error(err)
	}

}

// EncryptAES .
func (provider *ProviderDefault) EncryptAES(data []byte) ([]byte, error) {
	return aes.Encrypt(data, []byte(provider.AESKey()), []byte(provider.AESIV()))
}

// DecryptAES .
func (provider *ProviderDefault) DecryptAES(data []byte) ([]byte, error) {
	return aes.Decrypt(data, []byte(provider.AESKey()), []byte(provider.AESIV()))
}

// Auth .
func (provider *ProviderDefault) Auth(request *http.Request) (*pb.PlatformPlayer, error) {

	playerAccessToken, err := provider.parsePlayerAccessToken(request)

	if err != nil {
		glog.Error(err)
		return provider.player, err
	}

	provider.player.AccessToken = playerAccessToken

	body := struct {
		IPAddress string `json:"public_ip_address"`
	}{
		IPAddress: provider.PublicIpAddress,
	}

	response, err := provider.request(defaultAPIPath.PlayerTokenValidate, body)

	if err != nil {
		glog.Error(err)
		return provider.player, err
	}

	json := struct {
		UserID      string `json:"id"`
		Language    string `json:"language"`
		DisplayName string `json:"display_name"`
		Icon        string `json:"icon"`
		UserName    string `json:"name"`
		Error       struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		} `json:"error"`
		UserSN  int     `json:"user_sn"`
		Balance float64 `json:"balance"`
	}{}

	err = provider.client.JSONUnmarshal(response, &json)

	if err != nil {
		glog.Errorln("resp.Body:", string(response))
		glog.Errorln(err)
		return provider.player, err
	}

	if json.Error.Code > 0 || len(json.UserID) == 0 {
		glog.Errorln(string(response))
		return provider.player, errors.New(json.Error.Message)
	}

	provider.player.Id = json.UserID
	provider.player.Sn = int64(json.UserSN)
	provider.player.Name = json.UserName
	provider.player.IdAtPlatform = fmt.Sprintf("%s@%s", json.UserID, provider.ProviderName())
	provider.player.DisplayName = json.DisplayName
	provider.player.Balance = json.Balance
	provider.player.Platform = provider.ProviderName()
	provider.player.Icon = json.Icon

	return provider.player, err
}

func (provider *ProviderDefault) request(requestPath string, body interface{}) ([]byte, error) {
	err := provider.checkThenSetClientAuthToken()

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	url, err := url.Parse(provider.ApiUrlBase)

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	url.Path = path.Join(url.Path, requestPath)

	provider.client.SetRetryCount(0)

	req := provider.client.R().
		SetHeader("X-Access-Token", provider.player.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body)

	if provider.RunMode == constant.RunModeDev || constant.RunModeDev == constant.RunModeDebug {
		req.EnableTrace()
	}

	response, err := req.Post(url.String())

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("HttpStatusCode[%d]:[%s]", response.StatusCode(), string(response.Body()))
	}

	if provider.RunMode == constant.RunModeDev || constant.RunModeDev == constant.RunModeDebug {
		glog.Infoln("=============RequestInfo=============")
		glog.Infoln("url:", url)
		glog.Infoln("body:", body)
		glog.Infoln("req.Header:", req.Header)
		glog.Infoln("response.Body:", string(response.Body()))
		glog.Infoln("=====================================")
		provider.showTraceInfo(response.Request.TraceInfo())
	}

	return response.Body(), nil
}

func (provider *ProviderDefault) genAESkey() ([]byte, error) {
	key := make([]byte, defaultAESKeyByteLength)
	_, err := rand.Read(key)
	return key, err
}

func (provider *ProviderDefault) genAESIV() ([]byte, error) {
	iv := make([]byte, defaultAESIVByteLength)
	_, err := rand.Read(iv)
	return iv, err
}

func (provider *ProviderDefault) showTraceInfo(ti resty.TraceInfo) {
	glog.Infoln("========TraceInfo========")
	glog.Infoln("DNSLookup    :", ti.DNSLookup)
	glog.Infoln("ConnTime     :", ti.ConnTime)
	glog.Infoln("TLSHandshake :", ti.TLSHandshake)
	glog.Infoln("ServerTime   :", ti.ServerTime)
	glog.Infoln("ResponseTime :", ti.ResponseTime)
	glog.Infoln("TotalTime    :", ti.TotalTime)
	glog.Infoln("IsConnReused :", ti.IsConnReused)
	glog.Infoln("IsConnWasIdle:", ti.IsConnWasIdle)
	glog.Infoln("ConnIdleTime :", ti.ConnIdleTime)
	glog.Infoln("========TraceInfo========")
}

func (provider *ProviderDefault) parsePlayerAccessToken(request *http.Request) (token string, err error) {

	token = request.Header.Get("X-Access-Token")

	if len(token) > 0 {
		return
	}

	tokenList := request.URL.Query()["token"]

	if len(tokenList) > 0 {
		token = tokenList[0]
		return
	}

	tokenList = request.URL.Query()["access_token"]
	if len(tokenList) > 0 {
		token = tokenList[0]
		return
	}

	err = fmt.Errorf("Player token is not exists")

	return
}

func (provider *ProviderDefault) checkThenSetClientAuthToken() error {

	if provider.shouldUseClientAccessToken() {

		err := provider.checkThenRenewAccessToken()

		if err != nil {
			glog.Error(err)
			return err
		}

		isExists, clientAccessToken := provider.checkClientAccessToken()

		if !isExists {
			return fmt.Errorf("client access token [%s] not exists", provider.ProviderName())
		}

		provider.client.SetAuthToken(clientAccessToken.Token)
	}

	return nil
}

func (provider *ProviderDefault) shouldUseClientAccessToken() bool {
	return provider.PlatformProvider.Auth.Type == pb.AuthType_AT_OAUTH
}

func (provider *ProviderDefault) checkClientAccessToken() (bool, *ClientAccessToken) {
	if clientAccessToken, isExists := defaultClientAccessTokenMap[provider.ProviderName()]; isExists {
		return true, clientAccessToken
	}
	return false, nil
}

func (provider *ProviderDefault) checkThenRenewAccessToken() error {
	if provider.isAccessTokenExpired() {
		return provider.requestOAuthAccessToken()
	}
	return nil
}

func (provider *ProviderDefault) isAccessTokenExpired() bool {

	clientAccessToken, isExists := defaultClientAccessTokenMap[provider.ProviderName()]

	if !isExists {
		return true
	}

	if time.Since(clientAccessToken.LastUpdated) >= clientAccessToken.Timeout {
		return true
	}

	return false
}

func (provider *ProviderDefault) requestOAuthAccessToken() error {

	url, err := url.Parse(provider.ApiUrlBase)

	url.Path = path.Join(url.Path, defaultAPIPath.OauthAccessToken)

	req := provider.client.R().
		SetFormData(map[string]string{
			"client_id":     provider.PlatformProvider.Auth.Id,
			"client_secret": provider.PlatformProvider.Auth.Secret,
			"grant_type":    provider.PlatformProvider.Auth.GrantType,
			"scope":         provider.PlatformProvider.Auth.Scope,
		})

	if provider.RunMode == constant.RunModeDev || constant.RunModeDev == constant.RunModeDebug {
		req.EnableTrace()
	}

	response, err := req.Post(url.String())

	if err != nil {
		glog.Error(err)
		return err
	}

	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("HttpStatusCode[%d]:[%s]", response.StatusCode(), string(response.Body()))
	}

	if provider.RunMode == constant.RunModeDev || constant.RunModeDev == constant.RunModeDebug {
		glog.Infoln("========ValidateToken========")
		provider.showTraceInfo(response.Request.TraceInfo())
	}

	oAuthAccessTokenResponse := &OAuthAccessTokenResponseDefault{}

	err = provider.client.JSONUnmarshal(response.Body(), &oAuthAccessTokenResponse)

	if err != nil {
		glog.Errorln("req.URL:", req.URL)
		glog.Errorln("req.Header:", req.Header)
		glog.Errorln("req.Body:", req.Body)
		glog.Errorln("resp.Body:", string(response.Body()))
		glog.Errorln(err)
		return err
	}

	if oAuthAccessTokenResponse.Error.Code > 0 {
		glog.Error(err, string(response.Body()))
		return errors.New(oAuthAccessTokenResponse.Error.Message)
	}

	defaultClientAccessTokenMap[provider.ProviderName()] = &ClientAccessToken{
		Token:       oAuthAccessTokenResponse.AccessToken,
		LastUpdated: time.Now(),
		Timeout:     defaultAccessTokenTimeout,
	}

	return err
}
