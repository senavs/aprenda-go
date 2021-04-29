package main

import (
	"fmt"
)

func main() {
	// first line that will run in the program

	// Hello world
	fmt.Println("Hello world")

	// get fmt.Println output
	nBytes, error := fmt.Println("My first go code.", "Another param")
	fmt.Println(nBytes, error) // <nil> == null == None

	// how to discard variables (blank identifier) (to avoid "declared but not used")
	_, newError := fmt.Println("foo bar")
	fmt.Println(newError)

	// last line that will run in the program
}
