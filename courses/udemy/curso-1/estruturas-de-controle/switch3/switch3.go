package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {

	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "I don't know"
	}

}

func main() {

	fmt.Println(tipo(3.2))
	fmt.Println(tipo("Wilson"))
	fmt.Println(tipo(5))
	fmt.Println(tipo(sum))
	fmt.Println(tipo(time.Now()))
}

func sum() {

}
