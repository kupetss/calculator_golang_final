package main

import (
	"calc_golang_final/internal/db"
	"calc_golang_final/internal/handlers"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepo := &db.UserRepository{DB: database}
	authHandler := &handlers.AuthHandler{UserRepo: userRepo}

	http.HandleFunc("/api/v1/login", authHandler.Login)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
