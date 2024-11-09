package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	nome, sobrenome string
}

func (p Pessoa) getNomeCompleto() string {
	return strings.Join([]string{p.nome, p.sobrenome}, " ")
}

func (p *Pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome, p.sobrenome = partes[0], partes[1]
}

func main() {

	pessoa := Pessoa{nome: "Wilson", sobrenome: "Costa"}

	fmt.Println(pessoa.getNomeCompleto())
	pessoa.setNomeCompleto("Wilson Nascimento")
	fmt.Println(pessoa.getNomeCompleto())

}
