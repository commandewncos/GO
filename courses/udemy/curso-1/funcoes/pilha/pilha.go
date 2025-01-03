package main

import (
	"runtime/debug"
)

func funcao() {

	debug.PrintStack()
}

func funcao1() {
	funcao()

}

func funcao2() {
	funcao1()

}

func main() {
	funcao2()

}
