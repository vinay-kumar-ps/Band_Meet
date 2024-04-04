package models

type BandLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BandSignup struct {
	Name            string `json:"name"`
	UserName        string `json:"username"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type BandLoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type BandToken struct {
	Band         BandLoginResponse
	Token        string
	RefreshToken string
}
