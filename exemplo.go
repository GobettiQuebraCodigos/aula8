package main

import "fmt"

func main(){
	votos := make(map[string]int)

	votos["Ana"] += 1
	votos["Carlos"] += 1
	votos["João"] += 1
	votos["Ana"] += 1
	votos["Ana"] += 1
	

	for candidato, total := range votos {
		fmt.Printf("%s recebeu %d votos\n", candidato, total)
	}

}

