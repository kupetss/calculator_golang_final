package auth

import (
	"errors"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var (
	users = make(map[string]string) // login -> hashedPassword
	mu    sync.Mutex                // защита от гонок
)

// Register создает нового пользователя
func Register(login, password string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[login]; exists {
		return errors.New("user already exists")
	}

	hashed, err := HashPassword(password)
	if err != nil {
		return err
	}

	users[login] = hashed
	return nil
}

// Login проверяет логин/пароль
func Login(login, password string) (bool, error) {
	mu.Lock()
	defer mu.Unlock()

	hashed, exists := users[login]
	if !exists {
		return false, errors.New("user not found")
	}

	return CheckPassword(password, hashed), nil
}
