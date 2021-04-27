package main

import (
	"fmt"
)

type myType int

var x myType
var y int

func main() {
	fmt.Printf("Valor: %v\n", x)
	fmt.Printf("Tipo: %T\n", x)

	x = 42
	fmt.Printf("Novo valor: %v\n", x)

	y = int(x)
	fmt.Printf("Valor de x: %v\n", x)
	fmt.Printf("Valor de y: %v\n", y)
}
