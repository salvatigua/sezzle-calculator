package handler

import (
	"encoding/json"
	"net/http"

	"sezzle-calculator/backend/internal/calculator"
)

type CalculatorHandler struct {
	service *calculator.Service
}

func NewCalculatorHandler(service *calculator.Service) *CalculatorHandler {
	return &CalculatorHandler{service: service}
}

// Structs para request y response JSON
type CalculationRequest struct {
	Operator string   `json:"operator"`
	A        *float64 `json:"a"`
	B        *float64 `json:"b,omitempty"`
}

type CalculationResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *CalculatorHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	var req CalculationRequest

	// Decodificar JSON del body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Payload JSON inválido o malformado"})
		return
	}

	// Validación: El operando 'a' es obligatorio
	if req.A == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "El operando 'a' es obligatorio"})
		return
	}

	// Ejecutar cálculo
	result, err := h.service.Calculate(req.Operator, *req.A, req.B)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	// Respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalculationResponse{Result: result})
}