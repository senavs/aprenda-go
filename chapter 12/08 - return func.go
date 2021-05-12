package main

import "fmt"

func multiplcaPor(x int) func(int) int {
	return func(i int) int {
		return i * x
	}
}

func main() {
	f := multiplcaPor(10)

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
}
