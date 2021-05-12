package main

import (
	"fmt"
)

func main() {
	ss := [][]string{
		{
			"miu",
			"milton",
			"encher o saco",
		},
		{
			"mimi",
			"martha",
			"pedir comida",
		},
		{
			"meus alunos queridos",
			"que estudam bastante",
			"fazer os exercícios ninja",
		},
	}

	for _, v := range ss {
		fmt.Println(v)
	}

	fmt.Print("\n\n")

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}

}
