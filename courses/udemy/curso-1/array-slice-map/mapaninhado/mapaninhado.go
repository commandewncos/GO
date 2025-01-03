package main

import (
	"fmt"
)

func main() {

	funcsPorLetra := map[string]map[string]float32{
		"A": {
			"Ana":       5.5,
			"Ana Maria": 5.,
		},

		"W": {
			"Wilson": 7.4,
		},
	}

	for letra, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Println(letra, nome, salario)
		}
	}
}
