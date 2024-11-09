package main

import "fmt"

/*
Caso os dois trabalhos sejam satisfatorios, então comprará a TV de 50 polegadas e comprará Sorvete.
Caso UM DOS dois trabalhos sejam satisfatorios, então comprará a TV de 32 polegadas e comprará Sorvete.
Caso NÃO OCORRÁ dois trabalhos sejam satisfatorios, então sem TV e Sorvete.
*/

func comprar(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTV50Polegadas := trabalho1 && trabalho2
	comprarTV32Polegadas := trabalho1 != trabalho2 //OU exclusivo
	compraSorvete := trabalho1 || trabalho2

	return comprarTV50Polegadas, comprarTV32Polegadas, compraSorvete
}

func main() {
	// tv50, tv32, sorvete := comprar(true, true)
	// fmt.Printf("TV50: %t, TV32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)

	// tv50, tv32, sorvete := comprar(false, true)
	// fmt.Printf("TV50: %t, TV32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete := comprar(false, false)
	fmt.Printf("TV50: %t, TV32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)

}
