package main

import "fmt"

type Item struct {
	produtoID, quantidade int
	preco                 float64
}

type Pedido struct {
	userID int
	itens  []Item
}

func (p Pedido) valorTotal() float64 {
	total := 0.
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)

	}
	return total
}

func main() {

	pedido := Pedido{
		userID: 12,
		itens: []Item{
			{produtoID: 123, quantidade: 413, preco: 7.85},
			{produtoID: 124, quantidade: 13, preco: 12.15},
		},
	}

	fmt.Printf("Valor total:\t%.2f\n", pedido.valorTotal())
	fmt.Println(pedido.itens[0].preco)

}
