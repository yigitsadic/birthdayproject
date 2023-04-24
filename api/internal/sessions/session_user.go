package sessions

type SessionUser struct {
	ID           int
	Email        string
	PasswordHash string
	CompanyId    int
}
