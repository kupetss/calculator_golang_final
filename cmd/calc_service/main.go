package main

import (
	"calc_golang_final/internal/handlers"
	"calc_golang_final/middleware"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "calc.db")
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            login TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        )
    `)
	if err != nil {
		log.Fatal("Ошибка при создании таблицы users:", err)
	}

	log.Println("Таблица users готова")
	http.HandleFunc("/api/v1/register", handlers.RegisterHandler(db))
	http.HandleFunc("/api/v1/login", handlers.LoginHandler(db))
	http.Handle("/api/v1/calculate", middleware.AuthMiddleware(http.HandlerFunc(handlers.CalculateHandler)))
	log.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", nil)
}
