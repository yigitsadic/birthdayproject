package employees

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	monthYearMap = []int{
		0,
		31,
		28,
		31,
		30,
		31,
		30,
		31,
		31,
		30,
		31,
		30,
		31,
	}
	ErrEmployeeBirthDayMonth = errors.New("invalid day for selected month")
)

type EmployeeDto struct {
	CompanyId  int    `json:"-"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	BirthDay   int    `json:"birth_day"`
	BirthMonth int    `json:"birth_month"`
}

func (dto EmployeeDto) Validate() error {
	validations := validation.ValidateStruct(&dto,
		validation.Field(&dto.FirstName, validation.Required, validation.Length(2, 30)),
		validation.Field(&dto.LastName, validation.Required, validation.Length(2, 30)),
		validation.Field(&dto.Email, validation.Required, validation.Length(5, 60), is.Email),
		validation.Field(&dto.BirthDay, validation.Required, validation.Min(1), validation.Max(31)),
		validation.Field(&dto.BirthMonth, validation.Required, validation.Min(1), validation.Max(12)),
	)
	if validations != nil {
		return validations
	}

	// check custom validations

	maxDay := monthYearMap[dto.BirthMonth]

	if maxDay >= dto.BirthDay {
		return nil
	}

	return ErrEmployeeBirthDayMonth
}
