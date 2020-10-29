package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticao := repetir("a")
	esperado := "aaaaa"

	if esperado != repeticao {
		t.Errorf("Esperado '%s', mas obteve '%s'", esperado, repeticao)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repetir("a")
	}
}

func ExampleRepetir() {
	fmt.Println("a")
	// Output: a
}
