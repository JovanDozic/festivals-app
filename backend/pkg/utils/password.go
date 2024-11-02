package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) bool {
	passwordBytes := []byte(password)

	hashedPasswordBytes := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
	return err == nil
}
