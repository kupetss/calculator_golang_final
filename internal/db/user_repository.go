package db

import (
	"calc_golang_final/internal/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (login, password) VALUES (?, ?)",
		user.Login,
		user.Password,
	)
	return err
}

func (r *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(
		"SELECT id, login, password FROM users WHERE login = ?",
		login,
	).Scan(&user.ID, &user.Login, &user.Password)

	return user, err
}
