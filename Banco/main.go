package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return fmt.Sprintf("Saque no valor de %.2f R$ realizado com sucesso", valorDoSaque)
	} else {
		return "Saldo insuficiente"
	}
}
func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return fmt.Sprintf("Deposito no valor de %.2f R$ realizado com sucesso", valorDoDeposito)
	} else {
		return "Insira um número positivo"
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(300))
	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Depositar(200))
	fmt.Println(contaDaSilvia.saldo)

}
