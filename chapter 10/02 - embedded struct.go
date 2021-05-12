package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa  // prommoted fields
	titulo  string
	salario int
}

func main() {
	p1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}
	p2 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}
	p3 := profissional{
		pessoa:  p1,
		titulo:  "Dev",
		salario: 500,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.nome)
	fmt.Println(p3)
}
