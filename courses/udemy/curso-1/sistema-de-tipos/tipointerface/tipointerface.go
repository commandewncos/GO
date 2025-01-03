package main

import "fmt"

type Curso struct {
	nome string
}

func main() {
	var coisa interface {
	}

	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type Dinamico interface{}

	var outraCoisa Dinamico = "Epa!"

	fmt.Println(outraCoisa)
	outraCoisa = 35
	fmt.Println(outraCoisa)
	outraCoisa = true
	fmt.Println(outraCoisa)

	outraCoisa = Curso{nome: "GO"}
	fmt.Println(outraCoisa)
}
