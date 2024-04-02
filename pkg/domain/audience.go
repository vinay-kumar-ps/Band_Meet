package domain

type User struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone" gorm:"unique"`
	Password string `json:"password"`
}
