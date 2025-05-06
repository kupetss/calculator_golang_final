package db

import (
	"database/sql"
	"fmt"
	"log"
)

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`)

	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	log.Println("Database tables created successfully")
	return nil
}
