package employees

type EmployeeModel struct {
	ID         int    `json:"id"`
	CompanyId  int    `json:"company_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	BirthDay   int    `json:"birth_day"`
	BirthMonth int    `json:"birth_month"`
}
