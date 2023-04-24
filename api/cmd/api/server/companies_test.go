package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/birthday-app-api/internal/companies"
)

var companyMockRepo = companies.NewMockCompanyStore()

func createHandleCompanyDetailResp(t *testing.T, companyId string) (int, companies.CompanyModel) {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	ctxWithRoute := injectParamKeyToChiCtx(t, req, []string{"id"}, []string{companyId})
	authReq := loginWithRequestLevel(
		t,
		ctxWithRoute,
		"1",
		strconv.Itoa(companyMockRepo.Store[0].ID),
	)

	s := Server{CompanyRepository: companyMockRepo}
	s.HandleCompanyDetail(w, authReq)

	res := w.Result()
	defer res.Body.Close()

	parsed := companies.CompanyModel{}
	json.NewDecoder(res.Body).Decode(&parsed)

	return res.StatusCode, parsed
}

func createHandleCompanyUpdateResp(
	t *testing.T,
	companyId string,
	bodyAsStr string,
) (int, companies.CompanyModel) {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBufferString(bodyAsStr))

	ctxWithRoute := injectParamKeyToChiCtx(t, req, []string{"id"}, []string{companyId})
	authReq := loginWithRequestLevel(
		t,
		ctxWithRoute,
		"1",
		strconv.Itoa(companyMockRepo.Store[0].ID),
	)

	s := Server{CompanyRepository: companyMockRepo}

	s.HandleCompanyUpdate(w, authReq)

	res := w.Result()
	defer res.Body.Close()

	parsed := companies.CompanyModel{}
	json.NewDecoder(res.Body).Decode(&parsed)

	return res.StatusCode, parsed
}

func Test_handleCompanyDetail(t *testing.T) {
	t.Run("it should handle not found", func(t *testing.T) {
		companyMockRepo.RaiseErrorOnFind = true
		defer func() {
			companyMockRepo.RaiseErrorOnFind = false
		}()

		statusCode, _ := createHandleCompanyDetailResp(t, strconv.Itoa(companyMockRepo.Store[0].ID))
		assert.Equal(t, http.StatusNotFound, statusCode)
	})

	t.Run(
		"it should handle authenticated user's company and requested different",
		func(t *testing.T) {
			statusCode, _ := createHandleCompanyDetailResp(t, "666")

			assert.Equal(t, http.StatusNotFound, statusCode)
		},
	)

	t.Run("it should render successful response", func(t *testing.T) {
		statusCode, parsed := createHandleCompanyDetailResp(
			t,
			strconv.Itoa(companyMockRepo.Store[0].ID),
		)

		assert.Equal(t, http.StatusOK, statusCode)
		assert.Equal(t, companyMockRepo.Store[0].Name, parsed.Name)
	})
}

func Test_handleCompanyUpdate(t *testing.T) {
	goodBody := `{"name": "very nice company"}`
	recordId := strconv.Itoa(companyMockRepo.Store[0].ID)

	t.Run("it should check authenticated company id and param", func(t *testing.T) {
		status, _ := createHandleCompanyUpdateResp(t, "666", goodBody)

		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("it should respond 422 with bad body", func(t *testing.T) {
		status, _ := createHandleCompanyUpdateResp(t, recordId, "Gandalf")

		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("it should respond 422 with invalid parameters", func(t *testing.T) {
		status, _ := createHandleCompanyUpdateResp(t, recordId, `{"name": "sh"}`)

		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("it should handle service error", func(t *testing.T) {
		companyMockRepo.RaiseErrorOnUpdate = true
		defer func() {
			companyMockRepo.RaiseErrorOnUpdate = false
		}()

		status, _ := createHandleCompanyUpdateResp(t, recordId, goodBody)

		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("it should respond with success", func(t *testing.T) {
		status, _ := createHandleCompanyUpdateResp(t, recordId, goodBody)

		assert.Equal(t, http.StatusOK, status)
	})
}
