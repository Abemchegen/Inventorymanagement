// infrastructure/jwt_service.go
package infrastructure

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTService handles JWT token generation and validation.
type JWTService struct {
	secretKey string
}

// NewJWTService creates a new instance of JWTService.
func NewJWTService(secretKey string) *JWTService {
	return &JWTService{secretKey: secretKey}
}

// GenerateToken generates a JWT token for the given username.
func (j *JWTService) GenerateToken(username string) (string, error) {
	// Create a new JWT token with claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expiry time: 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validates the JWT token and returns the username from it.
func (j *JWTService) ValidateToken(tokenString string) (string, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		log.Printf("Error validating token: %v", err)
		return "", err
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if username, ok := claims["username"].(string); ok {
			return username, nil
		}
	}

	return "", errors.New("invalid token claims")
}
