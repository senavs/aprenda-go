package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"alfredo": 5550111,
		"joana":   5550222,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 5550333

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])

	// comma ok idiom
	if será, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(será)
	}
}
