package storage

import (
	"database/sql"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// Создаем временную БД для тестов
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:") // Используем БД в памяти
	if err != nil {
		t.Fatal(err)
	}

	// Создаем таблицы (упрощенная версия)
	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL
		);
		CREATE TABLE tasks (
			id TEXT PRIMARY KEY,
			user_id INTEGER NOT NULL,
			expression TEXT NOT NULL,
			status TEXT NOT NULL,
			result FLOAT
		);
	`)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestRegisterAndLogin(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Тест регистрации
	err := RegisterUser(db, "testuser", "password123")
	if err != nil {
		t.Errorf("RegisterUser failed: %v", err)
	}

	// Тест аутентификации
	userID, err := AuthenticateUser(db, "testuser", "password123")
	if err != nil {
		t.Errorf("AuthenticateUser failed: %v", err)
	}
	if userID == 0 {
		t.Error("Expected userID > 0, got 0")
	}

	// Тест неверного пароля
	_, err = AuthenticateUser(db, "testuser", "wrongpass")
	if err == nil {
		t.Error("Expected error for wrong password, got nil")
	}
}

func TestTasksCRUD(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Сначала создаем тестового пользователя
	err := RegisterUser(db, "testuser", "password123")
	if err != nil {
		t.Fatal(err)
	}

	// Тест создания задачи
	taskID, err := SaveTask(db, 1, "2+2", "pending")
	if err != nil {
		t.Errorf("SaveTask failed: %v", err)
	}
	if taskID == "" {
		t.Error("Expected taskID, got empty string")
	}

	// Тест завершения задачи
	err = CompleteTask(db, taskID, 4.0)
	if err != nil {
		t.Errorf("CompleteTask failed: %v", err)
	}

	// Тест получения задач
	tasks, err := GetUserTasks(db, 1)
	if err != nil {
		t.Errorf("GetUserTasks failed: %v", err)
	}
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Result != 4.0 {
		t.Errorf("Expected result 4.0, got %v", tasks[0].Result)
	}
}

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("testpass")
	if err != nil {
		t.Errorf("HashPassword failed: %v", err)
	}
	if hash == "" {
		t.Error("Expected hash, got empty string")
	}

	// Проверяем, что пароль проверяется корректно
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte("testpass"))
	if err != nil {
		t.Errorf("Password verification failed: %v", err)
	}
}
