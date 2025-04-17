package main

import ("fmt")


func dividir (dividendo int, divisor int) (int, string){
	if divisor == 0{
		return 0, "Erro na divisão por zero"
	}
	return dividendo / divisor, "Sem erro"

}
func operadocaoBasica( a int, b int) (int, int, int){
	soma:= a + b
	subtracao:= a - b
	multiplicacao:= a * b
	return soma, subtracao, multiplicacao
}
func main() {
	resultado, erro := dividir(10,2)
	if erro != "Sem erro"{
		fmt.Println(erro)
	}else{
		fmt.Println("Resultado: ", resultado)

	}

	soma, sub, multip := operadocaoBasica(10, 2)
	fmt.Println(soma)
	fmt.Println(sub)
	fmt.Println(multip)
}
