package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	conditional := t.Hour()
	switch { // switch true: acha o primeiro case que seja verdadeiro

	case conditional < 12:
		fmt.Println("Bom dia!")
	case conditional < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite")

	}

	n := 7.4
	r := notaParaConceito(n)
	fmt.Println(r)
}

func notaParaConceito(n float64) string {

	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "Numero inválido!"
	}

}
