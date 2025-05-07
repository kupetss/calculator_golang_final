package db

import "database/sql"

type User struct {
	ID       int64
	Login    string // Было: Username → Login
	Password string
}

type UserRepository interface {
	CreateUser(login, password string) error
	GetUserByCredentials(login, password string) (*User, error)
	GetUserByLogin(login string) (*User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(login, password string) error {
	_, err := r.db.Exec(
		"INSERT INTO users (login, password) VALUES (?, ?)", // Исправлено: username → login
		login, password,
	)
	return err
}

func (r *userRepository) GetUserByCredentials(login, password string) (*User, error) {
	var user User
	err := r.db.QueryRow(
		"SELECT id, login FROM users WHERE login = ? AND password = ?", // Исправлено: username → login
		login, password,
	).Scan(&user.ID, &user.Login)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByLogin(login string) (*User, error) {
	var user User
	err := r.db.QueryRow(
		"SELECT id, login, password FROM users WHERE login = ?",
		login,
	).Scan(&user.ID, &user.Login, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
