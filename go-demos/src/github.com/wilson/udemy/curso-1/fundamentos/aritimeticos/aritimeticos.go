package main

import (
	"fmt"
	"math"
)

var (
	a = 3
	b = 2
)

func main() {

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Modulo: Resto da divisão =>", a%b)

	/*
		Operation bitwise, bit to bit

		exemple:
	*/
	fmt.Println("E =>", a&b)   //11 & 10 = 10
	fmt.Println("OU =>", a|b)  //11 | 10 = 11
	fmt.Println("Xor =>", a^b) //11 ^10 = 01

	// Operations with package MATH:

	c := 5.5
	d := 6.7

	fmt.Println("Maior:", math.Max(float64(a), float64(b))) //needs to be float
	fmt.Println("Menor:", math.Min(c, d))
	fmt.Println("Exponenciação:", math.Pow(c, d))
}
