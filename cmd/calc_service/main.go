package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/kupetss/calculator_golang_final/internal/grpc"
	"github.com/kupetss/calculator_golang_final/internal/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Инициализация БД
	db, err := sql.Open("sqlite3", "calculator.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Запуск gRPC сервера в горутине
	go func() {
		if err := grpc.StartGRPCServer(db); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	// HTTP роутинг
	http.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)
	http.HandleFunc("/api/v1/history", func(w http.ResponseWriter, r *http.Request) {
		handlers.HistoryHandler(w, r, db)
	})

	// Запуск HTTP сервера
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
