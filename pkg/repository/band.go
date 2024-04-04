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
