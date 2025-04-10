package main

import ("fmt")

func main(){
var saldo int 
var numero int
var escolha int

fmt.Println("Insira o valor que você gostaria de sacar :")
fmt.Scan(&saldo)

    fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Sacar")
	fmt.Println("2 - Depositar")
	fmt.Print("Opção: ")
	fmt.Scan(&escolha)

	if escolha == 1 {
		fmt.Print("Insira o valor que você gostaria de sacar: ")
		fmt.Scan(&numero)
		if numero > saldo {
			fmt.Println("Saldo insuficiente.")
		} else {
			saldo -= numero
			fmt.Println("Saque realizado com sucesso!")
		}
	} else if escolha == 2 {
		fmt.Print("Insira o valor que você gostaria de depositar: ")
		fmt.Scan(&numero)
		saldo += numero
		fmt.Println("Depósito realizado com sucesso!")
	} else {
		fmt.Println("Opção inválida.")
	}

	fmt.Println("Saldo atual:", saldo)
}
	
