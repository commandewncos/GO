package main

import "fmt"

func main() {

	a := 10
	b := &a // 'b' is a pointer of variable 'a'

	fmt.Println("Value of variable:\t", a)
	fmt.Println("Memory address:\t", b)

	// Use to read value of memory address: Pointer

	fmt.Println("Read value from address:\t", *b)

	// Change value with pointer

	*b = 45
	fmt.Println(a)

}
