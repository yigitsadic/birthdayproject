package responses

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func handleTestServer(t *testing.T, renderFunc func(http.ResponseWriter)) (ErrorMessage, int) {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderFunc(w)
	}))

	resp, err := http.DefaultClient.Get(srv.URL)
	defer srv.Close()

	require.Nil(t, err)

	parsed := ErrorMessage{}
	require.Nil(t, json.NewDecoder(resp.Body).Decode(&parsed))

	defer resp.Body.Close()

	return parsed, resp.StatusCode
}

func handleTestServerWithMessage(t *testing.T, renderFunc func(http.ResponseWriter, string), msg string) (ErrorMessage, int) {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderFunc(w, msg)
	}))

	resp, err := http.DefaultClient.Get(srv.URL)
	defer srv.Close()

	require.Nil(t, err)

	parsed := ErrorMessage{}
	require.Nil(t, json.NewDecoder(resp.Body).Decode(&parsed))

	defer resp.Body.Close()

	return parsed, resp.StatusCode
}

func TestRenderUnauthenticated(t *testing.T) {
	parsed, statusCode := handleTestServerWithMessage(t, RenderUnauthenticated, UnauthenticatedMessage)

	assert.Equal(t, http.StatusUnauthorized, statusCode)
	assert.Equal(t, UnauthenticatedMessage, parsed.Message)
}

func TestRenderNotFound(t *testing.T) {
	parsed, statusCode := handleTestServer(t, RenderNotFound)

	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Equal(t, notFoundMessage, parsed.Message)
}

func TestRenderInternalServerError(t *testing.T) {
	parsed, statusCode := handleTestServer(t, RenderInternalServerError)

	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Equal(t, internalServerErrorMessage, parsed.Message)
}

func TestRenderUnprocessableEntity(t *testing.T) {
	msg := "company name is not valid"

	parsed, statusCode := handleTestServerWithMessage(t, RenderUnprocessableEntity, msg)
	assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
	assert.Equal(t, msg, parsed.Message)
}
