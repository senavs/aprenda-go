package main

import "fmt"

func main() {
	var x interface{}

	x = true

	switch x.(type) {
	case nil:
		fmt.Println("X é do tio nil")
	case float64:
		fmt.Println("X é do tio float64")
	case int:
		fmt.Println("X é do tio int")
	case string:
		fmt.Println("X é do tio string")
	case bool:
		fmt.Println("X é do tio bool")
	default:
		fmt.Println("X é desconhecido", x)
	}
}
