package main

import (
	"fmt"
	"strings"
)

type imprimivel interface {
	toString() string
}

type Pessoa struct {
	nome, sobrenome string
}

type Produto struct {
	nome  string
	preco float64
}

// Interfaces são implementadas implicitamente

func (p Pessoa) toString() string {
	return strings.Join([]string{p.nome, p.sobrenome}, " ")

}

func (p Produto) toString() string {
	return fmt.Sprintf("%s - R$%.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {

	var coisa imprimivel = Pessoa{nome: "Wilson", sobrenome: "Costa"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = Produto{nome: "Caneta", preco: 5.74}
	fmt.Println(coisa.toString())
}
