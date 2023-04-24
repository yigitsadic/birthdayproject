package users

import (
	"context"
	"errors"
)

type MockUserStore struct {
	Store               []*UserModel
	RaiseErrorOnUpdate  bool
	RaiseErrorOnGetUser bool
}

// Ensure MockUserStore follows UserRepository
var _ UserRepository = (*MockUserStore)(nil)

func NewMockUserStore() *MockUserStore {
	s := MockUserStore{}

	s.Store = append(s.Store, &UserModel{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@google.com",
	})
	s.Store = append(s.Store, &UserModel{
		ID:        2,
		FirstName: "Yigit",
		LastName:  "Sadic",
		Email:     "yigit.sadic@google.com",
	})

	return &s
}

func (m *MockUserStore) GetUser(context.Context, int) (*UserModel, error) {
	if m.RaiseErrorOnGetUser {
		return nil, errors.New("it is an error")
	}

	return m.Store[0], nil
}

func (m *MockUserStore) UpdateUser(_ context.Context, userId int, dto UserDto) (*UserModel, error) {
	if m.RaiseErrorOnUpdate {
		return nil, errors.New("an error")
	}

	m.Store[0].FirstName = dto.FirstName
	m.Store[0].LastName = dto.LastName

	return m.Store[0], nil
}
