package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado inválido: Resultado: %d. Esperado: %d", total, 30)
	}
}

func TestSub(t *testing.T) {
	total := sub(10, 5)

	if total != 5 {
		t.Errorf("Resultado inválido: Resultado: %d. Esperado:  %d", total, 5)
	}
}

func TestTimes(t *testing.T) {
	total := times(2, 5)

	if total != 10 {
		t.Errorf("Resultado inválido: Resultado: %d. Esperado:  %d", total, 5)
	}
}
