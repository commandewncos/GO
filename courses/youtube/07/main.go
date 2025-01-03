package main

import (
	"fmt"
)

func main() {

	// Variables
	// i := 1
	// name := "Wilson"

	// Method
	// for i <= 10 {
	// 	if i <= len(name) {
	// 		fmt.Println(i, "\t", name)
	// 	} else {
	// 		break
	// 	}

	// 	i++
	// }

	// Short method
	// for i := 1; i <= 20; i++ {
	// 	fmt.Println(i)
	// }

	for i := 1; i <= 100; i++ {

		if checkModule(i, 15) {
			fmt.Println("FizzBuzz")

		} else if checkModule(i, 3) {
			fmt.Println("Fizz")
		} else if checkModule(i, 5) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}

}

func checkModule(num1, num2 int) bool {
	x := ((num1 % num2) == 0)
	return x

}
