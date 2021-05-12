package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	c1 := cliente{
		nome:      "João",
		sobrenome: "da Silva",
		fumante:   false,
	}
	c2 := cliente{"Joana", "Pereira", true}

	fmt.Println(c1)
	fmt.Println(c2)
}
