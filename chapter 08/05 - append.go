package main

import (
	"fmt"
)

func main() {
	slice1 := []int{0, 1, 2, 3}
	slice2 := []int{10, 11, 12, 13}

	slice := append(slice1, slice2...)
	slice = append(slice, 20, 21, 22, 23)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice)
}
