package calculator

import (
	"math"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		expected float64
		wantErr  bool
	}{
		{"Simple addition", "2 + 2", 4, false},
		{"Subtraction", "5 - 3", 2, false},
		{"Multiplication", "3 * 4", 12, false},
		{"Division", "10 / 2", 5, false},
		{"Complex expression", "2 + 3 * 4", 14, false},
		{"With parentheses", "(2 + 3) * 4", 20, false},
		{"Square root", "sqrt(9)", 3, false},
		{"Power function", "pow(2, 3)", 8, false},
		{"Sin function", "sin(0)", 0, false},
		{"Invalid expression", "2 + ", 0, true},
		{"Division by zero", "1 / 0", 0, true},
		{"Unknown function", "test(1)", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Evaluate(tt.expr)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error for expression '%s', got nil", tt.expr)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error for expression '%s': %v", tt.expr, err)
				return
			}
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("For expression '%s' expected %v, got %v", tt.expr, tt.expected, result)
			}
		})
	}
}
func TestEvaluateEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		wantErr bool
	}{
		{"Empty expression", "", true},
		{"Whitespace expression", "   ", true},
		{"Negative square root", "sqrt(-1)", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Evaluate(tt.expr)
			if tt.wantErr && err == nil {
				t.Errorf("Expected error for case '%s', got nil", tt.name)
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error for case '%s': %v", tt.name, err)
			}
		})
	}
}
