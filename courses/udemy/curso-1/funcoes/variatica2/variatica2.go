package main

import "fmt"

func imprimirAprovados(aprovados ...string) {

	fmt.Println(aprovados)

	for i, aprovado := range aprovados {

		fmt.Printf("%d)\t%s\n", i+1, aprovado)
	}

}

func main() {
	aprovados := []string{"Wilson", "Jose", "Maria", "Jessica"}
	imprimirAprovados(aprovados...)
}
