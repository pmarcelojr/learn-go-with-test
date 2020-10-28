package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
		}
	}

	t.Run("para uma pessoa", func(t *testing.T) {
		resultado := ola("Marcelo", "")
		esperado := "Olá, Marcelo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("string vazia", func(t *testing.T) {
		resultado := ola("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em frances", func(t *testing.T) {
		resultado := ola("Ney", "frances")
		esperado := "Bonjour, Ney"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
