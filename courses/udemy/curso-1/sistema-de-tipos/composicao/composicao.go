package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type Luxuoso interface {
	fazerBaliza()
}

type EsportivoLuxuoso interface {
	Esportivo
	Luxuoso

	// Incluir a composição dos demais metodos ou interfaces
}

type BMW struct {
	modelo string
}

func (b BMW) ligarTurbo() {
	fmt.Println("Turbo ligado!")
}

func (b BMW) fazerBaliza() {
	fmt.Println("Balizando...")
}

func main() {
	var b EsportivoLuxuoso = BMW{modelo: "Seven"}

	b.fazerBaliza()
	b.ligarTurbo()
}
