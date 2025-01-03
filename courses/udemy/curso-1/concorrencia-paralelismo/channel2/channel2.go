package main

import (
	"fmt"
	"time"
)

// Channel (canal): é a forma de comunicação entre goroutines
// channel é um tipo da linguagem

func doisTresQuatroVezes(base int, channel chan int) {
	time.Sleep(time.Second)
	channel <- 2 * base // Enviando dados para o canal

	time.Sleep(time.Second)
	channel <- 3 * base

	time.Sleep(time.Second)
	channel <- 4 * base

}
func main() {
	c := make(chan int)

	go doisTresQuatroVezes(5, c)
	fmt.Println("A")

	a, b := <-c, <-c //Recebendo dados do canal
	fmt.Println(a, b)
	fmt.Println("B")

	fmt.Println(<-c)
	fmt.Println("C")

}
