package main

import (
	"fmt"
	"math"
)

const (
	a = 1
	b = 2
	c = 3
)

var (
	x = 11
	y = 12
	z = 13
)

func main() {
	m, n, o, p := 7, 8.9, true, "Vish!"
	const PI float64 = 3.1415

	var raio = 3.2

	area := PI * math.Pow(raio, 2)
	fmt.Println(PI, raio, area)
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
	fmt.Println(m, n, o, p)
}
