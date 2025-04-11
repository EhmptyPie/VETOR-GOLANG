package main

import (
	"fmt"
)

var saldo float32

func saque() {
	var valorSaque float32
	fmt.Println("Seu saldo atual é:", saldo)
	fmt.Print("Quanto você quer sacar? ")
	fmt.Scan(&valorSaque)

	if valorSaque > saldo {
		fmt.Println("Saldo insuficiente!")
	} else {
		saldo -= valorSaque
		fmt.Println("Saque realizado com sucesso!")
		fmt.Println("Seu saldo atual é:", saldo)
	}
}

func deposito() {
	var valorDeposito float32
	fmt.Println("Seu saldo atual é:", saldo)
	fmt.Print("Quanto você quer depositar? ")
	fmt.Scan(&valorDeposito)

	if valorDeposito <= 0 {
		fmt.Println("Valor de depósito inválido!")
	} else {
		saldo += valorDeposito
		fmt.Println("Depósito realizado com sucesso!")
		fmt.Println("Seu saldo atual é:", saldo)
	}
}

func main() {
	saldo = 10000
	var opcao int

	for {
		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1 - Sacar")
		fmt.Println("2 - Depositar")
		fmt.Println("3 - Sair")
		fmt.Print("Opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			saque()
		case 2:
			deposito()
		case 3:
			fmt.Println("Saindo... Obrigado!")
			return
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}
