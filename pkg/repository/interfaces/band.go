package interfaces

import "github.com/Anandhu4456/band-meet/pkg/utils/models"

type BandRepo interface {
	BandSignup(bandSignup models.BandSignup) (models.BandLoginResponse, error)
	BandUserAvailability(email string) bool
}
