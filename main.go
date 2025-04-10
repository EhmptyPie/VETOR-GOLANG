package main

import ("fmt")

func main(){
	var numeros [5]int
	var soma int

	fmt.Println("Digite 5 números inteiros:")


	for i := 0; i < 5; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numeros[i])
		soma += numeros[i]
	}
	fmt.Printf("A soma dos números digitados é: %d\n", soma)
}