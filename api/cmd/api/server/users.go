package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yigitsadic/birthday-app-api/internal/common"
	"github.com/yigitsadic/birthday-app-api/internal/responses"
	"github.com/yigitsadic/birthday-app-api/internal/users"
)

// HandleUserDetail Serves user's details.
func (s *Server) HandleUserDetail(w http.ResponseWriter, r *http.Request) {
	id := paramConverter("id", w, r)

	userId, _ := r.Context().Value(common.UserIdCtxKey).(string)
	if userId != strconv.Itoa(id) {
		responses.RenderNotFound(w)

		return
	}

	result, err := s.UserRepository.GetUser(r.Context(), id)
	if err != nil {
		responses.RenderNotFound(w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// HandleUserUpdate Updates users's informations.
func (s *Server) HandleUserUpdate(w http.ResponseWriter, r *http.Request) {
	id := paramConverter("id", w, r)
	userId, _ := r.Context().Value(common.UserIdCtxKey).(string)
	if userId != strconv.Itoa(id) {
		responses.RenderNotFound(w)
		return
	}

	var dto users.UserDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		responses.RenderUnprocessableEntity(w, err.Error())
		return
	}

	if err := dto.Validate(); err != nil {
		responses.RenderUnprocessableEntity(w, err.Error())
		return
	}

	result, err := s.UserRepository.UpdateUser(r.Context(), id, dto)
	if err != nil {
		responses.RenderUnprocessableEntity(w, "unable to update user")
		return
	}

	json.NewEncoder(w).Encode(result)
}
