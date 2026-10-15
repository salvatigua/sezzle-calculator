import type { CalculationRequest, CalculationResponse, ErrorResponse } from '../types/calculator';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8081';

export async function calculate(request: CalculationRequest): Promise<number> {
  const response = await fetch(`${API_BASE_URL}/api/v1/calculate`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(request),
  });

  const data = await response.json();

  if (!response.ok) {
    const errorData = data as ErrorResponse;
    throw new Error(errorData.error || 'Ocurrió un error inesperado');
  }

  return (data as CalculationResponse).result;
}