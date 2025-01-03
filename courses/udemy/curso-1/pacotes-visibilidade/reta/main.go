package main

import "fmt"

func main() {
	ponto1 := Ponto{2., 2.}
	ponto2 := Ponto{2., 4.}

	fmt.Println(catetos(ponto1, ponto2))
	fmt.Println(Distancia(ponto1, ponto2))
}
