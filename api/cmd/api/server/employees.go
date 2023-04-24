package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yigitsadic/birthday-app-api/internal/common"
	"github.com/yigitsadic/birthday-app-api/internal/employees"
	"github.com/yigitsadic/birthday-app-api/internal/responses"
)

// EmployeeUpdateHandler updates given employee for authenticated company.
func (s *Server) EmployeeUpdateHandler(w http.ResponseWriter, r *http.Request) {
	companyId := paramConverter("companyId", w, r)
	companyIdFromCtx, _ := r.Context().Value(common.CompanyIdCtxKey).(string)
	if companyIdFromCtx != strconv.Itoa(companyId) {
		responses.RenderNotFound(w)

		return
	}
	employeeId := paramConverter("id", w, r)

	var dto employees.EmployeeDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	defer r.Body.Close()
	if err != nil {
		responses.RenderUnprocessableEntity(w, "unable to parse body")
		return
	}

	if err = dto.Validate(); err != nil {
		responses.RenderUnprocessableEntity(w, err.Error())
		return
	}

	result, err := s.EmployeeRepository.Update(r.Context(), companyId, employeeId, dto)
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// EmployeeDeleteHandler deletes given employee for authenticated company.
func (s *Server) EmployeeDeleteHandler(w http.ResponseWriter, r *http.Request) {
	companyId := paramConverter("companyId", w, r)
	companyIdFromCtx, _ := r.Context().Value(common.CompanyIdCtxKey).(string)
	if companyIdFromCtx != strconv.Itoa(companyId) {
		responses.RenderNotFound(w)

		return
	}
	employeeId := paramConverter("id", w, r)

	err := s.EmployeeRepository.Delete(r.Context(), companyId, employeeId)
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// EmployeeDetailHandler serves detail of an employee for authenticated company.
func (s *Server) EmployeeDetailHandler(w http.ResponseWriter, r *http.Request) {
	companyId := paramConverter("companyId", w, r)
	companyIdFromCtx, _ := r.Context().Value(common.CompanyIdCtxKey).(string)
	if companyIdFromCtx != strconv.Itoa(companyId) {
		responses.RenderNotFound(w)

		return
	}
	employeeId := paramConverter("id", w, r)

	result, err := s.EmployeeRepository.FindOne(r.Context(), companyId, employeeId)
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// EmployeeCreateHandler Creates an employee for authenticated company.
func (s *Server) EmployeeCreateHandler(w http.ResponseWriter, r *http.Request) {
	id := paramConverter("companyId", w, r)
	companyId, _ := r.Context().Value(common.CompanyIdCtxKey).(string)
	if companyId != strconv.Itoa(id) {
		responses.RenderNotFound(w)

		return
	}

	var dto employees.EmployeeDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	defer r.Body.Close()
	if err != nil {
		responses.RenderUnprocessableEntity(w, "unable to parse body")
		return
	}

	dto.CompanyId = id

	if err = dto.Validate(); err != nil {
		responses.RenderUnprocessableEntity(w, err.Error())
		return
	}

	result, err := s.EmployeeRepository.Create(r.Context(), id, dto)
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// EmployeeListHandler Serves list of employees of authenticated company.
func (s *Server) EmployeeListHandler(w http.ResponseWriter, r *http.Request) {
	id := paramConverter("companyId", w, r)
	companyId, _ := r.Context().Value(common.CompanyIdCtxKey).(string)
	if companyId != strconv.Itoa(id) {
		responses.RenderNotFound(w)

		return
	}

	result, err := s.EmployeeRepository.FetchAll(r.Context(), id)
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	json.NewEncoder(w).Encode(&result)
}
