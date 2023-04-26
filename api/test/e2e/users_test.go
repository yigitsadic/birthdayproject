package e2e

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func (testSuite *IntegrationTestSuite) Test_Users_Detail() {
	assert := assert.New(testSuite.T())

	var recordId int
	err := testSuite.DB.QueryRow("select id from users limit 1").Scan(&recordId)
	assert.Nil(err)

	token, err := Login(testSuite.TestServer.URL)

	assert.Nil(err)

	genEndpoint := func(id int) string {
		return fmt.Sprintf("%s/users/%d", testSuite.TestServer.URL, id)
	}

	testSuite.Run("it should render details of user", func() {
		req, err := http.NewRequest(http.MethodGet, genEndpoint(recordId), nil)
		assert.Nil(err)

		req.Header.Add("Authorization", "Bearer "+token)

		res, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	testSuite.Run("it should return not found if not found", func() {
		req, err := http.NewRequest(http.MethodGet, genEndpoint(99999), nil)
		assert.Nil(err)

		req.Header.Add("Authorization", "Bearer "+token)

		res, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusNotFound, res.StatusCode)
	})
}

func (testSuite *IntegrationTestSuite) Test_Users_Update() {
	assert := assert.New(testSuite.T())

	var recordId int
	err := testSuite.DB.QueryRow("select id from users limit 1").Scan(&recordId)
	assert.Nil(err)

	token, err := Login(testSuite.TestServer.URL)

	assert.Nil(err)

	genEndpoint := func(id int) string {
		return fmt.Sprintf("%s/users/%d", testSuite.TestServer.URL, id)
	}

	testSuite.Run("it should update user data", func() {
		goodBody := `{"first_name": "John", "last_name": "Doe"}`
		req, err := http.NewRequest(
			http.MethodPut,
			genEndpoint(recordId),
			bytes.NewBufferString(goodBody),
		)
		req.Header.Add("Authorization", "Bearer "+token)

		assert.Nil(err)

		resp, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusOK, resp.StatusCode)
	})

	testSuite.Run("it should handle bad input", func() {
		badInput := `{"first_name": "A", "last_name": ""}`
		req, err := http.NewRequest(
			http.MethodPut,
			genEndpoint(recordId),
			bytes.NewBufferString(badInput),
		)
		req.Header.Add("Authorization", "Bearer "+token)

		assert.Nil(err)

		resp, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
	})

	testSuite.Run("it should handle if authenticated tries to access another resource", func() {
		goodBody := `{"first_name": "John", "last_name": "Doe"}`
		req, err := http.NewRequest(
			http.MethodPut,
			genEndpoint(9999999),
			bytes.NewBufferString(goodBody),
		)
		req.Header.Add("Authorization", "Bearer "+token)

		assert.Nil(err)

		resp, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusNotFound, resp.StatusCode)
	})
}
