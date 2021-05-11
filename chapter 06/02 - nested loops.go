package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 5; y++ {
			fmt.Println("x:", x, "y:", y)
		}
	}
}
