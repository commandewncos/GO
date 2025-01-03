package main

import (
	"fmt"
)

type constraintCustom interface {
	int | float64 | string
}

func reverse[T constraintCustom](slice []T) []T {
	news := make([]T, len(slice))

	lenght := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		news[lenght] = slice[i]
		lenght--
	}

	return news
}

func main() {
	q := []float64{0.95, 5.5, 12, 8}
	r := []string{"a", "e", "i", "o", "u"}
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	reverseQ := reverse(q)
	reverseS := reverse(s)
	reverseR := reverse(r)

	fmt.Println(reverseQ)
	fmt.Println(reverseS)
	fmt.Println(reverseR)

}
