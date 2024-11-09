package main

import "fmt"

type Nota float64

func (n Nota) entreNotas(notaInicio, notaFim float64) bool {
	return float64(n) >= notaInicio && float64(n) <= notaFim
}
func notaParaConceito(n Nota) string {
	if n.entreNotas(9.0, 10.0) {
		return "A"

	} else if n.entreNotas(7.0, 8.99) {
		return "B"
	} else if n.entreNotas(5.0, 6.99) {
		return "C"
	} else if n.entreNotas(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}

}

func main() {

	fmt.Println(notaParaConceito(5.45))
	fmt.Println(notaParaConceito(2.45))

}
