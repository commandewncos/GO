package main

import "fmt"

func main() {

	/*
		Contagem é feita pelo compilador
		Se tirarmos as redisensias o será um slice
	*/

	numeros := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, numero := range numeros {
		fmt.Println("Indice:", i, "Valor:", numero)
	}

	// Ignorando o primeiro retorno
	for _, numero := range numeros {
		fmt.Println("Apenas o valor:", numero)
	}
}
