package main

import "fmt"

func main() {
	x, y := 1, 10

	if x >= y {
		fmt.Printf("%d is more than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", x, y)
	}

	color := "white"

	switch color {
	case "orange":
		fmt.Println("Color is", color)
	case "red":
		fmt.Println("Color is", color)
	case "blue":
		fmt.Println("Color is", color)
	case "green":
		fmt.Println("Color is", color)
	default:
		fmt.Println("I don't know this color")
	}
}
