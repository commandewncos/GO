package main

import "fmt"

func main() {

	s := make([]int, 10, 20)
	s0 := append(s, 1, 2, 3)

	fmt.Println(s, s0)

	s[1] = 7

	fmt.Println(s, s0)
}
