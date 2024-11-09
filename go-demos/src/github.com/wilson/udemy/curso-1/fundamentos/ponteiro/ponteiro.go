package main

import "fmt"

func main() {

	i := 5

	// Criando um ponteiro:
	var ponteiro *int = nil

	ponteiro = &i // Pegando o endereço da variavel
	*ponteiro++   //Desreferencindo (pegando o valor)
	i++
	fmt.Println("Endereço:", ponteiro, "Valor:", *ponteiro)
	fmt.Println("Valor:", i, "Endereço:", &i)

}
