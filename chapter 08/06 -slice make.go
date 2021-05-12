package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5
	fmt.Println(slice, len(slice), cap(slice))

	for i := 6; i <= 10; i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 11)
	fmt.Println(slice, len(slice), cap(slice))
}
