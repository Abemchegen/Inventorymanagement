package infrastructure

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

// NewPasswordService creates a new instance of PasswordService.
func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

// HashPassword hashes a plain text password.
func (p *PasswordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares a plain text password with a hashed password.
func (p *PasswordService) ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
