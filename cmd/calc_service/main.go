package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kupetss/calculator_golang_final/internal/auth"
	"github.com/kupetss/calculator_golang_final/internal/grpc"
	"github.com/kupetss/calculator_golang_final/internal/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 1. Инициализация базы данных
	db, err := sql.Open("sqlite3", "./calculator.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Создаем таблицы если их нет
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS calculations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			expression TEXT NOT NULL,
			result TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`); err != nil {
		log.Fatal("Failed to create tables:", err)
	}

	// 2. Инициализация компонентов
	auth.InitAuth(db)
	grpc.InitCalculatorService(db)

	// 3. Запуск gRPC сервера в отдельной горутине
	go func() {
		log.Println("Starting gRPC server on :50051...")
		if err := grpc.StartGRPCServer(); err != nil {
			log.Fatal("gRPC server failed:", err)
		}
	}()

	// 4. Настройка HTTP маршрутов
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/register", auth.RegisterHandler)
	router.HandleFunc("/api/v1/login", auth.LoginHandler)

	protectedRouter := http.NewServeMux()
	protectedRouter.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)
	protectedRouter.HandleFunc("/api/v1/history", handlers.HistoryHandler)

	// 5. Настройка middleware
	mainHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v1/register", "/api/v1/login":
			router.ServeHTTP(w, r)
		default:
			auth.AuthMiddleware(protectedRouter).ServeHTTP(w, r)
		}
	})

	// 6. Запуск HTTP сервера
	server := &http.Server{
		Addr:    ":8080",
		Handler: mainHandler,
	}

	log.Println("HTTP server starting on :8080...")
	log.Fatal(server.ListenAndServe())
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
