package main

import "fmt"

func main(){
	nomes := []string{"Ana", "Bruno", "Carlos", "Diana", "Eduardo"}

    fmt.Println("Dois primeiros:", nomes[:2])
    fmt.Println("Dois últimos:", nomes[len(nomes)-2:])
    fmt.Println("Nome do meio:", nomes[len(nomes)/2:len(nomes)/2+1])

}
