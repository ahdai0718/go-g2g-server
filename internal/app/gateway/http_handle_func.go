package gateway

import (
	"net/http"
	"ohdada/g2gserver/internal/pkg/auth"
	"ohdada/g2gserver/internal/pkg/pb"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	authenticator = auth.DefaultAuthenticatorSimpleFactory.Create(&pb.Auth{
		Type:   pb.AuthType_AT_JWT,
		Secret: "12345678",
	})
)

// AllowIntranetOnly .
func AllowIntranetOnly(context *gin.Context) {
	isAllow := false

	if strings.Contains(context.Request.RemoteAddr, "192.168.") {
		isAllow = true
	}

	if strings.Contains(context.Request.RemoteAddr, "127.0.0.1") {
		isAllow = true
	}

	if strings.Contains(context.Request.RemoteAddr, "0.0.0.0") {
		isAllow = true
	}

	if !isAllow {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}

// HandleRequestServerInfo .
func HandleRequestServerInfo(context *gin.Context) {
	context.ProtoBuf(http.StatusOK, ServerInfo())
}

// HandleGetUserTransactions .
func HandleGetUserTransactions(context *gin.Context) {

}

// CheckOAuthToken .
func CheckOAuthToken(context *gin.Context) {
	if authenticator.Type() == pb.AuthType_AT_OAUTH {
		authorization := context.GetHeader("Authorization")
		strList := strings.Split(authorization, "Bearer")
		if len(strList) != 2 {
			context.JSON(400, gin.H{
				"error": gin.H{
					"code":    -1,
					"message": "Authorization: not bearer.",
				},
			})
			context.Abort()
		} else {
			accessToken := strings.Trim(strList[1], " ")
			if accessToken != "12345678" {
				context.JSON(400, gin.H{
					"error": gin.H{
						"code":    -1,
						"message": "Authorization: access token invalid.",
					},
				})
				context.Abort()
			}
		}
	}
}

// CheckAccessToken .
func CheckAccessToken(context *gin.Context) {

	accessToken := context.GetHeader("X-Access-Token")

	if len(strings.Trim(accessToken, " ")) == 0 {
		context.JSON(400, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": "Access token is required.",
			},
		})
	}
}

// HandleOAuthAccessToken .
func HandleOAuthAccessToken(context *gin.Context) {
	clientID, _ := context.GetPostForm("client_id")
	clientSecret, _ := context.GetPostForm("client_secret")
	grantType, _ := context.GetPostForm("grant_type")
	scope, _ := context.GetPostForm("scope")

	if clientID == "default" &&
		clientSecret == "12345678" &&
		grantType == "client_credentials" &&
		scope == "bet" {
		context.JSON(200, gin.H{
			"access_token": "12345678",
			"expires_in":   time.Now().Add(time.Hour).Unix(),
			"token_type":   "Bearer",
			"scope":        "bet",
		})
	} else {
		context.JSON(401, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": "OAuth failed.",
			},
		})
	}
}

// HandlePlayerTokenValidate .
func HandlePlayerTokenValidate(context *gin.Context) {

	accessToken := context.GetHeader("X-Access-Token")

	authResult, err := authenticator.Verify(accessToken)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	context.JSON(200, gin.H{
		"id":            authResult.PlayerID,
		"name":          authResult.PlayerID,
		"language":      "zh-CN",
		"balance":       1e6,
		"currency_code": "TWD",
		"user_sn":       0,
		"icon":          "01",
	})
}
