package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct para json

	produto := Produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 1589.49,
		Tags:  []string{"Promoção", "Eletrônico"},
	}

	produtoJSON, _ := json.Marshal(produto)

	fmt.Println(string(produtoJSON))

	var outroProduto Produto

	jsonString := `
		{
			"id":2,
			"nome": "Caneta",
			"preco": 8.90,
			"tags": ["Papelaria", "Importando"]
		}
	`
	json.Unmarshal([]byte(jsonString), &outroProduto)

	fmt.Println(outroProduto.Tags)

}
