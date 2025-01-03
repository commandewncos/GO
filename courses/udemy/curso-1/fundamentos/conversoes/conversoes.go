package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 5

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	// cuidado...Isso printa o valor do unicode
	fmt.Println("Teste" + string(9750))

	// Converssão int para string
	fmt.Println("Teste" + strconv.Itoa(9750))

	// Converssão de string para int
	num, _ := strconv.Atoi("123")

	fmt.Println(num)

	b, _ := strconv.ParseBool("0")

	if b {
		fmt.Println("O valor de b:", b)
	} else {
		fmt.Println("O valor de b:", b)
	}

}
