package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{

		"José Maria":   11325.45,
		"Maria José":   15462.78,
		"Rocha D'agua": 11457.77,
	}

	funcsESalarios["Filho José"] = 1350.50

	fmt.Println(funcsESalarios)
	delete(funcsESalarios, "Inexistente")
	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios {
		fmt.Printf("%s\t%f\n", nome, salario)
	}

}
