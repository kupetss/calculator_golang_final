package middleware

import (
	"calc_golang_final/internal/auth"
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Получаем токен из заголовка
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Токен отсутствует", http.StatusUnauthorized)
			return
		}

		// 2. Проверяем токен
		login, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}

		// 3. Добавляем логин в контекст запроса
		ctx := context.WithValue(r.Context(), "login", login)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
