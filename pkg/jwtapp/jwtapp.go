package jwtapp

import (
	"fmt"
	"time"

	"firebase.google.com/go/auth"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	User *auth.UserRecord `json:"user"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(user *auth.UserRecord) (string, error) {
	// Set session expiration to 5 days.
	expiresIn := time.Now().Add(time.Hour * 24 * 5)
	// Set token claims
	claims := MyClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresIn), // Token expires in 24 hours
		},
	}

	// Create a new token with the claims and the HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("rahasia") // Replace with a strong secret key

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyJWTToken verifies the JWT token and returns the claims on success
func VerifyJWTToken(tokenString string, secretKey string) (*MyClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if secretKey == "" {
			secretKey = "rahasia"
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract claims from token")
	}

	return claims, nil
}
