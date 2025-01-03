package main

import "fmt"

// NAO TEM TERNARIO
func main() {
	nota := 7.
	resultado := obterResultado(nota)

	fmt.Println(resultado)

}

func obterResultado(nota float64) string {
	if nota >= 7 {
		return "Aprovado"
	} else if nota >= 5 {
		return "Recuperação"
	} else {
		return "Reprovado"
	}
}
