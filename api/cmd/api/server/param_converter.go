package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yigitsadic/birthday-app-api/internal/responses"
)

// paramConverter converts given url param to integer. Used for id convertions.
func paramConverter(field string, w http.ResponseWriter, r *http.Request) int {
	idAsStr := chi.URLParam(r, field)

	id, err := strconv.Atoi(idAsStr)

	if err != nil {
		responses.RenderNotFound(w)

		return 0
	}

	if id < 1 {
		responses.RenderNotFound(w)

		return 0
	}

	return id
}
