package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type Ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// Para alterar o valor da propriedade é preciso passar um ponteiro da struct
func (f *Ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro := Ferrari{modelo: "F40", turboLigado: false, velocidadeAtual: 0}
	carro.ligarTurbo()
	fmt.Println(carro.turboLigado)

	var carroEsportivo Esportivo = &Ferrari{modelo: "F40", turboLigado: false, velocidadeAtual: 0}

	carroEsportivo.ligarTurbo()
	fmt.Println(carroEsportivo)

}
