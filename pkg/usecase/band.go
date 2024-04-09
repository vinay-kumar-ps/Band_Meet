package usecase

import (
	"errors"

	"github.com/Anandhu4456/band-meet/pkg/helper"
	"github.com/Anandhu4456/band-meet/pkg/repository/interfaces"
	services "github.com/Anandhu4456/band-meet/pkg/usecase/interfaces"
	"github.com/Anandhu4456/band-meet/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

type BandUsecase struct {
	BandRepo interfaces.BandRepo
}

func NewBandUsecase(bandRepo interfaces.BandRepo) services.BandUsecase {
	return &BandUsecase{
		BandRepo: bandRepo,
	}
}

func (b *BandUsecase) BandUserSignup(user models.BandSignup) (models.BandToken, error) {
	isAlready := b.BandRepo.BandUserAvailability(user.Email)
	if isAlready {
		return models.BandToken{}, errors.New("band user already exist..please login")
	}
	if user.Password != user.ConfirmPassword {
		return models.BandToken{}, errors.New("password doesnt match")
	}
	// password hashing
	hashedPass, err := helper.PasswordHashing(user.Password)
	if err != nil {
		return models.BandToken{}, err
	}
	// set the hashed pass
	user.Password = hashedPass

	bandLoginResponse, err := b.BandRepo.BandSignup(user)
	if err != nil {
		return models.BandToken{}, err
	}
	// generate token for band user
	token, refreshtoken, err := helper.GenerateBandUserToken(bandLoginResponse)
	if err != nil {
		return models.BandToken{}, err
	}
	return models.BandToken{
		Band:         bandLoginResponse,
		Token:        token,
		RefreshToken: refreshtoken,
	}, nil
}

func (b *BandUsecase) BandUserLogin(user models.BandLogin) (models.BandToken, error) {
	// check band user already signed or not
	ok := b.BandRepo.BandUserAvailability(user.Email)
	if !ok {
		return models.BandToken{}, errors.New("band user not exist..please signup first")
	}
	// find the band user details to check the password
	bandUserDetails, err := b.BandRepo.BandUserDetails(user.Email)
	if err != nil {
		return models.BandToken{}, err
	}
	// comparing the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(bandUserDetails.Password)); err != nil {
		return models.BandToken{}, err
	}
	var bandLoginResponse models.BandLoginResponse
	bandLoginResponse.ID = bandUserDetails.Id
	bandLoginResponse.Username = bandUserDetails.UserName
	bandLoginResponse.Email = bandUserDetails.Email
	bandLoginResponse.Phone = bandUserDetails.Phone

	token, refreshToken, err := helper.GenerateBandUserToken(bandLoginResponse)
	if err != nil {
		return models.BandToken{}, err
	}
	return models.BandToken{
		Band:         bandLoginResponse,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
