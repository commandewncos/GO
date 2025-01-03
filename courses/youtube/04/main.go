package main

import "fmt"

func main() {
	name := "Wilson"
	age := 35
	fmt.Println(greeting(name, age))
}

func greeting(name string, age int) (string, int) {
	return name, age
}
