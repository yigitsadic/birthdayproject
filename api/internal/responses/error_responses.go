package responses

import (
	"encoding/json"
	"net/http"
)

var (
	UnauthenticatedMessage     = "Invalid/expired token."
	notFoundMessage            = "Resource is not found."
	internalServerErrorMessage = "Internal server error occurred."
)

// RenderUnauthenticated renders unauthorized response
func RenderUnauthenticated(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(NewErrorMessage(msg))
}

// RenderNotFound renders not found response.
func RenderNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(NewErrorMessage(notFoundMessage))
}

// RenderInternalServerError renders internal server error response.
func RenderInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(NewErrorMessage(internalServerErrorMessage))
}

// RenderUnprocessableEntity renders error message for 422 responses.
func RenderUnprocessableEntity(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(NewErrorMessage(message))
}
