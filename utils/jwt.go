package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateToken generates a JWT token for the specified user ID and role
func GenerateToken(userID uint, role string) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set the expiration time to 24 hours

	// Generate the signed token string
	tokenString, err := token.SignedString([]byte(GetJWTSecret()))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func GenerateRandomKey(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer), nil
}

func GetJWTSecret() string {
	return "rPzH6Pi8s5nZkLWokkCc1fJr4e4zP-vO0V4X08XMGTk=" // Replace with your actual secret key
}