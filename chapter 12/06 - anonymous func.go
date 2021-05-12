package main

import "fmt"

func main() {
	x := 10

	func(y int) {
		fmt.Println("O valor passado foi", y)
	}(x)
}
