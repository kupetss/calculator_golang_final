package handlers

import (
	"calc_golang_final/internal/auth"
	"calc_golang_final/internal/types"
	"database/sql"
	"encoding/json"
	"net/http"
)

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неверный формат данных", http.StatusBadRequest)
			return
		}

		hashedPassword, err := auth.HashPassword(req.Password)
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (login, password) VALUES (?, ?)",
			req.Login, hashedPassword)
		if err != nil {
			http.Error(w, "Логин уже занят", http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Пользователь зарегистрирован"))
	}
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Парсим запрос
		var req types.RegisterRequest // Используем ту же структуру
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неверный формат данных", http.StatusBadRequest)
			return
		}

		// 2. Ищем пользователя в БД
		var storedPassword string
		err := db.QueryRow("SELECT password FROM users WHERE login = ?", req.Login).Scan(&storedPassword)
		if err != nil {
			http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
			return
		}

		// 3. Проверяем пароль
		if !auth.CheckPassword(req.Password, storedPassword) {
			http.Error(w, "Неверный пароль", http.StatusUnauthorized)
			return
		}

		// 4. Генерируем JWT-токен
		token, err := auth.GenerateToken(req.Login)
		if err != nil {
			http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
			return
		}

		// 5. Отправляем токен
		json.NewEncoder(w).Encode(types.LoginResponse{Token: token})
	}
}
