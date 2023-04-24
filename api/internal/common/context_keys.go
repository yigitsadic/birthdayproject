package common

var (
	UserIdCtxKey    = &contextKey{"UserIdCtxKey"}
	CompanyIdCtxKey = &contextKey{"CompanyIdCtxKey"}
)

type contextKey struct {
	name string
}
