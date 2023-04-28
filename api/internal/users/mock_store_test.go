package users

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MockUserStore(t *testing.T) {
	store := NewMockUserStore()
	assert := assert.New(t)
	firstRecord := store.Store[0]

	resetError := func() {
		store.RaiseErrorOnGetUser = false
		store.RaiseErrorOnUpdate = false
	}

	t.Run("it should return user", func(t *testing.T) {
		resetError()

		got, err := store.GetUser(context.TODO(), 4)

		assert.Nil(err)
		assert.Equal(firstRecord, got)
	})

	t.Run("it should return an error on GetUser", func(t *testing.T) {
		resetError()

		store.RaiseErrorOnGetUser = true
		got, err := store.GetUser(context.TODO(), 4)

		assert.NotNil(err)
		assert.Nil(got)
	})

	t.Run("it should update user", func(t *testing.T) {
		resetError()

		got, err := store.UpdateUser(context.TODO(), 4, UserDto{
			FirstName: "John",
			LastName:  "Doe",
		})

		assert.Nil(err)
		assert.Equal("John", got.FirstName)
		assert.Equal("Doe", got.LastName)
	})

	t.Run("it should return an error on UpdateUser", func(t *testing.T) {
		resetError()

		store.RaiseErrorOnUpdate = true

		got, err := store.UpdateUser(context.TODO(), 4, UserDto{
			FirstName: "John",
			LastName:  "Doe",
		})

		assert.NotNil(err)
		assert.Nil(got)
	})
}
