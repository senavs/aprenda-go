package main

import "fmt"

func main() {
	// for <bool> == while <bool>
	var x int
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for {
		// infinity loop
	}
}
