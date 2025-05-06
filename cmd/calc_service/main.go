package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kupetss/calculator_golang_final/internal/db"
	"github.com/kupetss/calculator_golang_final/internal/handlers"
	"github.com/kupetss/calculator_golang_final/internal/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "./calc.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	if err := db.Migrate(database); err != nil {
		log.Fatal("Migration failed:", err)
	}

	userRepo := db.NewUserRepository(database)

	router := http.NewServeMux()

	router.HandleFunc("POST /api/v1/register", handlers.RegisterHandler(userRepo))
	router.HandleFunc("POST /api/v1/login", handlers.LoginHandler(userRepo))

	protected := http.NewServeMux()
	protected.HandleFunc("POST /api/v1/calculate", handlers.CalculateHandler())
	router.Handle("/api/v1/", middleware.AuthMiddleware(protected))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
