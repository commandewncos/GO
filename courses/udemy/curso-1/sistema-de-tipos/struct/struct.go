package main

import "fmt"

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: funcão com receiver (receptor)
func (p Produto) precoDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	novoProduto := Produto{
		nome:     "Lapis",
		preco:    1.50,
		desconto: 0.05,
	}

	produto2 := Produto{"Caneta", 2.50, 0.03}

	fmt.Println(novoProduto)
	fmt.Printf("%.2f\n", novoProduto.precoDesconto())

	fmt.Printf("%.2f\n", produto2.precoDesconto())
}
