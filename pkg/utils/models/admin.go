package models

type AdminLoginDetails struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
}
type AdminLoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdminToken struct {
	Admin        AdminLoginResponse
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}
