package main

import "fmt"

func main() {
	produtos := make(map[string]int)

	produtos["Coxinha"] = 10
	produtos["Pão de queijo"] = 15
	produtos["Refrigerante"] = 20

	for alimento, total := range produtos {
		fmt.Printf("%s no estoque: %d\n", alimento, total)
	}

	produtos["Coxinha"] -= 2
	produtos["Pão de queijo"] -= 1

	fmt.Println("------------------")

	for alimento, total := range produtos {
		fmt.Printf("%s no estoque: %d\n", alimento, total)
	}
}