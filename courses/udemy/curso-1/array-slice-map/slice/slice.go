package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [3]int{1, 2, 3} //Array
	b := []int{1, 2, 3}  //Slice

	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	arr := [5]int{1, 2, 3, 4, 5} // Array com 5 posições

	//Slice não é um array
	// Slice define um pedaço de um array

	slc := arr[1:3]
	slc2 := arr[:2]
	fmt.Println(slc, slc2)

}
