package main

import "fmt"

func teste1() {
	defer fmt.Println("1. com defer")
	defer fmt.Println("2. com defer")
	fmt.Println("3. sem defer")
}

func teste2() {
	var x int

	defer fmt.Println("1", x)

	x = 10
	fmt.Println("2", x)
}

func main() {
	teste1()
	fmt.Println()
	teste2()
}
