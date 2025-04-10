package main

import ("fmt")

func main(){
	var idade int
	fmt.Println("Quantos anos você tem? ")
	fmt.Scan(&idade)
	
	if idade < 18 {
		fmt.Println("Você é menor de idade")
	}else if idade <= 60 {
	fmt.Println("Você é adulto")
	}else{
		fmt.Println("Você é idoso")
	}
}
