package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Excutou!") // Nessa linha o buffer não tinha mais espaço
	ch <- 5

}

func main() {
	ch := make(chan int, 3) //Buffer com 3 posições

	go rotina(ch)
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
