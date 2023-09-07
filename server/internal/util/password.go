package util

import "golang.org/x/crypto/bcrypt"

const HASHCONST = 12

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), HASHCONST)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
