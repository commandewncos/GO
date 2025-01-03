package main

import "fmt"

func clousure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}

func main() {

	x := 33
	fmt.Println(x)

	imprimeX := clousure()
	imprimeX()

}
