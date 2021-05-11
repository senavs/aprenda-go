package main

import (
	"fmt"
)

func main() {
	var a int = 200
	if a > 100 {
		fmt.Println("Hello, playground")
	}

	if x := 10; !(x > 100) {
		fmt.Println("Hello, playground")
	}
}
