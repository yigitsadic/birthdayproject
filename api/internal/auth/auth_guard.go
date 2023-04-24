package auth

import (
	"net/http"

	"github.com/yigitsadic/birthday-app-api/internal/common"
	"github.com/yigitsadic/birthday-app-api/internal/responses"
)

// AuthGuard guards route. Intercepts unauthorized requests.
// Looks for user_id and company_id on request's context.
func AuthGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, userOK := r.Context().Value(common.UserIdCtxKey).(string)
		companyId, companyOK := r.Context().Value(common.CompanyIdCtxKey).(string)

		if userOK && companyOK && userId != "" && companyId != "" {
			next.ServeHTTP(w, r)
		} else {
			responses.RenderUnauthenticated(w, responses.UnauthenticatedMessage)
		}
	})
}
