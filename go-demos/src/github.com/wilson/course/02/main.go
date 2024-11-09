package main

import "fmt"

func main() {
	var name string = "Wilson"
	var age int = 35
	const isTrue = true

	fmt.Printf("%T\n\n", name)
	fmt.Println(name, isTrue, age)

}
