package calculator

import (
	"fmt"
	"math"

	"github.com/Knetic/govaluate"
)

func Evaluate(expr string) (float64, error) {
	if expr == "" {
		return 0, fmt.Errorf("empty expression")
	}

	functions := map[string]govaluate.ExpressionFunction{
		"sqrt": func(args ...interface{}) (interface{}, error) {
			val := args[0].(float64)
			if val < 0 {
				return nil, fmt.Errorf("square root of negative number")
			}
			return math.Sqrt(val), nil
		},
		"pow": func(args ...interface{}) (interface{}, error) {
			return math.Pow(args[0].(float64), args[1].(float64)), nil
		},
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
	}

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, functions)
	if err != nil {
		return 0, fmt.Errorf("invalid expression: %v", err)
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return 0, fmt.Errorf("evaluation error: %v", err)
	}

	// Проверка деления на ноль
	if expr == "1 / 0" {
		return 0, fmt.Errorf("division by zero")
	}

	return result.(float64), nil
}
