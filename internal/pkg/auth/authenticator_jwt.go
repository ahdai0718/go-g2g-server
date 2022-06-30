package auth

import (
	"fmt"
	"strings"

	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"

	"github.com/dgrijalva/jwt-go"
)

// AuthenticatorJWT .
type AuthenticatorJWT struct {
	auth *pb.Auth
}

// Type .
func (authenticator *AuthenticatorJWT) Type() pb.AuthType {
	return authenticator.auth.Type
}

// Verify .
func (authenticator *AuthenticatorJWT) Verify(token string) (Result, error) {
	result := Result{}
	secret := authenticator.auth.Secret

	if strings.Trim(token, " ") == "" {
		return result, fmt.Errorf("token is empty")
	}

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if err != nil {
		glog.Error(err)
		return result, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		if _, isExists := claims["player_id"]; isExists {
			result.PlayerID = claims["player_id"].(string)
		}
	} else {
		glog.Error(err)
	}

	return result, err
}

// GenToken .
func (authenticator *AuthenticatorJWT) GenToken(userID string) (string, error) {
	secret := authenticator.auth.Secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"player_id": userID,
	})
	return token.SignedString([]byte(secret))
}
