package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"sezzle-calculator/backend/internal/calculator"
	"sezzle-calculator/backend/internal/handler"
)

func TestCalculateHandler(t *testing.T) {
	service := calculator.NewService()
	h := handler.NewCalculatorHandler(service)

	tests := []struct {
		name           string
		payload        map[string]interface{}
		expectedStatus int
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "Operación de suma exitosa",
			payload:        map[string]interface{}{"operator": "add", "a": 10.0, "b": 5.0},
			expectedStatus: http.StatusOK,
			expectedResult: 15.0,
			expectError:    false,
		},
		{
			name:           "División por cero devuelve 422",
			payload:        map[string]interface{}{"operator": "divide", "a": 10.0, "b": 0.0},
			expectedStatus: http.StatusUnprocessableEntity,
			expectError:    true,
		},
		{
			name:           "Payload incompleto (falta operando 'a') devuelve 400",
			payload:        map[string]interface{}{"operator": "add", "b": 5.0},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "Raíz cuadrada exitosa (sin requerir b)",
			payload:        map[string]interface{}{"operator": "sqrt", "a": 25.0},
			expectedStatus: http.StatusOK,
			expectedResult: 5.0,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			h.Calculate(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("se esperaba status %d, pero se obtuvo %d", tt.expectedStatus, res.StatusCode)
			}

			if !tt.expectError {
				var resp handler.CalculationResponse
				if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
					t.Fatalf("error decodificando respuesta JSON: %v", err)
				}
				if resp.Result != tt.expectedResult {
					t.Errorf("se esperaba resultado %f, pero se obtuvo %f", tt.expectedResult, resp.Result)
				}
			}
		})
	}
}
