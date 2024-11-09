package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Mesma")
	fmt.Print("linha")
	fmt.Println("Mesma linha...")

	x := 3.14

	// Essa concatenação não é aceitavel porque x não é uma string
	// fmt.Println("O valor de x é:" + x)

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é:" + xs)

	fmt.Printf("O valor de pi é %.3f", math.Pi)

	fmt.Printf("\n\t%t %d %s %f", false, 35, "Wilson", 3.14)
	fmt.Printf("\n\t%v %v %v %v", false, 35, "Wilson", 3.14)
}
