package main

import (
	"fmt"
)

// package level scope
var y = 20

func main() {
	z := 10

	printInt(z)
}

func printInt(x int) {
	fmt.Println(x)
	fmt.Println(y)
	// fmt.Println(z)  // error: undefined
}
