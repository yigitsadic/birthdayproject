package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func init() {
	time.Local = time.UTC
}

const (
	AccessTokenKey  = "AccessToken"
	RefreshTokenKey = "RefreshToken"
)

type JWTClaims struct {
	CompanyId string `json:"company_id,omitempty"`
	TokenType string `json:"token_type"`

	jwt.RegisteredClaims
}

type JWTStore struct {
	secret []byte
}

// NewJWT creates a new instance of JWT functionality.
func NewJWT(secret string) JWTStore {
	return JWTStore{secret: []byte(secret)}
}

// GenerateAccessToken generates JWT token with given parameters as claims.
func (j JWTStore) GenerateAccessToken(user_id, company_id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		TokenType: AccessTokenKey,
		CompanyId: company_id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "birthday-app",
			Subject:   user_id,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(time.Minute * 15))),
		},
	})

	return token.SignedString(j.secret)
}

// GenerateRefreshToken generates JWT token that lasts 7 days with given parameters as claims.
func (j JWTStore) GenerateRefreshToken(user_id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		TokenType: RefreshTokenKey,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "birthday-app",
			Subject:   user_id,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(time.Hour * 24 * 7))),
		},
	})

	return token.SignedString(j.secret)
}

// ParseToken parses given token and returns JWT claims that stores
// first_name, last_name like values.
func (j JWTStore) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(tok *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		if err == nil {
			err = errors.New("token is not valid")
		}

		return nil, err
	}
}

// RefreshToken creates new JWT token with same claims from given token.
func (j JWTStore) RefreshToken(token string) (string, error) {
	parsed, err := j.ParseToken(token)

	if err != nil {
		return "", err
	}

	if parsed.TokenType == AccessTokenKey {
		return j.GenerateAccessToken(parsed.Subject, parsed.CompanyId)
	} else if parsed.TokenType == RefreshTokenKey {
		return j.GenerateRefreshToken(parsed.Subject)
	} else {
		return "", errors.New("unknown token type")
	}
}
