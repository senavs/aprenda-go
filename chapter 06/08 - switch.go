package main

import "fmt"

func main() {
	x := 100

	switch {
	case x < 5:
		fmt.Println("X é menor que 5")
	case x == 5:
		fmt.Println("X é igual 5")
	case x > 5 && x <= 10:
		fmt.Println("X é maior que 5")
	default:
		fmt.Println("X é maior que 10")
	}

	y := "b"

	switch y {
	case "a":
		fmt.Println("Y é a letra A")
	case "b":
		fmt.Println("Y é a letra B")
		fallthrough
	case "c", "d", "e":
		fmt.Println("Y é a letra C, D ou E")
	default:
		fmt.Println("Y é qualquer outra coisa")
	}
}
