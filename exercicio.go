package main

import "fmt"

func analisarNotas(nota1, nota2 float64) (float64, string) {
	 media := (nota1 + nota2) / 2
	 var resultado string

	 if (media >= 6){
		resultado = "Aprovado"
	 } else {
		resultado = "Reprovado"
	 }
	 return media, resultado
	}

func main() {
	media, resultado := analisarNotas(7, 8)
	fmt.Println("Média:", media)
	fmt.Println("Você foi", resultado)
}