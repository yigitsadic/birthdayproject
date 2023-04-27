package e2e

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/birthday-app-api/internal/server"
)

func (testSuite *IntegrationTestSuite) Test_Sessions_Create() {
	assert := assert.New(testSuite.T())
	body := bytes.NewBufferString(`{
		"email":    "johndo@google.com",
		"password": "123456789"
    }`)
	endpoint := fmt.Sprintf("%s/sessions/create", testSuite.TestServer.URL)

	testSuite.Run("it should successfully login", func() {
		req, err := http.NewRequest(http.MethodPost, endpoint, body)
		assert.Nil(err)

		res, err := http.DefaultClient.Do(req)
		assert.Nil(err)

		assert.Equal(http.StatusCreated, res.StatusCode)
	})

	testSuite.Run("it should handle non present user email", func() {
		nonPresentBody := bytes.NewBufferString(`{
			"email": "something@twitter.com",
			"password": "123132"
			}`)
		req, err := http.NewRequest(http.MethodPost, endpoint, nonPresentBody)
		assert.Nil(err)

		res, err := http.DefaultClient.Do(req)
		assert.Nil(err)

		assert.Equal(http.StatusUnprocessableEntity, res.StatusCode)
	})

	testSuite.Run("it should handle wrong password", func() {
		wrongPasswordBody := bytes.NewBufferString(`{
		"email":    "johndo@google.com",
		"password": "12332123456789"
    }`)

		req, err := http.NewRequest(http.MethodPost, endpoint, wrongPasswordBody)
		assert.Nil(err)

		res, err := http.DefaultClient.Do(req)
		assert.Nil(err)

		assert.Equal(http.StatusUnprocessableEntity, res.StatusCode)
	})
}

func (testSuite *IntegrationTestSuite) Test_Sessions_Update() {
	assert := assert.New(testSuite.T())
	endpoint := fmt.Sprintf("%s/sessions/refresh", testSuite.TestServer.URL)

	var userId int
	err := testSuite.DB.QueryRow("select id from users limit 1").
		Scan(&userId)
	assert.Nil(err)

	genReq := func() *http.Request {
		req, err := http.NewRequest(http.MethodPost, endpoint, nil)
		assert.Nil(err)

		return req
	}

	testSuite.Run("it should generate refresh token for valid token", func() {
		refreshToken, err := testSuite.JWTStore.GenerateRefreshToken(
			strconv.Itoa(userId),
		)

		assert.Nil(err)

		req := genReq()
		req.AddCookie(&http.Cookie{
			Name:     server.RefreshTokenCookieName,
			Value:    refreshToken,
			HttpOnly: true,
		})

		resp, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusCreated, resp.StatusCode)
	})

	testSuite.Run("it should return 401 if user not found", func() {
		refreshToken, err := testSuite.JWTStore.GenerateRefreshToken("99999999")

		assert.Nil(err)

		req := genReq()
		req.AddCookie(&http.Cookie{
			Name:     server.RefreshTokenCookieName,
			Value:    refreshToken,
			HttpOnly: true,
		})

		resp, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusUnauthorized, resp.StatusCode)
	})
}
