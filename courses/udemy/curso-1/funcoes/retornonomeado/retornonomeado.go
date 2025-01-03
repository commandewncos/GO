package main

import "fmt"

func trocar(param1, param2 int) (segundo int, primeiro int) {
	segundo = param2
	primeiro = param1

	return //Retorno limpo

}

func main() {

	r, rr := trocar(3, 7)
	fmt.Println(r, rr)
}
