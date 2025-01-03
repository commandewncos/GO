package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {

	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second * 2)
		fmt.Printf("%s:\t%s\t(iteração\t%d)\n", pessoa, texto, i+1)
	}

}

func main() {

	// fale("Wilson", "Por que você não fala comigo?", 10)
	// fale("Bolão", "Só posso falar depois de você!", 1)

	// go fale("João", "Tá por ai?", 3)
	// go fale("José", "To sim...", 3)
	// fmt.Println("Hi")

	go fale("José", "Falando...", 10)
	fale("João", "Epa!", 5)

}
