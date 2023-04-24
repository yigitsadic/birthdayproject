package employees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EmployeeDtoValidations(t *testing.T) {
	t.Run("it should return nil if everything is valid", func(t *testing.T) {
		d := EmployeeDto{
			CompanyId:  1,
			FirstName:  "Yigit",
			LastName:   "Sadic",
			Email:      "yigit@google.com",
			BirthDay:   5,
			BirthMonth: 5,
		}
		assert.Nil(t, d.Validate())
	})

	t.Run("it should check standard validations", func(t *testing.T) {
		d := EmployeeDto{
			CompanyId:  1,
			FirstName:  "Yigit",
			LastName:   "S",
			Email:      "yigit",
			BirthDay:   5,
			BirthMonth: 5,
		}
		assert.NotNil(t, d.Validate())
	})

	t.Run("test max day of months", func(t *testing.T) {
		tests := []struct {
			name       string
			givenDay   int
			givenMonth int
		}{
			{
				name:       "02/29 should be invalid",
				givenDay:   29,
				givenMonth: 2,
			},
			{
				name:       "04/31 should be invalid",
				givenDay:   31,
				givenMonth: 4,
			},
			{
				name:       "06/31 should be invalid",
				givenDay:   31,
				givenMonth: 6,
			},
			{
				name:       "09/31 should be invalid",
				givenDay:   31,
				givenMonth: 9,
			},
			{
				name:       "011/31 should be invalid",
				givenDay:   31,
				givenMonth: 11,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				d := EmployeeDto{
					CompanyId:  1,
					FirstName:  "Yigit",
					LastName:   "Sadic",
					Email:      "yigit@google.com",
					BirthDay:   tt.givenDay,
					BirthMonth: tt.givenMonth,
				}
				err := d.Validate()

				assert.Equal(t, ErrEmployeeBirthDayMonth, err)
			})
		}
	})
}
