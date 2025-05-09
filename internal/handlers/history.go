package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func HistoryHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	// Здесь должна быть ваша логика получения истории из БД
	history := []string{"2+2=4", "3*3=9"} // Заглушка

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
