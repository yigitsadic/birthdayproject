package sessions

import "context"

type SessionRepository interface {
	FindUser(context.Context, string) (*SessionUser, error)
	FindById(context.Context, int) (*SessionUser, error)
}
