package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	var slic []int

	// O argumento da função append precisa ser um slice, essa sentença geraria um erro:
	// arr = append(arr, 1)

	slic = append(slic, 1, 2, 3, 4, 5)
	fmt.Println(arr, slic)

	slic2 := make([]int, 2)

	copy(slic2, slic) //Copia os dois primeiros elementos

	fmt.Println(slic2)
}
