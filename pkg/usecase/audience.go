package usecase

import (
	"errors"

	"github.com/Anandhu4456/band-meet/pkg/helper"
	"github.com/Anandhu4456/band-meet/pkg/repository/interfaces"
	services "github.com/Anandhu4456/band-meet/pkg/usecase/interfaces"
	"github.com/Anandhu4456/band-meet/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepo interfaces.UserRepo
}

func NewUserUsercase(userRepo interfaces.UserRepo) services.UserUsecase {
	return &UserUsecase{
		UserRepo: userRepo,
	}
}
func (u *UserUsecase) UserSignup(user models.UserSignup) (models.UserToken, error) {
	isAlready := u.UserRepo.UserAvailability(user.Email)
	if isAlready {
		return models.UserToken{}, errors.New("user is already exist ..please try to login")
	}
	if user.Password != user.ConfirmPassword {
		return models.UserToken{}, errors.New("password doesnt match")
	}
	//password hashing
	hashedpass, err := helper.PasswordHashing(user.Password)
	if err != nil {
		return models.UserToken{}, err
	}
	//set the hashed pass
	user.Password = hashedpass

	userLoginResponse, err := u.UserRepo.UserSignup(user)

	if err != nil {
		return models.UserToken{}, err

	}
	//generate token for user
	token, refreshToken, err := helper.GenerateUserToken(userLoginResponse)
	if err != nil {
		return models.UserToken{}, err
	}
	return models.UserToken{
		Token:        token,
		RefreshToken: refreshToken,
		User:         userLoginResponse,
	}, nil

}
func (u *UserUsecase) UserLogin(user models.UserLogin) (models.UserToken, error) {
	//check the user is already signed or not
	ok := u.UserRepo.UserAvailability(user.Email)
	if !ok {
		return models.UserToken{}, errors.New(" user not exist.. please signup first")

	}
	//find the user details to check the password
	userDetails, err := u.UserRepo.UserDetails(user.Email)
	if err != nil {
		return models.UserToken{}, err
	}
	//comparing the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password)); err != nil {
		return models.UserToken{}, err

	}
	var userLoginResponse models.UserLoginResponse
	userLoginResponse.ID = userDetails.Id
	userLoginResponse.Username = userDetails.UserName
	userLoginResponse.Email = userDetails.Email
	userLoginResponse.Phone = userDetails.Phone

	token, refreshToken, err := helper.GenerateUserToken(userLoginResponse)
	if err != nil {
		return models.UserToken{}, err
	}
	return models.UserToken{
		Token:        token,
		RefreshToken: refreshToken,
		User:         userLoginResponse,
	}, nil
}
