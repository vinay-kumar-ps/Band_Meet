package handlers

import (
	"net/http"

	"github.com/Anandhu4456/band-meet/pkg/usecase/interfaces"
	"github.com/Anandhu4456/band-meet/pkg/utils/models"
	"github.com/Anandhu4456/band-meet/pkg/utils/response"
	"github.com/gofiber/fiber/v2"
)

type BandHandler struct {
	BandUsecase interfaces.BandUsecase
}

func NewBandHandler(bandUsecase interfaces.BandUsecase) *BandHandler {
	return &BandHandler{
		BandUsecase: bandUsecase,
	}
}

func (b *BandHandler) BandSignupHandler(c *fiber.Ctx) error {
	var bandUserSignup models.BandSignup
	if err := c.BodyParser(&bandUserSignup); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errRes)
	}
	signupDetails, err := b.BandUsecase.BandUserSignup(bandUserSignup)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "signup failed", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errRes)
	}
	successRes := response.ClientResponse(http.StatusOK, "signed up successfully", signupDetails, nil)
	c.Status(http.StatusOK).JSON(successRes)
	return nil
}

func (b *BandHandler) BandLoginHandler(c *fiber.Ctx) error {
	var bandLoginDetails models.BandLogin
	if err := c.BodyParser(&bandLoginDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errRes)
	}
	loginDetails, err := b.BandUsecase.BandUserLogin(bandLoginDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "login failed", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errRes)
	}
	successRes := response.ClientResponse(http.StatusOK, "login successfully", loginDetails, nil)
	c.Status(http.StatusOK).JSON(successRes)
	return nil
}
