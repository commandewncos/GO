package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Variavel de inicialização antes do bloco.
	if i := numeroAleatorio(); i >= 5 {
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}
	numeroAleatorio()
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)

}
