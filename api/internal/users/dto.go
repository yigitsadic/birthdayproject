package users

import validation "github.com/go-ozzo/ozzo-validation"

type UserDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u UserDto) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.FirstName, validation.Required, validation.Length(2, 40)),
		validation.Field(&u.LastName, validation.Required, validation.Length(2, 40)),
	)
}
