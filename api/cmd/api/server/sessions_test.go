package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/birthday-app-api/internal/auth"
	"github.com/yigitsadic/birthday-app-api/internal/sessions"
)

var (
	sessionRepo = sessions.NewMockSessionStore()
	jwtStore    = auth.NewJWT("1234")
)

func makeCreateSessionRequest(t *testing.T, body string) (int, []*http.Cookie) {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/sessions/create",
		bytes.NewBufferString(body),
	)

	s := Server{SessionRepository: &sessionRepo, JWTStore: jwtStore}
	s.HandleCreateSession(w, req)

	res := w.Result()
	defer res.Body.Close()

	return res.StatusCode, res.Cookies()
}

func makeRefreshSessionRequest(t *testing.T, cookie *http.Cookie) (int, []*http.Cookie) {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/sessions/refresh",
		nil,
	)
	req.AddCookie(cookie)

	s := Server{SessionRepository: &sessionRepo, JWTStore: jwtStore}
	s.HandleRefreshSession(w, req)

	res := w.Result()
	defer res.Body.Close()

	return res.StatusCode, res.Cookies()
}

func Test_CreateSession(t *testing.T) {
	good_input := `{
		"email": "yigit@google.com",
		"password": "123456789"
	}`

	t.Run("it should handle bad input", func(t *testing.T) {
		status, _ := makeCreateSessionRequest(t, "lord of the rings")
		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("it should handle not found email", func(t *testing.T) {
		sessionRepo.RaiseNotFound = true
		defer func() {
			sessionRepo.RaiseNotFound = false
		}()

		status, _ := makeCreateSessionRequest(t, good_input)
		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("it should handle not matching password", func(t *testing.T) {
		not_matching_input := `{
			"email": "yigit@google.com",
			"password": "987654321"
		}`

		status, _ := makeCreateSessionRequest(t, not_matching_input)
		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("it should return auth response if everything ok", func(t *testing.T) {
		require.False(t, sessionRepo.RaiseNotFound)

		status, cookies := makeCreateSessionRequest(t, good_input)

		assert.Equal(t, http.StatusCreated, status)
		assert.True(t, len(cookies) > 0)
		assert.Equal(t, RefreshTokenCookieName, cookies[0].Name)
		assert.True(t, cookies[0].HttpOnly)
	})
}

func Test_RefreshToken(t *testing.T) {
	t.Run("it should return 401 if refresh token not present", func(t *testing.T) {
		status, _ := makeRefreshSessionRequest(t, &http.Cookie{})

		assert.Equal(t, http.StatusUnauthorized, status)
	})

	t.Run("it should return 401 if called with an access token", func(t *testing.T) {
		tok, err := jwtStore.GenerateAccessToken("1", "2")
		require.Nil(t, err)

		status, _ := makeRefreshSessionRequest(t, &http.Cookie{
			Name:     RefreshTokenCookieName,
			Value:    tok,
			HttpOnly: true,
		})

		assert.Equal(t, http.StatusUnauthorized, status)
	})

	t.Run("it should return 401 if no user found", func(t *testing.T) {
		sessionRepo.RaiseNotFound = true
		defer func() {
			sessionRepo.RaiseNotFound = false
		}()

		tok, err := jwtStore.GenerateRefreshToken("99")
		require.Nil(t, err)

		status, _ := makeRefreshSessionRequest(t, &http.Cookie{
			Name:     RefreshTokenCookieName,
			Value:    tok,
			HttpOnly: true,
		})

		assert.Equal(t, http.StatusUnauthorized, status)
	})

	t.Run("it should refresh token if everyhing ok", func(t *testing.T) {
		tok, err := jwtStore.GenerateRefreshToken("1")
		require.Nil(t, err)

		status, cookies := makeRefreshSessionRequest(t, &http.Cookie{
			Name:     RefreshTokenCookieName,
			Value:    tok,
			HttpOnly: true,
		})

		assert.Equal(t, http.StatusCreated, status)
		assert.True(t, len(cookies) > 0)
		assert.Equal(t, RefreshTokenCookieName, cookies[0].Name)
		assert.True(t, cookies[0].HttpOnly)
	})
}
