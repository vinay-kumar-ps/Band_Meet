package domain

type Admin struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Name     string `json:"name" gorm:"validate:required"`
	Email    string `json:"email" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
}
