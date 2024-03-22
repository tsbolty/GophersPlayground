package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Generates a new access token and refresh token for a user.
func GenerateNewTokens(userID uint) (string, string, error) {
	// Generate a new access token
	newAccessToken, err := GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}

	// Generate a new refresh token
	newRefreshToken, err := GenerateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

// Generates a JWT access token for a user.
func GenerateAccessToken(userID uint) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"typ": "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Generates a refresh token for a user.
func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // 1 week
		"typ": "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Validates a given refresh token and returns the user ID.
func ValidateRefreshToken(token string) (uint, error) {
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !parsedToken.Valid {
		return 0, err
	}

	if typ, ok := claims["typ"].(string); !ok || typ != "refresh" {
		return 0, jwt.NewValidationError("Not a refresh token", jwt.ValidationErrorClaimsInvalid)
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return 0, jwt.NewValidationError("User ID invalid", jwt.ValidationErrorClaimsInvalid)
	}

	return uint(userId), nil
}
