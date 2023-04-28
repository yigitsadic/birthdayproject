package sessions

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MockSessionStore(t *testing.T) {
	assert := assert.New(t)

	store := NewMockSessionStore()
	firstRecord := store.Store[0]

	resetError := func() {
		store.RaiseNotFound = false
	}

	t.Run("it should return session user", func(t *testing.T) {
		resetError()

		got, err := store.FindUser(context.TODO(), "lorem")

		assert.Nil(err)
		assert.Equal(firstRecord, got)
	})

	t.Run("it should return an error when find user", func(t *testing.T) {
		resetError()
		store.RaiseNotFound = true

		got, err := store.FindUser(context.TODO(), "lorem")

		assert.NotNil(err)
		assert.Nil(got)
	})

	t.Run("it should return session user by id", func(t *testing.T) {
		resetError()

		got, err := store.FindById(context.TODO(), 5)

		assert.Nil(err)
		assert.Equal(firstRecord, got)
	})

	t.Run("it should return an error when find user by id", func(t *testing.T) {
		resetError()
		store.RaiseNotFound = true

		got, err := store.FindById(context.TODO(), 5)

		assert.NotNil(err)
		assert.Nil(got)
	})
}
