package auth

import "ohdada/g2gserver/internal/pkg/pb"

// AuthenticatorSimpleFactory .
type AuthenticatorSimpleFactory struct{}

// Create .
func (factory AuthenticatorSimpleFactory) Create(auth *pb.Auth) Authenticator {
	var authenticator Authenticator

	switch auth.Type {
	default:
		authenticator = &AuthenticatorJWT{
			auth: auth,
		}
	}

	return authenticator
}
