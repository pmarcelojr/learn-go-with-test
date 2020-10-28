package inteiros

import "testing"

func TestAdicionador(t *testing.T) {
	soma := adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("Esperado '%d', Resultado '%d'", esperado, soma)
	}
}
