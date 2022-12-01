package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckedPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	// here when compare, by hashed password we already know the salt and cost, 
	// so when hashing password, the result will be same as hashed password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}