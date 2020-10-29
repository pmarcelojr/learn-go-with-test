package main

func soma(numeros [5]int) int {
	soma := 0
	for i := 0; i < 5; i++ {
		soma += numeros[i]
	}
	return soma
}
