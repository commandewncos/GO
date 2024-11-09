package main

import (
	"fmt"
	"math"
	"reflect"
)

var (
	a, b, c, d = 10, false, "Hello", 89.7
)

func main() {

	fmt.Println("Tipos de a:", reflect.TypeOf(a))
	fmt.Println("Tipos de b:", reflect.TypeOf(b))
	fmt.Println("Tipos de c:", reflect.TypeOf(c))
	fmt.Println("Tipos de c:", reflect.TypeOf(d))

	/*
		Só sinais positivos, ou seja, não reserva espaço para os sinais.
		Tipos de inteiros. (u: unsighted)

		1 - byte:	uint8	-> 1byte:	8	bits
		2 - short:	uint16	-> 2bytes:	16	bits
		3 - int:	uint32	-> 4Bytes:	32	bits
		4 - long: 	uint64	-> 8bytes:	64	bits
	*/

	var x byte = 255 //Alias para (uint8)
	fmt.Println("Reflect: O byte é:", reflect.TypeOf(x))

	/*
		Inteiros com sianais
		int8, int16, int32, int64
	*/

	y := math.MaxInt64
	fmt.Println("Valor maximo:", y)
	fmt.Println("Valor maximo:", reflect.TypeOf(y)) //-> Arquitetura de 64bits o valor maximo.

	// Tipo rune
	// Representa um mapeamento da tabela unicode (int32)
	// Representação inteira da tabela unicode

	var a rune = 'a'

	fmt.Println("Type:", reflect.TypeOf(a))
	fmt.Println("Representação de (a) em unicode (rune):", a)

	var f float32 = 56.66
	var g float64 = 89.23
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(g), "Tipo literal:", reflect.TypeOf(29.87))

	// Boolean

	bo := true
	fmt.Println("o tipo de bo:", reflect.TypeOf(bo))
	fmt.Println("Negação do valor:", !bo)

	// Ver o tamanho de uma string:

	fmt.Println("Tamanho do nome completo:", len("Wilson Nascimento Costa"))

	//Utilizando acento de crase
	fmt.Println("Tamanho do nome completo:", len(`Wilson Nascimento Costa`))

	// char
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
