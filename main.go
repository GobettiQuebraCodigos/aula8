package main

import "fmt"

func pegarNome() (string, string) {
	return "Lucas", "Gobetti"
}

func main() {
	nome, sobrenome := pegarNome()
	fmt.Println(nome, sobrenome)
}