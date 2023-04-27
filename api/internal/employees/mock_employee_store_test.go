package employees

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MockEmployeeStore(t *testing.T) {
	assert := assert.New(t)
	store := NewMockEmployeeStore()

	assert.NotNil(store)
	obj := store.Store[0]

	resetError := func() {
		store.RaiseErrorOnOperation = false
	}

	t.Run("FetchAll: it should return first item", func(t *testing.T) {
		defer resetError()
		got, err := store.FetchAll(context.TODO(), 3)

		assert.Nil(err)
		assert.Equal(obj.FirstName, got[0].FirstName)
	})

	t.Run("FetchAll: it should an error if configured", func(t *testing.T) {
		defer resetError()

		store.RaiseErrorOnOperation = true
		got, err := store.FetchAll(context.TODO(), 3)

		assert.NotNil(err)
		assert.Nil(got)
		assert.Equal(mockStoreErr, err)
	})

	t.Run("FindOne: it should return first item", func(t *testing.T) {
		defer resetError()
		got, err := store.FindOne(context.TODO(), 3, 5)

		assert.Nil(err)
		assert.Equal(obj.FirstName, got.FirstName)
	})

	t.Run("FindOne: it should an error if configured", func(t *testing.T) {
		defer resetError()

		store.RaiseErrorOnOperation = true
		got, err := store.FindOne(context.TODO(), 3, 5)

		assert.NotNil(err)
		assert.Nil(got)
		assert.Equal(mockStoreErr, err)
	})

	t.Run("Create: it should return first item", func(t *testing.T) {
		defer resetError()
		got, err := store.Create(context.TODO(), 3, EmployeeDto{})

		assert.Nil(err)
		assert.Equal(obj.FirstName, got.FirstName)
	})

	t.Run("Create: it should an error if configured", func(t *testing.T) {
		defer resetError()

		store.RaiseErrorOnOperation = true
		got, err := store.Create(context.TODO(), 3, EmployeeDto{})

		assert.NotNil(err)
		assert.Nil(got)
		assert.Equal(mockStoreErr, err)
	})

	t.Run("Update: it should return first item", func(t *testing.T) {
		defer resetError()
		got, err := store.Update(context.TODO(), 5, 3, EmployeeDto{})

		assert.Nil(err)
		assert.Equal(obj.FirstName, got.FirstName)
	})

	t.Run("Update: it should an error if configured", func(t *testing.T) {
		defer resetError()
		store.RaiseErrorOnOperation = true

		got, err := store.Update(context.TODO(), 5, 3, EmployeeDto{})

		assert.NotNil(err)
		assert.Nil(got)
		assert.Equal(mockStoreErr, err)
	})

	t.Run("Delete: it should return first item", func(t *testing.T) {
		defer resetError()
		err := store.Delete(context.TODO(), 5, 3)

		assert.Nil(err)
	})

	t.Run("Delete: it should an error if configured", func(t *testing.T) {
		defer resetError()
		store.RaiseErrorOnOperation = true
		err := store.Delete(context.TODO(), 5, 3)

		assert.NotNil(err)
		assert.Equal(mockStoreErr, err)
	})
}
