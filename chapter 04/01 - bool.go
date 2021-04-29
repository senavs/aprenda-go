package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x) // zero value

	x = true
	fmt.Println("x = true", x) // valor atribuido

	x = (10 < 100)
	fmt.Println("x = (10 < 100)", x)

	x = (10 == 100)
	fmt.Println("x = (10 == 100)", x)

	x = (10 > 100)
	fmt.Println("x = (10 > 100)", x)

	x = (10 >= 100)
	fmt.Println("x = (10 >= 100)", x)

	x = (10 <= 100)
	fmt.Println("x = (10 <= 100)", x)
}
