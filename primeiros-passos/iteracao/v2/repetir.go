package iteracao

const qtdRepeticoes = 5

func repetir(caracter string) string {
	var repeticoes string
	for i := 0; i < qtdRepeticoes; i++ {
		repeticoes += caracter
	}
	return repeticoes
}
