package dtos

type CurrentUser struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	AllClaims []string `json:"all_claims"`
	Roles     []string `json:"roles"`
}
