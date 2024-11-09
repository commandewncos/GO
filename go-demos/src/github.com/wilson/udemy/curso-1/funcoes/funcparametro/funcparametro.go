package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, param1, param2 int) int {
	return funcao(param1, param2)

}

func main() {
	resultado := exec(multiplicacao, 3, 9)
	fmt.Println(resultado)
}
