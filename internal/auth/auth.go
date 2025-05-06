package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("123")

func GenerateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(secretKey)
}
