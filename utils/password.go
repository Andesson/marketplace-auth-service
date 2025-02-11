package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func ComparePasswordHash(password, hash, salt string) bool {
	saltedPassword := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedPassword))
	return err == nil
}

func GenerateRandomSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(salt), nil
}
