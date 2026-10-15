import React, { useState } from 'react';
import { calculate } from '../api/calculatorApi';
import type { Operator } from '../types/calculator';

export const Calculator: React.FC = () => {
  const [display, setDisplay] = useState<string>('0');
  const [operandA, setOperandA] = useState<number | null>(null);
  const [operator, setOperator] = useState<Operator | null>(null);
  const [waitingForOperandB, setWaitingForOperandB] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  const handleDigit = (digit: string) => {
    setError(null);
    if (waitingForOperandB) {
      setDisplay(digit);
      setWaitingForOperandB(false);
    } else {
      setDisplay(display === '0' ? digit : display + digit);
    }
  };

  const handleDecimal = () => {
    setError(null);
    if (waitingForOperandB) {
      setDisplay('0.');
      setWaitingForOperandB(false);
      return;
    }
    if (!display.includes('.')) {
      setDisplay(display + '.');
    }
  };

  const handleClear = () => {
    setDisplay('0');
    setOperandA(null);
    setOperator(null);
    setWaitingForOperandB(false);
    setError(null);
  };

  const handleOperationSelect = (op: Operator) => {
    setError(null);
    const currentValue = parseFloat(display);

    if (op === 'sqrt' || op === 'percentage') {
      executeSingleOperandOp(op, currentValue);
      return;
    }

    setOperandA(currentValue);
    setOperator(op);
    setWaitingForOperandB(true);
  };

  const executeSingleOperandOp = async (op: Operator, val: number) => {
    setLoading(true);
    setError(null);
    try {
      const res = await calculate({ operator: op, a: val });
      setDisplay(String(res));
      setOperandA(res);
      setWaitingForOperandB(true);
    } catch (err: any) {
      setError(err.message || 'Error de cálculo');
    } finally {
      setLoading(false);
    }
  };

  const handleEquals = async () => {
    if (operator === null || operandA === null) return;

    const operandB = parseFloat(display);
    setLoading(true);
    setError(null);

    try {
      const result = await calculate({
        operator,
        a: operandA,
        b: operandB,
      });
      setDisplay(String(result));
      setOperandA(result);
      setOperator(null);
      setWaitingForOperandB(true);
    } catch (err: any) {
      setError(err.message || 'Error de cálculo');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded-2xl shadow-xl border border-gray-100">
      <h1 className="text-2xl font-bold !text-gray-900 text-center mb-4">
        Sezzle Fullstack Calculator
        </h1>

      {/* Display */}
      <div className="bg-gray-900 text-white p-4 rounded-xl text-right mb-4 min-h-[80px] flex flex-col justify-center">
        <div className="text-xs text-gray-400 h-4">
          {operandA !== null && operator && `${operandA} ${operator}`}
        </div>
        <div className="text-3xl font-mono overflow-x-auto">
          {loading ? 'Calculando...' : display}
        </div>
      </div>

      {/* Alerta de Error */}
    {error && (
    <div className="mb-4 p-3 bg-red-50 border-l-4 border-red-500 text-red-700 rounded-r-md text-sm font-medium">
        <p>{error}</p>
    </div>
    )}

      {/* Teclado */}
      <div className="grid grid-cols-4 gap-2">
        <button
          onClick={handleClear}
          className="col-span-2 p-3 bg-red-500 hover:bg-red-600 text-white font-bold rounded-lg transition"
        >
          AC
        </button>
        <button
          onClick={() => handleOperationSelect('sqrt')}
          className="p-3 bg-indigo-100 hover:bg-indigo-200 text-indigo-700 font-semibold rounded-lg transition"
        >
          √
        </button>
        <button
          onClick={() => handleOperationSelect('divide')}
          className="p-3 bg-indigo-600 hover:bg-indigo-700 text-white font-semibold rounded-lg transition"
        >
          ÷
        </button>

        <button onClick={() => handleDigit('7')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">7</button>
        <button onClick={() => handleDigit('8')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">8</button>
        <button onClick={() => handleDigit('9')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">9</button>
        <button
          onClick={() => handleOperationSelect('multiply')}
          className="p-3 bg-indigo-600 hover:bg-indigo-700 text-white font-semibold rounded-lg transition"
        >
          ×
        </button>

        <button onClick={() => handleDigit('4')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">4</button>
        <button onClick={() => handleDigit('5')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">5</button>
        <button onClick={() => handleDigit('6')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">6</button>
        <button
          onClick={() => handleOperationSelect('subtract')}
          className="p-3 bg-indigo-600 hover:bg-indigo-700 text-white font-semibold rounded-lg transition"
        >
          -
        </button>

        <button onClick={() => handleDigit('1')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">1</button>
        <button onClick={() => handleDigit('2')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">2</button>
        <button onClick={() => handleDigit('3')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">3</button>
        <button
          onClick={() => handleOperationSelect('add')}
          className="p-3 bg-indigo-600 hover:bg-indigo-700 text-white font-semibold rounded-lg transition"
        >
          +
        </button>

        <button onClick={() => handleDigit('0')} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">0</button>
        <button onClick={handleDecimal} className="p-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-semibold rounded-lg">.</button>
        <button
          onClick={() => handleOperationSelect('percentage')}
          className="p-3 bg-indigo-100 hover:bg-indigo-200 text-indigo-700 font-semibold rounded-lg transition"
        >
          %
        </button>
        <button
          onClick={() => handleOperationSelect('power')}
          className="p-3 bg-indigo-100 hover:bg-indigo-200 text-indigo-700 font-semibold rounded-lg transition"
        >
          xʸ
        </button>

        <button
          onClick={handleEquals}
          className="col-span-4 mt-2 p-3 bg-green-600 hover:bg-green-700 text-white font-bold rounded-lg shadow transition"
        >
          =
        </button>
      </div>
    </div>
  );
};