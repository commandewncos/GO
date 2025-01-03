package util

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"

	"log"
)

// Hash func
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func CheckHashPassword(hashPassword string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(pass))
}

func GenerateRandomToken(lenght int) string {
	b := make([]byte, lenght)

	if _, e := rand.Read(b); e != nil {
		log.Printf("Failed to generate token: %v", e)
	}

	return base64.URLEncoding.EncodeToString(b)

}
