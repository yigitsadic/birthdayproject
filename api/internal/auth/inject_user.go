package auth

import (
	"context"
	"net/http"

	"github.com/yigitsadic/birthday-app-api/internal/common"
)

// InjectUser injects user_id and company_id from JWT to request context.
func InjectUser(jwtStore JWTStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, err := GetJWTFromHeader(r, jwtStore)

			if err == nil {
				userId := claims.Subject
				companyId := claims.CompanyId

				r1 := r.WithContext(context.WithValue(r.Context(), common.UserIdCtxKey, userId))
				r2 := r1.WithContext(
					context.WithValue(r1.Context(), common.CompanyIdCtxKey, companyId),
				)

				next.ServeHTTP(w, r2)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
