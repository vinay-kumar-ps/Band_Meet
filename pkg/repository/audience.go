package repository

import (
	interfaces "github.com/Anandhu4456/band-meet/pkg/repository/interfaces"
	"github.com/Anandhu4456/band-meet/pkg/utils/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserRepo{
		DB: db,
	}
}
func (u *UserRepo) UserSignup(userSignup models.UserSignup) (models.UserLoginResponse, error) {
	var userLoginResponse models.UserLoginResponse
	err := u.DB.Exec("INSERT INTO users(username,name,email,phone,password)Values(?,?,?,?,?)RETURNING(username,id,email,phone)", userSignup.UserName, userSignup.Name, userSignup.Email, userSignup.Phone, userSignup.Password).Scan(&userLoginResponse).Error
	if err != nil {
		return models.UserLoginResponse{}, err
	}
	return userLoginResponse, nil
}
