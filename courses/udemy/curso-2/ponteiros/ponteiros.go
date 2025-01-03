package main

import "fmt"

func main() {
	x := 5
	y := &x

	*y = 10

	fmt.Println(&x == y)
}
