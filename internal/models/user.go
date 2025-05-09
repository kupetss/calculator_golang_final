package models

type User struct {
	ID           int    `db:"id"`
	Login        string `db:"login"` // Было Username
	PasswordHash string `db:"password_hash"`
}
