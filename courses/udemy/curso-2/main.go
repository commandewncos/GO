package main

import (
	"fmt"
	"time"
)

func main() {
	// addr := gostruct.Address{City: "São Paulo", Street: "Rua Projetada", Number: 15}
	// pers := gostruct.Person{Name: "Wilson", Nascimento: time.Date(1989, 06, 24, 0, 0, 0, 0, time.Local), Address: addr}
	// pers.Email = "wilsonnascimentocosta98@gmail.com"

	// fmt.Println(pers)
	// pers.ChangeEmail("Wilson.Costa@gmail.com")
	// fmt.Println(pers)

	// c := compra.Compras{Mercado: "Batalha", Data: time.Date(2024, 8, 12, 0, 0, 0, 0, time.Local), Item: []compra.Items{{Name: "Banana", Preco: 29.90}, {Name: "Uva", Preco: 15.5}}}

	// fmt.Println(c)

	func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}()

}
