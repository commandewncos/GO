package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!")

	if numero > 500 {
		fmt.Println("Valor alto...")
		return 500
	} else {
		fmt.Println("Valor baixo...")
		return numero
	}

}

func main() {
	fmt.Println(obterValorAprovado(600))
	fmt.Println(obterValorAprovado(400))
}
