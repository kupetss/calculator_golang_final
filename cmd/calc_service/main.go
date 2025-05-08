package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kupetss/calculator_golang_final/internal/auth"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./calculator.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	auth.InitAuth(db)

	router := http.NewServeMux()
	router.HandleFunc("/api/v1/register", auth.RegisterHandler)
	router.HandleFunc("/api/v1/login", auth.LoginHandler)

	protected := http.NewServeMux()
	protected.Handle("/api/v1/calculate", auth.AuthMiddleware(http.HandlerFunc(calculateHandler)))

	mainHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v1/register", "/api/v1/login":
			router.ServeHTTP(w, r)
		default:
			protected.ServeHTTP(w, r)
		}
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mainHandler))
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ваша логика калькулятора
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"result": "Calculation result",
	})
}
