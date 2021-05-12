package main

import "fmt"

func soma(x ...int) int {
	var result int
	for _, v := range x {
		result += v
	}
	return result
}

func somentePares(f func(x ...int) int, x ...int) int {
	var result []int

	for _, v := range x {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return f(result...)
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(somentePares(soma, 1, 2, 3, 4, 5, 6, 7, 8, 9))
}
