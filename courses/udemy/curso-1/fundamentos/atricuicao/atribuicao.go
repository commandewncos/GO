package main

import (
	"fmt"
	"reflect"
)

var (
	s = 7
	q = 9
)

func main() {
	var b byte = 5
	fmt.Println("Tipo:", reflect.TypeOf(b))

	i := 4 //inferência de tipo
	i += 4 // i = i + 4
	i -= 4 // i = i + 4
	i /= 4 // i = i / 4
	i *= 4 // i = i * 4
	i %= 4 // i = i % 4

	x, y := 1, 2
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

	fmt.Println(s, q)

}
