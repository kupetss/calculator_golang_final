package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kupetss/calculator_golang_final/internal/auth"
	"github.com/kupetss/calculator_golang_final/internal/db"
)

type AuthRequest struct {
	Login    string `json:"login"`
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

		// Хеширование пароля перед сохранением
		hashedPassword, err := auth.HashPassword(req.Password)
		if err != nil {
			response := AuthResponse{Error: "failed to hash password"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}

		err = repo.CreateUser(req.Login, hashedPassword)
		if err != nil {
			response := AuthResponse{Error: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// После регистрации сразу генерируем токен
		user, err := repo.GetUserByCredentials(req.Login, hashedPassword)
		if err != nil {
			response := AuthResponse{Error: "failed to login after registration"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}

		token, err := auth.GenerateToken(int(user.ID))
		if err != nil {
			response := AuthResponse{Error: "failed to generate token"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := AuthResponse{Token: token}
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

		// Получаем пользователя по логину
		user, err := repo.GetUserByLogin(req.Login)
		if err != nil {
			response := AuthResponse{Error: "invalid credentials"}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Проверяем пароль
		if !auth.CheckPassword(req.Password, user.Password) {
			response := AuthResponse{Error: "invalid credentials"}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Генерируем JWT токен
		token, err := auth.GenerateToken(int(user.ID))
		if err != nil {
			response := AuthResponse{Error: "failed to generate token"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := AuthResponse{Token: token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
