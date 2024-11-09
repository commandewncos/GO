package main

import "fmt"

// Quantidade de parametros indefinidos
func media(numeros ...float64) float64 {

	total := 0.
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))

}

func main() {
	numeros := []float64{1.5, 3.5, 5.5, 7.5, 9.5}
	fmt.Println(media(numeros...))
}
