package interfaces

import (


	"github.com/Anandhu4456/band-meet/pkg/utils/models"
)

type UserRepo interface{
	UserSignup( userSignup models.UserSignup) (models.UserLoginResponse,error)
}