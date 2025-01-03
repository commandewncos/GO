package main

import "fmt"

func funcao1() {
	fmt.Println("F!")
}

func funcao2(param1, param2 string) {
	fmt.Printf("Funcao 2: %s %s\n", param1, param2)

}

func funcao3() string {
	return "Hello!"
}

func funcao4(param1, param2 string) string {
	return fmt.Sprintf("Funcao 4: %s %s\n", param1, param2)
}

func funcao5() (string, string) {
	return "Wilson", "Costa"
}

func main() {
	funcao1()
	funcao2("A", "B")
	fmt.Println(funcao3())
	fmt.Println(funcao4("X", "Y"))
	nome, sobrenome := funcao5()
	fmt.Println(nome, sobrenome)
}
