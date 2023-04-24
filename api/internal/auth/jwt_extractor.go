package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthTokenFoundOnHeader = errors.New("no auth token found on header")

// GetJWTFromHeader extracts token from header.
func GetJWTFromHeader(r *http.Request, jStore JWTStore) (*JWTClaims, error) {
	headerValue := r.Header.Get("Authorization")
	tok := strings.ReplaceAll(headerValue, "Bearer ", "")

	if tok == "" {
		return nil, ErrNoAuthTokenFoundOnHeader
	}

	return jStore.ParseToken(tok)
}
