package interfaces

import "github.com/Anandhu4456/band-meet/pkg/utils/models"

type UserUsecase interface{
	UserLogin(user models.UserLogin) (models.UserToken,error)
	UserSignup(user models.UserSignup) (models.UserToken,error)
} 