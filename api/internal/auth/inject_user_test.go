package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/birthday-app-api/internal/common"
)

type valueFound struct {
	CompanyIdFound bool
	CompanyId      string
	UserIdFound    bool
	UserId         string
}

func init() {
	jStore = NewJWT("lorem")
}

func genInjectUserSrv(t *testing.T, v *valueFound) *httptest.Server {
	t.Helper()

	r := chi.NewRouter()

	r.Use(InjectUser(jStore))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		userId, userIdOk := r.Context().Value(common.UserIdCtxKey).(string)
		companyId, companyIdOk := r.Context().Value(common.CompanyIdCtxKey).(string)

		if userIdOk {
			v.UserIdFound = true
			v.UserId = userId
		}

		if companyIdOk {
			v.CompanyIdFound = true
			v.CompanyId = companyId
		}
	})

	srv := httptest.NewServer(r)

	return srv
}

func TestInjectUser(t *testing.T) {
	t.Run("it should successfully inject user id and company id", func(t *testing.T) {
		v := &valueFound{}

		srv := genInjectUserSrv(t, v)
		defer srv.Close()

		req, err := http.NewRequest(
			http.MethodGet,
			srv.URL,
			nil,
		)
		require.Nil(t, err)

		userId := "123"
		companyId := "4123"

		tok, err := jStore.GenerateAccessToken(userId, companyId)
		require.Nil(t, err)

		req.Header.Add("Authorization", "Bearer "+tok)

		_, err = http.DefaultClient.Do(req)

		assert.Nil(t, err)
		assert.True(t, v.UserIdFound)
		assert.Equal(t, userId, v.UserId)
		assert.True(t, v.CompanyIdFound)
		assert.Equal(t, companyId, v.CompanyId)
	})

	t.Run("it should fail silently when an error occurs", func(t *testing.T) {
		v := &valueFound{}

		srv := genInjectUserSrv(t, v)
		defer srv.Close()

		_, err := http.DefaultClient.Get(srv.URL)

		require.Nil(t, err)
		assert.False(t, v.UserIdFound)
		assert.False(t, v.CompanyIdFound)
		assert.Equal(t, "", v.UserId)
		assert.Equal(t, "", v.CompanyId)
	})
}
