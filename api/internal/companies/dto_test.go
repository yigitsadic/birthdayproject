package companies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CompanyUpdateDtoValidations(t *testing.T) {
	t.Run("it should return nil when good", func(t *testing.T) {
		c := CompanyUpdateDto{
			Name: "A company",
		}

		assert.Nil(t, c.Validate())
	})

	t.Run("it should return an error when invalid", func(t *testing.T) {
		a := CompanyUpdateDto{
			Name: "AB",
		}

		assert.NotNil(t, a.Validate())
	})
}
