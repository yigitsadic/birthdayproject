package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_paramConverter(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := paramConverter("id", w, r)

		w.Write([]byte(fmt.Sprintf("Given id was: %d", id)))
	})

	srv := httptest.NewServer(r)

	defer srv.Close()

	t.Run("it should render not found if not integer", func(t *testing.T) {
		resp, err := http.Get(srv.URL + "/legolas")
		require.Nil(t, err)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("it should render not found if not positive", func(t *testing.T) {
		resp, err := http.Get(srv.URL + "/-1")
		require.Nil(t, err)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("it should do nothing if correct integer id", func(t *testing.T) {
		resp, err := http.Get(srv.URL + "/34")
		require.Nil(t, err)

		defer resp.Body.Close()

		cont, err := io.ReadAll(resp.Body)
		require.Nil(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "Given id was: 34", string(cont))
	})

}
