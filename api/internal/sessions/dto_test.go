package sessions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SessionDtoValidations(t *testing.T) {
	assert := assert.New(t)

	t.Run("it should return nil when valid", func(t *testing.T) {
		s := SessionDto{
			Email:    "good@google.com",
			Password: "12345678",
		}

		assert.Nil(s.Validate())
	})

	t.Run("it should return an error when email is invalid", func(t *testing.T) {
		s := SessionDto{
			Email:    "goodgoogle.com",
			Password: "123456789",
		}

		assert.NotNil(s.Validate())

		s = SessionDto{
			Email:    "g@l.c",
			Password: "123456789",
		}
		assert.NotNil(s.Validate())
	})

	t.Run("it should return an error when password is invalid", func(t *testing.T) {
		s := SessionDto{
			Email:    "good@google.com",
			Password: "1234",
		}

		assert.NotNil(s.Validate())
	})
}
