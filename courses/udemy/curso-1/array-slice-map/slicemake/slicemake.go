package main

import "fmt"

func main() {

	/*
		O slice tem um ponteiro para um array, ou seja, ele retorna os valores referenciados
		Criando um slice com a funão make
		Parametros: tipo do slice, posições
	*/

	s := make([]int, 10)
	s[7] = 35
	fmt.Println(s)

	s = make([]int, 10, 20)

	fmt.Println(s, "Tamanho:", len(s), "Capacidade:", cap(s)) // cap: capacidade do array interno

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(s)
	fmt.Println(s, "Tamanho:", len(s), "Capacidade:", cap(s))

	//Ele não gera erro caso você acrescente um valor, ele dobrará de tamanho:
	s = append(s, 10, 11, 23)
	fmt.Println(s, "Tamanho:", len(s), "Capacidade:", cap(s))

}
