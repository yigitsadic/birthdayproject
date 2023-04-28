package companies

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MockCompanyStore(t *testing.T) {
	s := NewMockCompanyStore()
	assert := assert.New(t)
	firstRecord := s.Store[0]

	resetNormal := func() {
		s.RaiseErrorOnFind = false
		s.RaiseErrorOnUpdate = false
	}

	t.Run("it should return company", func(t *testing.T) {
		resetNormal()
		got, err := s.FetchOne(context.TODO(), 1)

		assert.Nil(err)
		assert.Equal(firstRecord.Name, got.Name)
		assert.Equal(firstRecord.ID, got.ID)
	})

	t.Run("it should return an error when wanted", func(t *testing.T) {
		resetNormal()
		s.RaiseErrorOnFind = true
		got, err := s.FetchOne(context.TODO(), 1)

		assert.NotNil(err)
		assert.Nil(got)
	})

	t.Run("it should update company", func(t *testing.T) {
		resetNormal()
		got, err := s.Update(context.TODO(), 1, CompanyUpdateDto{})

		assert.Nil(err)
		assert.Equal(firstRecord.Name, got.Name)
		assert.Equal(firstRecord.ID, got.ID)
	})

	t.Run("it should return an error when wanted for update", func(t *testing.T) {
		resetNormal()
		s.RaiseErrorOnUpdate = true
		got, err := s.Update(context.TODO(), 1, CompanyUpdateDto{})

		assert.NotNil(err)
		assert.Nil(got)
	})
}
