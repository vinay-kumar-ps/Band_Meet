package helper

import (
	"time"

	"github.com/Anandhu4456/band-meet/pkg/utils/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHashing(pass string) (string, error) {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}
	return string(bytePass), nil
}

type CustomAuthClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateBandUserToken(bandUser models.BandLoginResponse) (string, string, error) {
	tokenClaims := &CustomAuthClaims{
		Id:    bandUser.ID,
		Email: bandUser.Email,
		Role:  "band_user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	refreshTokenClaims := &CustomAuthClaims{
		Id:    bandUser.ID,
		Email: bandUser.Email,
		Role:  "band_user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte("bandusersecret"))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshtokenString, err := refreshToken.SignedString([]byte("bandusersecret"))
	if err != nil {
		return "", "", err
	}
	return tokenString, refreshtokenString, nil
}
