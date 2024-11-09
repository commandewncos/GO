package main

import "fmt"

func main() {
	// Arrays

	var arr [2]string

	arr[0] = "Wilson"
	arr[1] = "Jessica"

	fruitArr := [3]string{"Banana", "Apple", "Orange"}
	for i := 0; i < len(fruitArr); i++ {
		fmt.Println(fruitArr[i])
	}

	fmt.Println(fruitArr, arr)
}
