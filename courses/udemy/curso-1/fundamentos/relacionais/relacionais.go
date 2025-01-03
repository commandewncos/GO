package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings: Igualdade", "Banana" == "Maça")
	fmt.Println("Diferente: !=", 1 != 2)
	fmt.Println("Menor que: <", 5 < 9)
	fmt.Println("Maior que: >", 5 > 9)
	fmt.Println("Maior igual que: >=", 5 >= 9)
	fmt.Println("Menor igual que: <=", 5 <= 9)

	data1 := time.Unix(0, 0)
	data2 := time.Unix(0, 0)

	fmt.Println("Datas:", data1 == data2)
	fmt.Println("Datas, equal", data1.Equal(data2))

	type Person struct {
		Name string
	}

	personOne := Person{Name: "João"}
	personTwo := Person{Name: "Maria"}

	fmt.Println("Pessoas:", personOne == personTwo)
}
