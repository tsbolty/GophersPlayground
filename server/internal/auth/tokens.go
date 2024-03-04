package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// generate JWT token
func GenerateToken(userID uint) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// sign the token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
