package main

import "fmt"

func fatorial(numero uint) uint {

	if numero == 0 || numero == 1 {
		return 1
	} else {
		fatorialAnterior := fatorial(numero - 1)
		return numero * fatorialAnterior
	}

}

func main() {
	r := fatorial(6)
	fmt.Println(r)
}
