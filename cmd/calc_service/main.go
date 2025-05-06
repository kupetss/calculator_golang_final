package main

import (
	"calc_golang_final/internal/db"
	"calc_golang_final/internal/handlers"
	"calc_golang_final/internal/middleware"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Инициализация БД
	database, err := sql.Open("sqlite3", "calc.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Применение миграций
	if err := db.Migrate(database); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Инициализация репозиториев
	userRepo := db.NewUserRepository(database)

	// Настройка маршрутов
	http.HandleFunc("/api/v1/register", handlers.RegisterHandler(userRepo))
	http.HandleFunc("/api/v1/login", handlers.LoginHandler(userRepo))

	// Защищенные маршруты
	protected := http.NewServeMux()
	protected.HandleFunc("/api/v1/calculate", handlers.CalculateHandler())

	http.Handle("/api/v1/", middleware.AuthMiddleware(protected))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
