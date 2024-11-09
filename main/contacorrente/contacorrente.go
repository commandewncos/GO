package contacorrente

import (
	"fmt"

	"github.com/wilson/golang/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldoConta    float64
}

// SALDO
func (saldo *ContaCorrente) Saldo() {
	fmt.Println("Saldo:", saldo.saldoConta)
}

// SACAR
func (s *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque <= 0 {
		fmt.Println("\nValor incorreto!")

	} else if s.saldoConta >= valorDoSaque && valorDoSaque > 0 {
		s.saldoConta -= valorDoSaque
		fmt.Println("\nOperação realizada com sucesso!")
	} else {
		fmt.Println("\nSaldo insuficiente!\b")
	}

	return "Saldo:", s.saldoConta
}

// TRANFERENCIA
func (t *ContaCorrente) Transferir(valor float64, conta *ContaCorrente) bool {
	if valor < t.saldoConta && valor > 0 {
		t.saldoConta -= valor
		conta.Depositar(valor)
		fmt.Println("Saldo", t.saldoConta)
		return true

	} else {
		fmt.Println("Sem saldo!")
		return false
	}

}

// DEPOSITAR
func (d *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		d.saldoConta = d.saldoConta + valorDoDeposito
		fmt.Println("\nOperação realizada com sucesso!")

	} else {
		fmt.Println("\nValor incorreto!")

	}
	return "Saldo:", d.saldoConta
}
