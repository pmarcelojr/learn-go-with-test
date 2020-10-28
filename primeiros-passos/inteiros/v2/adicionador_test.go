package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	soma := adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("Esperado '%d', Resultado '%d'", esperado, soma)
	}
}

func ExampleAdiciona() {
	soma := adiciona(1, 5)
	fmt.Println(soma)
	// Output: 6
}
