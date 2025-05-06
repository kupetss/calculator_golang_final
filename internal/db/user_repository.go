package db

import (
	"database/sql"
)

type User struct {
	ID       int64
	Username string
	Password string // На практике пароль должен храниться в хэшированном виде
}

type UserRepository interface {
	CreateUser(username, password string) error
	GetUserByCredentials(username, password string) (*User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(username, password string) error {
	// Реализация сохранения пользователя в БД
	_, err := r.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}

func (r *userRepository) GetUserByCredentials(username, password string) (*User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, username FROM users WHERE username = ? AND password = ?", username, password).Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
