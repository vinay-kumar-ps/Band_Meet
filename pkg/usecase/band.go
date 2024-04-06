package usecase

import (
	"github.com/Anandhu4456/band-meet/pkg/repository/interfaces"
	services "github.com/Anandhu4456/band-meet/pkg/usecase/interfaces"
	"github.com/Anandhu4456/band-meet/pkg/utils/models"
)

type BandUsecase struct {
	BandRepo interfaces.BandRepo
}

func NewBandUsecase(bandRepo interfaces.BandRepo)services.BandUsecase{
	return &BandUsecase{
		BandRepo: bandRepo,
	}
}

func (b *BandUsecase)BandUserSignup(user models.BandSignup)(models.BandToken,error){
	
}