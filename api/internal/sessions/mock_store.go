package sessions

import (
	"context"
	"errors"

	"github.com/yigitsadic/birthday-app-api/internal/auth"
)

type MockSessionStore struct {
	Store         []*SessionUser
	RaiseNotFound bool
}

// Ensure MockSessionStore follows SessionRepository
var _ SessionRepository = (*MockSessionStore)(nil)

func NewMockSessionStore() MockSessionStore {
	pass1, _ := auth.HashPassword("123456789")
	pass2, _ := auth.HashPassword("987654321")

	s := MockSessionStore{}
	s.Store = []*SessionUser{
		{
			ID:           1,
			Email:        "yigit@google.com",
			PasswordHash: pass1,
			CompanyId:    1,
		},
		{
			ID:           2,
			Email:        "johndoe@google.com",
			PasswordHash: pass2,
			CompanyId:    2,
		},
	}

	return s
}

func (m *MockSessionStore) FindUser(context.Context, string) (*SessionUser, error) {
	if m.RaiseNotFound {
		return nil, errors.New("an error occurred")
	}

	return m.Store[0], nil
}

func (m *MockSessionStore) FindById(context.Context, int) (*SessionUser, error) {
	if m.RaiseNotFound {
		return nil, errors.New("an error occurred")
	}

	return m.Store[0], nil
}
