package main

import (
	"fmt"
)

type myType int

var x myType

func main() {
	fmt.Printf("Valor: %v\n", x)
	fmt.Printf("Tipo: %T\n", x)

	x = 42
	fmt.Printf("Novo valor: %v\n", x)

	// types https://golang.org/ref/spec#Types
}
