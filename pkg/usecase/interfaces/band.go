package interfaces

import "github.com/Anandhu4456/band-meet/pkg/utils/models"

type BandUsecase interface {
	BandUserSignup(user models.BandSignup) (models.BandToken, error)
	BandUserLogin(user models.BandLogin) (models.BandToken, error)
}
