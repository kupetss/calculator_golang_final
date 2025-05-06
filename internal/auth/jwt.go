package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Секретный ключ для подписи (замените на свой в продакшене)
var jwtSecret = []byte("123")

// Claims - структура для хранения данных в токене
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// ValidateToken проверяет JWT токен и возвращает ID пользователя
func ValidateToken(tokenString string) (int, error) {
	// Парсим токен с нашими claims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	// Извлекаем claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, errors.New("invalid token")
}

// GenerateToken создает новый JWT токен для пользователя
func GenerateToken(userID int) (string, error) {
	// Устанавливаем срок действия (24 часа)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Создаем claims с данными
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "calc-service",
		},
	}

	// Создаем токен с алгоритмом HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен
	return token.SignedString(jwtSecret)
}
