package main

import "fmt"

const frances = "frances"
const espanhol = "espanhol"
const prefixoBonjour = "Bonjour, "
const prefixoOla = "Olá, "
const prefixoHola = "Hola, "

func ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	if idioma == espanhol {
		return prefixoHola + nome
	}
	if idioma == frances {
		return prefixoBonjour + nome
	}
	return prefixoOla + nome
}

func main() {
	fmt.Println(ola("", ""))
}
