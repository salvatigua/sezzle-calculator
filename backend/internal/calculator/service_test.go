package calculator_test

import (
	"math"
	"testing"

	"sezzle-calculator/backend/internal/calculator"
)

func helperFloatPtr(val float64) *float64 {
	return &val
}

func TestCalculate(t *testing.T) {
	service := calculator.NewService()

	tests := []struct {
		name        string
		operator    string
		a           float64
		b           *float64
		expected    float64
		expectErr   bool
		expectedErr error
	}{
		// Suma
		{"Suma exitosa", calculator.OpAdd, 10, helperFloatPtr(5), 15, false, nil},
		{"Suma sin b", calculator.OpAdd, 10, nil, 0, true, nil},

		// Resta
		{"Resta exitosa", calculator.OpSubtract, 10, helperFloatPtr(4), 6, false, nil},

		// Multiplicación
		{"Multiplicación exitosa", calculator.OpMultiply, 3, helperFloatPtr(7), 21, false, nil},

		// División
		{"División exitosa", calculator.OpDivide, 20, helperFloatPtr(4), 5, false, nil},
		{"División por cero", calculator.OpDivide, 10, helperFloatPtr(0), 0, true, calculator.ErrDivisionByZero},

		// Exponenciación
		{"Exponenciación exitosa", calculator.OpPower, 2, helperFloatPtr(3), 8, false, nil},

		// Raíz cuadrada
		{"Raíz cuadrada exitosa", calculator.OpSqrt, 16, nil, 4, false, nil},
		{"Raíz de número negativo", calculator.OpSqrt, -9, nil, 0, true, calculator.ErrNegativeSqrt},

		// Porcentaje
		{"Porcentaje un solo operando (10%)", calculator.OpPercentage, 10, nil, 0.1, false, nil},
		{"Porcentaje de dos operandos (20% de 150)", calculator.OpPercentage, 150, helperFloatPtr(20), 30, false, nil},

		// Operador inválido
		{"Operador no soportado", "invalid_op", 5, helperFloatPtr(5), 0, true, calculator.ErrInvalidOperator},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Calculate(tt.operator, tt.a, tt.b)

			if tt.expectErr {
				if err == nil {
					t.Errorf("se esperaba un error pero se obtuvo nil")
				}
				if tt.expectedErr != nil && err != tt.expectedErr {
					t.Errorf("se esperaba el error '%v', pero se obtuvo '%v'", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("no se esperaba error, pero se obtuvo: %v", err)
				}
				if math.Abs(result-tt.expected) > 1e-9 {
					t.Errorf("se esperaba resultado %f, pero se obtuvo %f", tt.expected, result)
				}
			}
		})
	}
}