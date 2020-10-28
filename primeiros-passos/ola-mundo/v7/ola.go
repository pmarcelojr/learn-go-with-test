package main

import "fmt"

const frances = "frances"
const espanhol = "espanhol"
const prefixoOla = "Olá, "
const prefixoHola = "Hola, "
const prefixoBonjour = "Bonjour, "

func ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	prefixo := prefixoOla
	switch idioma {
	case espanhol:
		prefixo = prefixoHola
	case frances:
		prefixo = prefixoBonjour
	}

	return prefixo + nome
}

func main() {
	fmt.Println(ola("Mundo", "frances"))
}
