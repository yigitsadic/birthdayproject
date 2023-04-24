package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/birthday-app-api/internal/users"
)

var userMockRepo = users.NewMockUserStore()

func createHandleUserDetailResp(t *testing.T, userId string) int {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	ctxWithRoute := injectParamKeyToChiCtx(t, req, []string{"id"}, []string{userId})
	authReq := loginWithRequestLevel(t, ctxWithRoute, "1", strconv.Itoa(userMockRepo.Store[0].ID))

	s := Server{UserRepository: userMockRepo}
	s.HandleUserDetail(w, authReq)

	res := w.Result()
	defer res.Body.Close()

	return res.StatusCode
}

func createHandleUserUpdateResp(t *testing.T, userId string, bodyAsStr string) int {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBufferString(bodyAsStr))

	ctxWithRoute := injectParamKeyToChiCtx(t, req, []string{"id"}, []string{userId})
	authReq := loginWithRequestLevel(t, ctxWithRoute, "1", strconv.Itoa(userMockRepo.Store[0].ID))

	s := Server{UserRepository: userMockRepo}
	s.HandleUserUpdate(w, authReq)

	res := w.Result()
	defer res.Body.Close()

	return res.StatusCode
}

// func Test_MountUserRoutes(t *testing.T) {
// 	router := chi.NewRouter()
// 	loginRouterLevel(t, router, "1", "2")
//
// 	MountUserRoutes(router, userMockRepo)
// 	srv := httptest.NewServer(router)
// 	defer srv.Close()
//
// 	t.Run("it should mount user detail handler", func(t *testing.T) {
// 		req, err := http.NewRequest(
// 			http.MethodGet,
// 			fmt.Sprintf("%s/users/%d", srv.URL, userMockRepo.Store[0].ID),
// 			nil,
// 		)
// 		require.Nil(t, err)
//
// 		resp, err := http.DefaultClient.Do(req)
// 		require.Nil(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	})
//
// 	t.Run("it should mount user update handler", func(t *testing.T) {
// 		goodInput := `{
// 			"first_name": "Yigit",
// 			"last_name": "Sadic"
// 		}`
//
// 		require.False(t, userMockRepo.RaiseErrorOnUpdate)
//
// 		req, err := http.NewRequest(
// 			http.MethodPut,
// 			fmt.Sprintf("%s/users/%d", srv.URL, userMockRepo.Store[0].ID),
// 			bytes.NewBufferString(goodInput),
// 		)
// 		require.Nil(t, err)
//
// 		resp, err := http.DefaultClient.Do(req)
// 		require.Nil(t, err)
//
// 		defer resp.Body.Close()
//
// 		var parsed users.UserModel
// 		err = json.NewDecoder(resp.Body).Decode(&parsed)
//
// 		require.Nil(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 		assert.Equal(t, "Yigit", parsed.FirstName)
// 		assert.Equal(t, "Sadic", parsed.LastName)
// 	})
// }

func Test_handleUserDetail(t *testing.T) {
	t.Run("it should match authenticated user and route param", func(t *testing.T) {
		status := createHandleUserDetailResp(t, "666")

		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("it should handle not found user", func(t *testing.T) {
		userMockRepo.RaiseErrorOnGetUser = true
		defer func() {
			userMockRepo.RaiseErrorOnGetUser = false
		}()

		status := createHandleUserDetailResp(t, strconv.Itoa(userMockRepo.Store[0].ID))
		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("it should handle success", func(t *testing.T) {
		status := createHandleUserDetailResp(t, strconv.Itoa(userMockRepo.Store[0].ID))
		assert.Equal(t, http.StatusOK, status)
	})
}

func Test_handleUserUpdate(t *testing.T) {
	goodInput := `{
		"first_name": "Yigit",
		"last_name": "Sadic"
	}`

	t.Run("it should check authenticated user and route param match", func(t *testing.T) {
		status := createHandleUserUpdateResp(t, "666", goodInput)

		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("it should handle bad input", func(t *testing.T) {
		status := createHandleUserUpdateResp(
			t,
			strconv.Itoa(userMockRepo.Store[0].ID),
			"bad romance",
		)

		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("it should handle invalid input", func(t *testing.T) {
		invalidInput := `{
			"first_name": "",
			"last_name": "S"
		}`
		status := createHandleUserUpdateResp(
			t,
			strconv.Itoa(userMockRepo.Store[0].ID),
			invalidInput,
		)

		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("it should handle service error", func(t *testing.T) {
		userMockRepo.RaiseErrorOnUpdate = true
		defer func() {
			userMockRepo.RaiseErrorOnUpdate = false
		}()

		status := createHandleUserUpdateResp(t, strconv.Itoa(userMockRepo.Store[0].ID), goodInput)

		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("it should update user", func(t *testing.T) {
		status := createHandleUserUpdateResp(t, strconv.Itoa(userMockRepo.Store[0].ID), goodInput)

		assert.Equal(t, http.StatusOK, status)
	})
}
