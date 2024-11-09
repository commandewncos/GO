package main

import "fmt"

func main() {

	ch := make(chan int, 1)
	ch <- 1 //Enviando um valor para o canal (processo de escrita)
	<-ch    // Recebendo dados do canal (processo de leitura)
	ch <- 2
	fmt.Println(<-ch)

}
