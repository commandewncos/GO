package main

import (
	"fmt"

	"github.com/Wilson-cmd/html"
)

func encaminhar(oringem <-chan string, destino chan string) {
	for {

		destino <- <-oringem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(html.Titulo("https://google.com", "https://alura.com.br"), html.Titulo("https://youtube.com"))

	fmt.Println("Primeiro mais rapido:", <-c, "\nSegundo:", <-c, "\nTerceiro:", <-c)
}
