package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixoOla = "Olá, "
const prefixoHola = "Hola, "
const prefixoBonjour = "Bonjour, "

func ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case espanhol:
		prefixo = prefixoHola
	case frances:
		prefixo = prefixoBonjour
	default:
		prefixo = prefixoOla
	}
	return
}

func main() {
	fmt.Println("", "")
}
