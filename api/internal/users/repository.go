package users

import "context"

type UserRepository interface {
	GetUser(context.Context, int) (*UserModel, error)
	UpdateUser(context.Context, int, UserDto) (*UserModel, error)
}
