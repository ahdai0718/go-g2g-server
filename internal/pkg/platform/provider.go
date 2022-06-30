package platform

import (
	"net/http"
	"ohdada/g2gserver/internal/pkg/pb"
	"time"
)

// Provider .
type Provider interface {
	Init(platformProvider *pb.PlatformProvider) error

	SetPublicIPAddress(publicIPAddress string)
	SetRunMode(runMode string)
	SetLanguage(language string)
	SetPlayerAccessToken(token string)

	ProviderName() string
	PlatformPlayer() *pb.PlatformPlayer
	AESKey() string
	AESIV() string

	Tick()

	Auth(request *http.Request) (*pb.PlatformPlayer, error)

	EncryptAES(data []byte) ([]byte, error)
	DecryptAES(data []byte) ([]byte, error)
}

// ClientAccessToken .
type ClientAccessToken struct {
	LastUpdated time.Time
	Token       string
	Timeout     time.Duration
}
