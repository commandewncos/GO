package main

import (
	"fmt"
)

func main() {

	// ids := []int{33, 56, 78, 24, 31, 56, 98}

	// for i, id := range ids {
	// 	fmt.Printf("INDEX: %d\tID: %d\n", i, id)
	// }

	// for i, letter := range "WILSON NASCIMENTO COSTA" {
	// 	fmt.Println(i, "\t", string(letter))
	// }

	// sum := 0
	// num := []int{10, 20, 30, 50}
	// for _, n := range num {
	// 	sum += n

	// }

	// fmt.Println("Sum of list is equal:\t", sum)

	emails := map[string]string{"Wilson": "wilson@gmail.com", "Jessica": "jessica@gmail.com"}

	for key, value := range emails {
		fmt.Printf("%s:\t%s\n", key, value)
	}

}
