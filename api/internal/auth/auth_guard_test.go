package auth

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/birthday-app-api/internal/common"
)

func TestAuthGuard(t *testing.T) {
	t.Run("it should render unauthorized response if not user-id found on ctx", func(t *testing.T) {
		r := chi.NewRouter()
		r.Use(AuthGuard)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		})

		srv := httptest.NewServer(r)
		defer srv.Close()

		resp, err := http.DefaultClient.Get(srv.URL)
		require.Nil(t, err)
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)

		require.Nil(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.NotEqual(t, "hello", string(content))
	})

	t.Run("it should continue if user-id found on ctx", func(t *testing.T) {
		r := chi.NewRouter()

		r.Use(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r1 := r.WithContext(context.WithValue(r.Context(), common.UserIdCtxKey, "123"))
				r2 := r.WithContext(context.WithValue(r1.Context(), common.CompanyIdCtxKey, "321"))

				h.ServeHTTP(w, r2)
			})
		})
		r.Use(AuthGuard)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		})

		srv := httptest.NewServer(r)
		defer srv.Close()

		resp, err := http.DefaultClient.Get(srv.URL)
		require.Nil(t, err)
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)

		require.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "hello", string(content))
	})
}
