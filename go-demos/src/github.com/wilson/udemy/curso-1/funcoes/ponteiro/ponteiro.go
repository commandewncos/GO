package main

import "fmt"

// Ponteiro é representao por um asterisco (*)

func numero(n int) { n++ }

func numeroPonteiro(n *int) {
	/*
		- (*) Asterisco é usado para referenciar, ou seja,
		ter acesso ao valor no qual o ponteiro aponta
	*/

	*n++

}

func main() {
	n := 5
	numero(n) //Referencia por valor
	fmt.Println(n)

	// (E - Comercial) é usado para obter o endereço da variável

	numeroPonteiro(&n) //Referenciando a variavel
	fmt.Println(n)

}
