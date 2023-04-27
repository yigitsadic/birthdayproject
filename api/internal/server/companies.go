package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yigitsadic/birthday-app-api/internal/common"
	"github.com/yigitsadic/birthday-app-api/internal/companies"
	"github.com/yigitsadic/birthday-app-api/internal/responses"
)

func (s *Server) HandleCompanyDetail(w http.ResponseWriter, r *http.Request) {
	id := paramConverter("id", w, r)
	companyId, _ := r.Context().Value(common.CompanyIdCtxKey).(string)
	if companyId != strconv.Itoa(id) {
		responses.RenderNotFound(w)

		return
	}

	result, err := s.CompanyRepository.FetchOne(r.Context(), id)
	if err != nil {
		responses.RenderNotFound(w)

		return
	}

	json.NewEncoder(w).Encode(result)
}

func (s *Server) HandleCompanyUpdate(w http.ResponseWriter, r *http.Request) {
	id := paramConverter("id", w, r)

	companyId, _ := r.Context().Value(common.CompanyIdCtxKey).(string)
	if companyId != strconv.Itoa(id) {
		responses.RenderNotFound(w)
		return
	}

	var dto companies.CompanyUpdateDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		responses.RenderUnprocessableEntity(w, err.Error())
		return
	}

	if err := dto.Validate(); err != nil {
		responses.RenderUnprocessableEntity(w, err.Error())
		return
	}

	result, err := s.CompanyRepository.Update(r.Context(), id, dto)
	if err != nil {
		responses.RenderUnprocessableEntity(w, "unable to update company")
		return
	}

	json.NewEncoder(w).Encode(result)
}
