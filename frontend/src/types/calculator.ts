export type Operator = 'add' | 'subtract' | 'multiply' | 'divide' | 'power' | 'sqrt' | 'percentage';

export interface CalculationRequest {
  operator: Operator;
  a: number;
  b?: number;
}

export interface CalculationResponse {
  result: number;
}

export interface ErrorResponse {
  error: string;
}