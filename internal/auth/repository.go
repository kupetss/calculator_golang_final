package auth

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Init() error {
	_, err := r.db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            login TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL
        )
    `)
	return err
}

func (r *UserRepository) CreateUser(login, password string) error {
	_, err := r.db.Exec("INSERT INTO users (login, password) VALUES (?, ?)", login, password)
	return err
}

func (r *UserRepository) GetUser(login string) (*User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, login, password FROM users WHERE login = ?", login).Scan(
		&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
