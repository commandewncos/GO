package main

import "fmt"

func main() {

	// Homogênea (mesmo tipo) e estática (fixo)
	// Valores são indexados

	var notas [3]float64

	notas[0], notas[1], notas[2] = 5.5, 6.7, 8.9
	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	lenght := len(notas)
	mean := total / float64(lenght)

	fmt.Printf("Média: %.2f", mean)

}
