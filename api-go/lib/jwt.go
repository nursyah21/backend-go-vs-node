package lib

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const jwtSecretKey = "your-secret-key"

type CustomClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// Get the complete, signed token
func GenerateJwt(userID, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &CustomClaims{
		ID:       userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "user",
			Issuer:    "my-go-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}
