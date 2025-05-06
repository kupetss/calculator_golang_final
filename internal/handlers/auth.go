package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/kupetss/calculator_golang_final/internal/auth"
	"github.com/kupetss/calculator_golang_final/internal/types"
)

// RegisterHandler обрабатывает регистрацию пользователя
func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		hashedPassword, err := auth.HashPassword(req.Password)
		if err != nil {
			http.Error(w, "Failed to process password", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (login, password) VALUES (?, ?)", req.Login, hashedPassword)
		if err != nil {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"status": "user created"})
	}
}

// LoginHandler обрабатывает вход пользователя
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		var (
			userID     int
			storedPass string
		)
		err := db.QueryRow("SELECT id, password FROM users WHERE login = ?", req.Login).Scan(&userID, &storedPass)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if !auth.CheckPassword(req.Password, storedPass) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token, err := auth.GenerateToken(userID)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		response := types.LoginResponse{
			Token: token,
		}
		json.NewEncoder(w).Encode(response)
	}
}
