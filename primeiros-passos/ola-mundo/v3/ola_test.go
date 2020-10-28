package main

import "testing"

func TestOla(t *testing.T) {
	resultado := ola("Marcelo")
	esperado := "Olá, Marcelo"

	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}
