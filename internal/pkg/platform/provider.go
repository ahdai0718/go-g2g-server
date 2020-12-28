package platform

import (
	"ohdada/g2gserver/internal/pkg/pb"
	"net/http"
	"time"
)

// Provider .
type Provider interface {
	Init(platformProvider *pb.PlatformProvider) error

	SetPublicIPAddress(publicIPAddress string)
	SetRunMode(runMode string)
	SetLanguage(language string)
	SetPlayerAccessToken(token string)
	SetCurrencyCode(code string)

	ProviderName() string
	PlatformPlayer() *pb.PlatformPlayer
	AESKey() string
	AESIV() string

	Tick()

	Auth(request *http.Request) (*pb.PlatformPlayer, error)
	RequestPlayerBalance(*pb.Transaction) (*pb.PlatformPlayer, error)
	LockPlayerTransaction(*pb.Transaction) (*pb.PlatformPlayer, error)
	UnlockPlayerTransaction(*pb.Transaction) (*pb.PlatformPlayer, error)
	CancelPlayerTransaction(*pb.Transaction) (*pb.PlatformPlayer, error)
	PlacePlayerBet(*pb.Transaction) (*pb.PlatformPlayer, error)
	CancelPlayerBet(*pb.Transaction) (*pb.PlatformPlayer, error)
	SettlePlayerBet(*pb.Transaction) (*pb.PlatformPlayer, error)
	IsPlayerLock() (bool, error)

	EncryptAES(data []byte) ([]byte, error)
	DecryptAES(data []byte) ([]byte, error)
}

// ClientAccessToken .
type ClientAccessToken struct {
	Token       string
	LastUpdated time.Time
	Timeout     time.Duration
}
