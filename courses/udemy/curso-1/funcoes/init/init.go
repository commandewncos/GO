package main

import "fmt"

// Função init é um convenção em GO para ser executada na inicialização, sem a necessidade de chamar no escopo do main.

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
