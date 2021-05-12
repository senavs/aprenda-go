package main

import "fmt"

func main() {
	x := 10

	f := func(y int) {
		fmt.Println("O valor passado foi", y)
	}

	f(x)
}
