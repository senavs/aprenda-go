package main

import "fmt"

func somaMuitos(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x, _ := somaMuitos()
	y, _ := somaMuitos(s...)

	fmt.Println(x)
	fmt.Println(y)
}
