package main

import (
	"fmt"
	"time"
)

func exmain() {
	a, b := make(chan int), make(chan int)
	var x int = 250

	go receivechannel((x / 2), a)
	go receivechannel((x / 3), b)

	for i := 0; i <= x; i++ {
		select {
		case value := <-a:
			fmt.Printf("Channel A:\t%d\n", value)
		case value := <-b:
			fmt.Printf("Channel B:\t%d\n", value)
		}
	}
	close(a)
	close(b)
}

func receivechannel(z int, y chan int) {
	for i := 0; i <= z; i++ {
		y <- i
	}
	time.Sleep(2 * time.Second)
}
