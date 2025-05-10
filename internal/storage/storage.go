package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Task struct {
	ID         string  `json:"id"`
	UserID     int     `json:"user_id"`
	Expression string  `json:"expression"`
	Status     string  `json:"status"`
	Result     float64 `json:"result"`
	CreatedAt  string  `json:"created_at"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Добавьте новые константы в начало файла
const (
	JWTSecretKey = "your_very_secret_key_here" // Замените на случайную строку
	TokenExpire  = 24 * time.Hour
)

// Добавьте структуру для JWT claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tasks (
			id TEXT PRIMARY KEY,
			user_id INTEGER NOT NULL,
			expression TEXT NOT NULL,
			status TEXT NOT NULL,
			result FLOAT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`)
	return db, err
}

// User functions
func RegisterUser(db *sql.DB, username, password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		"INSERT INTO users (username, password_hash) VALUES (?, ?)",
		username, hash,
	)
	return err
}

func AuthenticateUser(db *sql.DB, username, password string) (int, error) {
	var id int
	var hash string

	err := db.QueryRow(
		"SELECT id, password_hash FROM users WHERE username = ?",
		username,
	).Scan(&id, &hash)

	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return id, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Task functions
func SaveTask(db *sql.DB, userID int, expr, status string) (string, error) {
	taskID := fmt.Sprintf("task_%d", time.Now().UnixNano())
	_, err := db.Exec(
		"INSERT INTO tasks (id, user_id, expression, status) VALUES (?, ?, ?, ?)",
		taskID, userID, expr, status,
	)
	return taskID, err
}

func CompleteTask(db *sql.DB, id string, result float64) error {
	_, err := db.Exec(
		"UPDATE tasks SET status = 'completed', result = ? WHERE id = ?",
		result, id,
	)
	return err
}

func GetUserTasks(db *sql.DB, userID int) ([]Task, error) {
	rows, err := db.Query(
		"SELECT id, expression, status, result FROM tasks WHERE user_id = ?", // Убрали created_at
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Expression,
			&task.Status,
			&task.Result,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
