package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := repetir("a")
	esperado := "aaaaa"

	if esperado != repeticoes {
		t.Errorf("Esperado '%s', mas obteve '%s'", esperado, repeticoes)
	}
}
