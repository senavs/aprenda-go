package main

import "fmt"

func main() {
	slice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(slice)
	fmt.Println(slice[2][1])
}
