package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kupetss/calculator_golang_final/internal/db"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

func RegisterHandler(repo db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Реальная реализация должна:
		// 1. Проверять, существует ли пользователь
		// 2. Хэшировать пароль
		// 3. Сохранять в БД
		err := repo.CreateUser(req.Username, req.Password)
		if err != nil {
			response := AuthResponse{Error: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := AuthResponse{Token: "generated-jwt-token"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func LoginHandler(repo db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Реальная реализация должна:
		// 1. Проверять учетные данные
		// 2. Генерировать JWT токен
		_, err := repo.GetUserByCredentials(req.Username, req.Password)
		if err != nil {
			response := AuthResponse{Error: "invalid credentials"}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := AuthResponse{Token: "generated-jwt-token"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
