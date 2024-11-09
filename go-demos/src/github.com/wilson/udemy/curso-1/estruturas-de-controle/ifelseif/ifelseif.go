package main

import (
	"fmt"
)

func main() {
	// var n float64
	n := 8.75
	r := notaParaConceito(n)
	fmt.Println(r)
}

// Nota para conceito
func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}
