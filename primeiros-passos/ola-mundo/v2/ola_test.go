package main

import "testing"

func TestOla(t *testing.T) {
	resultado := ola()
	esperado := "Olá, Mundo"

	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}
