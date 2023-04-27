package server

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/yigitsadic/birthday-app-api/internal/common"
)

func injectParamKeyToChiCtx(
	t *testing.T,
	r *http.Request,
	keys []string,
	vals []string,
) *http.Request {
	t.Helper()

	// Add param to context. chi reads /aa/12 from chi Ctx
	chiCtx := chi.Context{
		URLParams: chi.RouteParams{
			Keys:   keys,
			Values: vals,
		},
	}

	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, &chiCtx))
}

func loginWithRequestLevel(
	t *testing.T,
	r *http.Request,
	userId string,
	companyId string,
) *http.Request {
	t.Helper()

	r1 := r.WithContext(context.WithValue(r.Context(), common.UserIdCtxKey, userId))
	r2 := r.WithContext(context.WithValue(r1.Context(), common.CompanyIdCtxKey, companyId))

	return r2
}
