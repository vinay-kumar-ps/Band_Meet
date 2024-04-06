package helper

import "golang.org/x/crypto/bcrypt"

func PasswordHashing(pass string) (string, error) {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}
	return string(bytePass), nil
}
