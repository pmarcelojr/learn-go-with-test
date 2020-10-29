package iteracao

const qtdRepeticao = 5

func repetir(caracter string) string {
	var repeticoes string
	for i := 0; i < qtdRepeticao; i++ {
		repeticoes += caracter
	}
	return repeticoes
}
