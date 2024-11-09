package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {

	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i+1)

		}
	}()
	return c
}

func juntar(ent1, ent2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {

			case s := <-ent1:
				c <- s
			case s := <-ent2:
				c <- s
			}
		}
	}()

	return c
}

func main() {

	c := juntar(falar("João"), falar("Maria"))

	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)

}
