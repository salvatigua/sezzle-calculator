package calculator

import (
	"errors"
	"math"
)

// Tipos de operaciones soportadas
const (
	OpAdd        = "add"
	OpSubtract   = "subtract"
	OpMultiply   = "multiply"
	OpDivide     = "divide"
	OpPower      = "power"
	OpSqrt       = "sqrt"
	OpPercentage = "percentage"
)

// Errores de dominio
var (
	ErrDivisionByZero  = errors.New("division por cero no está permitida")
	ErrNegativeSqrt    = errors.New("no se puede calcular la raíz cuadrada de un número negativo")
	ErrInvalidOperator = errors.New("operador no válido o no soportado")
)

// CalculatorService contiene la interfaz y métodos de cálculo
type Service struct{}

func NewService() *Service {
	return &Service{}
}

// Calculate procesa la operación según los operandos ingresados
func (s *Service) Calculate(operator string, a float64, b *float64) (float64, error) {
	switch operator {
	case OpAdd:
		if b == nil {
			return 0, errors.New("el segundo operando (b) es requerido para la suma")
		}
		return a + *b, nil

	case OpSubtract:
		if b == nil {
			return 0, errors.New("el segundo operando (b) es requerido para la resta")
		}
		return a - *b, nil

	case OpMultiply:
		if b == nil {
			return 0, errors.New("el segundo operando (b) es requerido para la multiplicación")
		}
		return a * *b, nil

	case OpDivide:
		if b == nil {
			return 0, errors.New("el segundo operando (b) es requerido para la división")
		}
		if *b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / *b, nil

	case OpPower:
		if b == nil {
			return 0, errors.New("el segundo operando (b) es requerido para la exponenciación")
		}
		return math.Pow(a, *b), nil

	case OpSqrt:
		if a < 0 {
			return 0, ErrNegativeSqrt
		}
		return math.Sqrt(a), nil

	case OpPercentage:
		if b == nil {
			// Si solo hay un operando, retorna a / 100
			return a / 100.0, nil
		}
		// Si hay dos operandos, calcula el b% de a (a * (b/100))
		return a * (*b / 100.0), nil

	default:
		return 0, ErrInvalidOperator
	}
}