package main

import (
	"fmt"
)

var x int

// x = 10  // syntax error: non-declaration statement outside function body

func main() {
	fmt.Println(x)

	x = 10
	fmt.Println(x)
}
