package auth

import "ohdada/g2gserver/internal/pkg/pb"

var (
	// DefaultAuthenticatorSimpleFactory .
	DefaultAuthenticatorSimpleFactory = AuthenticatorSimpleFactory{}
)

// Result .
type Result struct {
	PlayerID string
}

// Authenticator .
type Authenticator interface {
	Type() pb.AuthType
	Verify(token string) (Result, error)
	GenToken(userID string) (string, error)
}
