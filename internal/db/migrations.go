package db

import (
	"database/sql"
	"fmt"
	"log"
)

func Migrate(db *sql.DB) error {
	// Простой SQL-запрос без многострочного форматирования
	query := "CREATE TABLE IF NOT EXISTS users (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT," +
		"login TEXT NOT NULL UNIQUE," +
		"password TEXT NOT NULL," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP" +
		")"

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	log.Println("Database tables created successfully")
	return nil
}
