package main

import "math"

// Iniciando métodos, structs, funções ..., por convenção será considerado como PÚBLICO
// Iniciando com letra minusculo, por convenção, é considerado PRIVADO dentro do pacote.

/*
	Por exemplo:
	iniciado com (.) "Ponto" ou (-) "Traço": gerará  algo público
	Iniciado com (.) "Pinto" ou (_) "underline" gerará algo privado.

*/

type Ponto struct {
	x, y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return
}

func Distancia(a, b Ponto) float64 {
	// Distancia é responsavel por calcular a distancia linear entre dois pontos
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
