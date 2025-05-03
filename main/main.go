package main

import (
	"calculator_golangV3/config/calculator"
	"calculator_golangV3/config/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	calculator.Init()
	handlers.Init()
	err := os.MkdirAll("database", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	file, err := os.OpenFile("database/results.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	file.Close()

	log.Println("Server is starting...")
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/calculate", handlers.HandleCompute).Methods("POST")
	router.HandleFunc("/api/v1/expressions/{id}", handlers.HandleGet).Methods("GET")
	router.HandleFunc("/api/v1/expressions", handlers.HandleList).Methods("GET")
	router.HandleFunc("/internal/task", handlers.HandleOrchestrate).Methods("POST")
	router.HandleFunc("/api/v1/history", handlers.HandleHistory).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	}).Methods("GET")

	http.ListenAndServe(":8080", enableCORS(loggingMiddleware(router)))
}

func main() {
	StartServer()
}
