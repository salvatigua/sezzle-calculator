import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import { Calculator } from './Calculator';

// Mock de la API REST
vi.mock('../api/calculatorApi', () => ({
  calculate: vi.fn().mockResolvedValue(10),
}));

describe('Calculator Component', () => {
  it('renderiza el título y la pantalla inicial', () => {
    render(<Calculator />);
    const heading = screen.getByRole('heading', { level: 1 });
    expect(heading.textContent).toContain('Sezzle Fullstack Calculator');
  });

  it('actualiza la pantalla al presionar los dígitos numéricos', () => {
    render(<Calculator />);
    
    // Hacer click en los botones buscando por su texto directo
    const button7 = screen.getByText('7');
    const button5 = screen.getByText('5');

    fireEvent.click(button7);
    fireEvent.click(button5);

    // Buscar el valor 75 renderizado en la pantalla
    const displayElement = screen.getByText('75');
    expect(displayElement).toBeTruthy();
  });

  it('limpia la pantalla al presionar el botón AC', () => {
    render(<Calculator />);
    
    const button9 = screen.getByText('9');
    const buttonAC = screen.getByText('AC');

    fireEvent.click(button9);
    fireEvent.click(buttonAC);

    // Verificar que el display vuelve a mostrar 0
    const displayElement = screen.getByText('0', { selector: '.font-mono' });
    expect(displayElement).toBeTruthy();
  });
});