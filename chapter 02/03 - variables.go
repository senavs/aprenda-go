package main

import (
	"fmt"
)

func main() {
	// declaration
	a := 10
	b := 10.
	c := "Olá, bom dia"
	x, y, z := 1, 2 != 2, 3 == 3

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)

	// assignment
	a = 20
	b = b * 10
	// c := "new string"  // error: no new variables on left side of :=

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
}
