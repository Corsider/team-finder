package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Print("Error while generating hash")
	}
	return string(hashed)
}

func ValidatePassword(hashedPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

func First[T, U any](val T, _ U) T {
	return val
}

func Second[T, U any](_ T, val U) U {
	return val
}
