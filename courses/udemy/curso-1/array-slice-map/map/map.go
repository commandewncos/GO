package main

import (
	"fmt"
)

func main() {

	// var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[12345678987] = "João"
	aprovados[32187698200] = "Maria"
	aprovados[34521865298] = "Fulano"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s\t(CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678987])
	delete(aprovados, 12345678987)
	fmt.Println(aprovados)
}
