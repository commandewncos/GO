package main

import (
	"fmt"
	"os"

	c "github.com/wilson/golang/contacorrente"
)

func main() {
	usuarioContaCorrente := new(c.ContaCorrente)
	usuarioContaCorrente.Titular.Nome = "Wilson Nascimento Costa"
	usuarioContaCorrente.NumeroAgencia = 1456
	usuarioContaCorrente.NumeroConta = 256

	for {

		var condicao int
		fmt.Println("\nEscolha a operação que deseja realiazar:")
		fmt.Println("0:\tSaldo\n1:\tSacar\n2:\tDepositar\n3:\tTranferir\n4:\tSair\n\b")

		fmt.Scan(&condicao)

		switch condicao {
		case 0:
			{
				usuarioContaCorrente.Saldo()
			}

		case 1:
			var valor float64
			fmt.Println("\nInforme o valor que deseja sacar:")
			fmt.Scan(&valor)
			fmt.Println(usuarioContaCorrente.Sacar(valor))
		case 2:
			var valor float64
			fmt.Println("Informe o valor que deseja despositar:")
			fmt.Scan(&valor)
			fmt.Print(usuarioContaCorrente.Depositar(valor))

		case 3:
			var valor float64
			operacaoContaCorrenteUsuario := new(c.ContaCorrente)

			fmt.Println("Informe o nome:")
			fmt.Scan(&operacaoContaCorrenteUsuario.Titular.Nome)
			fmt.Println("Informe o numero da agencia:")
			fmt.Scan(&operacaoContaCorrenteUsuario.NumeroAgencia)
			fmt.Println("Informe o numero da conta:")
			fmt.Scan(&operacaoContaCorrenteUsuario.NumeroConta)
			fmt.Println("Informe o numero da valor:")
			fmt.Scan(&valor)
			conta := operacaoContaCorrenteUsuario
			usuarioContaCorrente.Transferir(valor, conta)

		case 4:
			fmt.Println("\nTchaul!...")
			os.Exit(-1)

		default:
			fmt.Println("\nOpção invalida!")

		}
	}

}
