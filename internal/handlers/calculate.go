package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// CalculateRequest представляет структуру входящего запроса
type CalculateRequest struct {
	Expression string `json:"expression"`
}

// CalculateResponse представляет структуру ответа
type CalculateResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

// CalculateHandler создает обработчик для вычисления выражений
func CalculateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Проверяем метод запроса
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		// 2. Парсим тело запроса
		var req CalculateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		// 3. Проверяем, что выражение не пустое
		if req.Expression == "" {
			response := CalculateResponse{
				Error: "Expression cannot be empty",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// 4. Вычисляем результат
		result, err := evaluateExpression(req.Expression)
		if err != nil {
			response := CalculateResponse{
				Error: err.Error(),
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// 5. Возвращаем успешный ответ
		response := CalculateResponse{
			Result: result,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// evaluateExpression вычисляет значение математического выражения
func evaluateExpression(expr string) (float64, error) {
	// Удаляем лишние пробелы
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return 0, fmt.Errorf("empty expression")
	}

	// Разбиваем выражение на части
	parts := strings.Fields(expr)
	if len(parts) != 3 {
		return 0, fmt.Errorf("expression must be in format 'a operator b' (e.g. '2 + 3')")
	}

	// Парсим операнды
	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid first operand: %v", err)
	}

	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid second operand: %v", err)
	}

	// Выполняем операцию
	switch parts[1] {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator '%s'", parts[1])
	}
}
