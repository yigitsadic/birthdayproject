package e2e

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getEmployee(db *sql.DB) (companyId int, employeeId int, err error) {
	err = db.QueryRow("select id, company_id from employees limit 1").Scan(&employeeId, &companyId)
	return
}

func genListEndpoint(url string, companyId int) string {
	return fmt.Sprintf("%s/companies/%d/employees", url, companyId)
}

func genDetailEndpoint(url string, companyId int, employeeId int) string {
	return fmt.Sprintf("%s/companies/%d/employees/%d", url, companyId, employeeId)
}

func (testSuite *IntegrationTestSuite) Test_Employees_List() {
	assert := assert.New(testSuite.T())

	companyId, _, err := getEmployee(testSuite.DB)
	require.Nil(testSuite.T(), err)

	token, err := Login(testSuite.TestServer.URL)
	assert.Nil(err)

	testSuite.Run("it should render list of employees", func() {
		req, err := http.NewRequest(
			http.MethodGet,
			genListEndpoint(testSuite.TestServer.URL, companyId),
			nil,
		)
		assert.Nil(err)

		req.Header.Add("Authorization", "Bearer "+token)

		res, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	testSuite.Run("it should return 404 when not found company", func() {
		req, err := http.NewRequest(
			http.MethodGet,
			genListEndpoint(testSuite.TestServer.URL, 9999),
			nil,
		)
		assert.Nil(err)

		req.Header.Add("Authorization", "Bearer "+token)

		res, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusNotFound, res.StatusCode)
	})
}

func (testSuite *IntegrationTestSuite) Test_Employees_Detail() {
	assert := assert.New(testSuite.T())

	companyId, employeeId, err := getEmployee(testSuite.DB)
	require.Nil(testSuite.T(), err)

	token, err := Login(testSuite.TestServer.URL)
	assert.Nil(err)

	testSuite.Run("it should return an employee", func() {
		req, err := http.NewRequest(
			http.MethodGet,
			genDetailEndpoint(testSuite.TestServer.URL, companyId, employeeId),
			nil,
		)
		assert.Nil(err)

		req.Header.Add("Authorization", "Bearer "+token)

		res, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	testSuite.Run("it should return 404 when not found employee", func() {
		req, err := http.NewRequest(
			http.MethodGet,
			genDetailEndpoint(testSuite.TestServer.URL, companyId, 99999),
			nil,
		)
		assert.Nil(err)

		req.Header.Add("Authorization", "Bearer "+token)

		res, err := http.DefaultClient.Do(req)

		assert.Nil(err)
		assert.Equal(http.StatusNotFound, res.StatusCode)
	})
}

func (testSuite *IntegrationTestSuite) Test_Employees_Create() {
	assert := assert.New(testSuite.T())

	assert.True(true)
}

func (testSuite *IntegrationTestSuite) Test_Employees_Update() {
	assert := assert.New(testSuite.T())

	assert.True(true)
}

func (testSuite *IntegrationTestSuite) Test_Employees_Delete() {
	assert := assert.New(testSuite.T())

	assert.True(true)
}
