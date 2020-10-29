package iteracao

func repetir(caracter string) string {
	var repeticoes string
	for i := 0; i < 5; i++ {
		repeticoes = repeticoes + caracter
	}
	return repeticoes
}
