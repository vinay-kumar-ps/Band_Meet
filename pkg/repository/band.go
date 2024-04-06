package repository

import (
	"github.com/Anandhu4456/band-meet/pkg/repository/interfaces"
	"github.com/Anandhu4456/band-meet/pkg/utils/models"
	"gorm.io/gorm"
)

type BandRepo struct {
	DB *gorm.DB
}

func NewBandRepo(db *gorm.DB) interfaces.BandRepo {
	return &BandRepo{
		DB: db,
	}
}

func (b *BandRepo) BandSignup(bandSignup models.BandSignup) (models.BandLoginResponse, error) {
	var bandLoginResponse models.BandLoginResponse
	err := b.DB.Exec("INSERT INTO band_profile(name,username,email,phone,password)VALUES(?,?,?,?,?)RETURNING(id,username,email,phone)", bandSignup.Name, bandSignup.Email, bandSignup.Phone, bandSignup.Password).Scan(&bandLoginResponse).Error
	if err != nil {
		return models.BandLoginResponse{}, err
	}
	return bandLoginResponse, nil
}

func (b *BandRepo) BandUserAvailability(email string) bool {
	var bandUserCount int
	err := b.DB.Raw("SELECT COUNT(*)FROM band_profile WHERE email=?", email).Scan(&bandUserCount).Error
	if err != nil {
		return false
	}
	return bandUserCount > 0
}

func (b *BandRepo) BandUserDetails(email string) (models.BandUserSignupResponse, error) {
	var bandUsersignupResponse models.BandUserSignupResponse
	err := b.DB.Raw("SELECT *FROM band_profile WHERE email=?", email).Scan(&bandUsersignupResponse).Error
	if err != nil {
		return models.BandUserSignupResponse{}, err
	}
	return bandUsersignupResponse, nil
}
