package main

import (
	"fmt"
)

var arr [5]int = [5]int{1, 2, 3, 4, 5}
var slc []int = []int{1, 2, 3, 4, 5}

func main() {
	fmt.Println(arr)
	fmt.Println(slc)

	slc := append(slc, 6)
	fmt.Println(slc)

	// slc[20] = 19 // error
}
