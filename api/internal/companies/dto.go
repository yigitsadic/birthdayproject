package companies

import validation "github.com/go-ozzo/ozzo-validation"

type CompanyUpdateDto struct {
	Name string `json:"name"`
}

func (c CompanyUpdateDto) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Name, validation.Required, validation.Length(3, 50)),
	)
}
