package sessions

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SessionDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s SessionDto) Validate() error {
	return validation.ValidateStruct(
		&s,
		validation.Field(&s.Email, validation.Required, is.Email, validation.Length(6, 50)),
		validation.Field(&s.Password, validation.Required, validation.Length(8, 21)),
	)
}
