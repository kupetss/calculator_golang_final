package db

import (
	"calc_golang_final/internal/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		"SELECT id, login, password_hash FROM users WHERE login = ?",
		login,
	).Scan(&user.ID, &user.Login, &user.PasswordHash)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
