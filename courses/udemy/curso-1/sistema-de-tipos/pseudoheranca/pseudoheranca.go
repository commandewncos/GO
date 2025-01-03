package main

import "fmt"

type Carro struct {
	Nome            string
	velocidadeAtual int
}

type Ferrari struct {
	Carro      // Campo anonino (Não por herença e sim por composição)
	tudoLigado bool
}

func main() {
	f := Ferrari{}
	f.Nome = "F40"
	f.velocidadeAtual = 0
	f.tudoLigado = true
	// Percentual (t) '%t' valor booleano
	fmt.Printf("A Ferrari %s está com o turbo ligado?\t%t\n", f.Nome, f.tudoLigado)
	fmt.Println(f)
}
