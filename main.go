package goarea

import "math"

/*
Pi é uma proporção numerica definida pela relação entre o perimetro de uma circunferência  e seu diâmetro
*/
const Pi = 3.1415

// Circ é responsavel por calcular a área da circunferência
func Circ(raio float64) float64 {
        return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcular a área de um retangulo
func Rect(base, altura float64) float64 {

        return base * altura

}

// Não é uma função visivel
func _TrianguloEquilatero(base, altura float64) float64 {
        return (base * altura) / 2
}
