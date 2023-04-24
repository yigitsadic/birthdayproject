package sessions

type AuthenticationModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserId       int    `json:"user_id"`
	CompanyId    int    `json:"company_id"`
}
