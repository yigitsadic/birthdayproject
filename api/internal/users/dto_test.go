package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserDtoValidations(t *testing.T) {
	assert := assert.New(t)

	t.Run("it should return nil when valid", func(t *testing.T) {
		u := UserDto{
			FirstName: "John",
			LastName:  "Doe",
		}

		assert.Nil(u.Validate())
	})

	t.Run("it should return an error when first name is invalid", func(t *testing.T) {
		u := UserDto{
			FirstName: "J",
			LastName:  "Doe",
		}

		assert.NotNil(u.Validate())
	})

	t.Run("it should return an error when last name is invalid", func(t *testing.T) {
		u := UserDto{
			FirstName: "John",
			LastName:  "D",
		}

		assert.NotNil(u.Validate())
	})
}
