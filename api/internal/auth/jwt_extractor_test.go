package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var jStore JWTStore

func init() {
	jStore = JWTStore{[]byte("ABCDEFG")}
}

func Test_GetJWTFromHeader(t *testing.T) {
	t.Run("it should return an error if no header found", func(t *testing.T) {
		var jwtErr error

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, jwtErr = GetJWTFromHeader(r, jStore)
		}))
		defer srv.Close()

		_, err := http.DefaultClient.Get(srv.URL)

		require.Nil(t, err)
		assert.NotNil(t, jwtErr)
		assert.ErrorIs(t, jwtErr, ErrNoAuthTokenFoundOnHeader)
	})

	t.Run("it should extract token from header", func(t *testing.T) {
		tok, err := jStore.GenerateAccessToken("1", "2")
		require.Nil(t, err)

		c := &JWTClaims{}

		r := chi.NewRouter()
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			c, err = GetJWTFromHeader(r, jStore)

			if err != nil {
				t.Fatalf("should not get an error: %s", err)
			}
		})

		srv := httptest.NewServer(r)
		defer srv.Close()

		req, err := http.NewRequest(
			http.MethodGet,
			srv.URL,
			nil,
		)
		require.Nil(t, err)

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tok))

		_, err = http.DefaultClient.Do(req)

		assert.Nil(t, err)

		assert.Equal(t, "1", c.Subject)
		assert.Equal(t, "2", c.CompanyId)
		assert.Equal(t, AccessTokenKey, c.TokenType)
	})
}
