package main

import "fmt"

func print(x interface{}) {
	fmt.Println(x)
}

func soma(x, y int) int {
	return x + y
}

func somaMuitos(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}

func main() {
	var x int = soma(1, 1)
	var y, yTotal int = somaMuitos(1, 2, 3, 4, 5, 6, 7, 8, 9)

	print(x)
	print(y)
	print(yTotal)
}
