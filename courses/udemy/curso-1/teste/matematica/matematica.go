package matematica

//Média é reponsavel para calcular a média

func Media(numbers ...float64) float64 {
	sum := 0.

	for _, n := range numbers {
		sum += n

	}

	l := len(numbers)
	return sum / float64(l)
}
