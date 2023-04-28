package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/birthday-app-api/internal/employees"
)

var mockEmployeeStore = employees.NewMockEmployeeStore()

func Test_employeeListHandler(t *testing.T) {
	makeReq := func(companyId string) int {
		t.Helper()

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		ctxWithRoute := injectParamKeyToChiCtx(t, req, []string{"companyId"}, []string{companyId})
		authReq := loginWithRequestLevel(t, ctxWithRoute, "1", "1")

		s := Server{EmployeeRepository: mockEmployeeStore}
		s.EmployeeListHandler(w, authReq)

		res := w.Result()
		defer res.Body.Close()

		return res.StatusCode
	}

	t.Run(
		"it should check authorized user's company id and queried company id",
		func(t *testing.T) {
			statusCode := makeReq("2")

			assert.Equal(t, http.StatusNotFound, statusCode)
		},
	)

	t.Run("it should handle repository error", func(t *testing.T) {
		mockEmployeeStore.RaiseErrorOnOperation = true
		defer func() {
			mockEmployeeStore.RaiseErrorOnOperation = false
		}()

		statusCode := makeReq("1")

		assert.Equal(t, http.StatusInternalServerError, statusCode)
	})

	t.Run("it should not give an error when list is empty", func(t *testing.T) {
		var copied []*employees.EmployeeModel

		copied = append(copied, mockEmployeeStore.Store...)

		mockEmployeeStore.Store = []*employees.EmployeeModel{}

		statusCode := makeReq("1")

		mockEmployeeStore.Store = copied
		assert.Equal(t, 1, len(mockEmployeeStore.Store))
		assert.Equal(t, http.StatusOK, statusCode)
	})

	t.Run("it should serve employee list", func(t *testing.T) {
		statusCode := makeReq("1")

		assert.Equal(t, http.StatusOK, statusCode)
	})
}

func Test_employeeCreateHandler(t *testing.T) {
	makeReq := func(dto employees.EmployeeDto, companyId string) int {
		t.Helper()

		body, err := json.Marshal(&dto)
		require.Nil(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))

		ctxWithRoute := injectParamKeyToChiCtx(t, req, []string{"companyId"}, []string{companyId})
		authReq := loginWithRequestLevel(t, ctxWithRoute, "1", "1")

		s := Server{EmployeeRepository: mockEmployeeStore}
		s.EmployeeCreateHandler(w, authReq)

		res := w.Result()
		defer res.Body.Close()

		return res.StatusCode
	}

	t.Run("it should validate bad input", func(t *testing.T) {
		statusCode := makeReq(employees.EmployeeDto{}, "1")

		assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
	})

	t.Run("it should validate invalid input", func(t *testing.T) {
		statusCode := makeReq(employees.EmployeeDto{
			CompanyId:  1,
			FirstName:  "lorem",
			LastName:   "ipsum",
			Email:      "loremipsum@google.com",
			BirthDay:   15,
			BirthMonth: 13,
		}, "1")

		assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
	})

	t.Run("it should check authenticated company id and requested company id", func(t *testing.T) {
		statusCode := makeReq(employees.EmployeeDto{
			CompanyId:  1,
			FirstName:  "lorem",
			LastName:   "ipsum",
			Email:      "loremipsum@google.com",
			BirthDay:   15,
			BirthMonth: 13,
		}, "2")

		assert.Equal(t, http.StatusNotFound, statusCode)
	})

	t.Run("it should handle repository error", func(t *testing.T) {
		mockEmployeeStore.RaiseErrorOnOperation = true
		defer func() {
			mockEmployeeStore.RaiseErrorOnOperation = false
		}()

		statusCode := makeReq(employees.EmployeeDto{
			CompanyId:  1,
			FirstName:  "lorem",
			LastName:   "ipsum",
			Email:      "loremipsum@google.com",
			BirthDay:   15,
			BirthMonth: 4,
		}, "1")

		assert.Equal(t, http.StatusInternalServerError, statusCode)
	})

	t.Run("it should respond with 201 when ok", func(t *testing.T) {
		statusCode := makeReq(employees.EmployeeDto{
			CompanyId:  1,
			FirstName:  "lorem",
			LastName:   "ipsum",
			Email:      "loremipsum@google.com",
			BirthDay:   15,
			BirthMonth: 4,
		}, "1")

		assert.Equal(t, http.StatusCreated, statusCode)
	})
}

func Test_employeeDetailHandler(t *testing.T) {
	makeReq := func(companyId, employeeId string) int {
		t.Helper()

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		ctxWithRoute := injectParamKeyToChiCtx(
			t,
			req,
			[]string{"companyId", "id"},
			[]string{companyId, employeeId},
		)
		authReq := loginWithRequestLevel(t, ctxWithRoute, "1", "1")

		s := Server{EmployeeRepository: mockEmployeeStore}
		s.EmployeeDetailHandler(w, authReq)

		res := w.Result()
		defer res.Body.Close()

		return res.StatusCode
	}

	t.Run("it should match with company id with auth", func(t *testing.T) {
		statusCode := makeReq("2", "1")

		assert.Equal(t, http.StatusNotFound, statusCode)
	})

	t.Run("it should handle repository error", func(t *testing.T) {
		mockEmployeeStore.RaiseErrorOnOperation = true
		defer func() {
			mockEmployeeStore.RaiseErrorOnOperation = false
		}()

		statusCode := makeReq("1", "1")

		assert.Equal(t, http.StatusInternalServerError, statusCode)
	})

	t.Run("it should handle success", func(t *testing.T) {
		statusCode := makeReq("1", "1")

		assert.Equal(t, http.StatusOK, statusCode)
	})
}

func Test_employeeUpdateHandler(t *testing.T) {
	validDto := employees.EmployeeDto{
		FirstName:  "Yigit",
		LastName:   "Sadic",
		Email:      "yigitsad@google.com",
		BirthDay:   14,
		BirthMonth: 2,
	}

	makeReq := func(companyId string, employeeId string, dto employees.EmployeeDto) int {
		t.Helper()

		body, err := json.Marshal(&dto)
		require.Nil(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))

		ctxWithRoute := injectParamKeyToChiCtx(
			t,
			req,
			[]string{"companyId", "id"},
			[]string{companyId, employeeId},
		)
		authReq := loginWithRequestLevel(t, ctxWithRoute, "1", "1")

		s := Server{EmployeeRepository: mockEmployeeStore}
		s.EmployeeUpdateHandler(w, authReq)

		res := w.Result()

		defer res.Body.Close()

		return res.StatusCode
	}

	t.Run("it should validate auth company id and param company id", func(t *testing.T) {
		statusCode := makeReq("2", "1", validDto)

		assert.Equal(t, http.StatusNotFound, statusCode)
	})

	t.Run("it should validate input", func(t *testing.T) {
		invalidDto := employees.EmployeeDto{
			FirstName:  "Yigit",
			LastName:   "Sadic",
			Email:      "yigitsad@google.com",
			BirthDay:   44,
			BirthMonth: 15,
		}

		statusCode := makeReq("1", "1", invalidDto)

		assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
	})

	t.Run("it should handle repository error", func(t *testing.T) {
		mockEmployeeStore.RaiseErrorOnOperation = true
		defer func() {
			mockEmployeeStore.RaiseErrorOnOperation = false
		}()

		statusCode := makeReq("1", "1", validDto)

		assert.Equal(t, http.StatusInternalServerError, statusCode)
	})

	t.Run("it should respond with success", func(t *testing.T) {
		statusCode := makeReq("1", "1", validDto)

		assert.Equal(t, http.StatusOK, statusCode)
	})
}

func Test_employeeDeleteHandler(t *testing.T) {
	makeReq := func(companyId, employeeId string) int {
		t.Helper()

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)

		ctxWithRoute := injectParamKeyToChiCtx(
			t,
			req,
			[]string{"companyId", "id"},
			[]string{companyId, employeeId},
		)
		authReq := loginWithRequestLevel(t, ctxWithRoute, "1", "1")

		s := Server{EmployeeRepository: mockEmployeeStore}
		s.EmployeeDeleteHandler(w, authReq)

		res := w.Result()
		defer res.Body.Close()

		return res.StatusCode
	}

	t.Run("it should validate auth company id and param company id", func(t *testing.T) {
		statusCode := makeReq("2", "1")

		assert.Equal(t, http.StatusNotFound, statusCode)
	})

	t.Run("it should handle repository error", func(t *testing.T) {
		mockEmployeeStore.RaiseErrorOnOperation = true
		defer func() {
			mockEmployeeStore.RaiseErrorOnOperation = false
		}()

		statusCode := makeReq("1", "1")

		assert.Equal(t, http.StatusInternalServerError, statusCode)
	})

	t.Run("it should handle success", func(t *testing.T) {
		statusCode := makeReq("1", "1")

		assert.Equal(t, http.StatusOK, statusCode)
	})
}

// func Test_MountEmployeesRoutes(t *testing.T) {
// 	r := chi.NewRouter()
//
// 	loginRouterLevel(t, r, "1", "1")
// 	MountEmployeesRoutes(r, &mockEmployeeStore)
//
// 	srv := httptest.NewServer(r)
// 	defer srv.Close()
//
// 	employeeGeneralUrl := srv.URL + "/companies/1/employees"
// 	employeeDetailUrl := srv.URL + "/companies/1/employees/1"
//
// 	t.Run("it should mount employeeUpdateHandler", func(t *testing.T) {
// 		dto := employees.EmployeeDto{
// 			FirstName:  "lorem",
// 			LastName:   "ipsum",
// 			Email:      "loremipsum@google.com",
// 			BirthDay:   15,
// 			BirthMonth: 4,
// 		}
// 		body, err := json.Marshal(dto)
// 		require.Nil(t, err)
//
// 		req, err := http.NewRequest(http.MethodPut, employeeDetailUrl, bytes.NewBuffer(body))
// 		require.Nil(t, err)
//
// 		resp, err := http.DefaultClient.Do(req)
//
// 		require.Nil(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	})
//
// 	t.Run("it should mount employeeDeleteHandler", func(t *testing.T) {
// 		req, err := http.NewRequest(http.MethodDelete, employeeDetailUrl, nil)
// 		require.Nil(t, err)
//
// 		resp, err := http.DefaultClient.Do(req)
//
// 		require.Nil(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	})
//
// 	t.Run("it should mount employeeCreateHandler", func(t *testing.T) {
// 		dto := employees.EmployeeDto{
// 			CompanyId:  1,
// 			FirstName:  "lorem",
// 			LastName:   "ipsum",
// 			Email:      "loremipsum@google.com",
// 			BirthDay:   15,
// 			BirthMonth: 4,
// 		}
// 		body, err := json.Marshal(dto)
// 		require.Nil(t, err)
//
// 		req, err := http.NewRequest(http.MethodPost, employeeGeneralUrl, bytes.NewBuffer(body))
// 		require.Nil(t, err)
//
// 		resp, err := http.DefaultClient.Do(req)
// 		require.Nil(t, err)
//
// 		assert.Equal(t, http.StatusCreated, resp.StatusCode)
// 	})
//
// 	t.Run("it should mount employeeListHandler", func(t *testing.T) {
// 		resp, err := http.Get(employeeGeneralUrl)
//
// 		require.Nil(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	})
//
// 	t.Run("it should mount employeeDetailHandler", func(t *testing.T) {
// 		resp, err := http.Get(employeeDetailUrl)
//
// 		require.Nil(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	})
// }
