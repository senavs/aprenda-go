package main

import (
	"fmt"
)

func main() {
	x := 10
	y := &x

	fmt.Printf("%v\t%v\t%T\n", &x, x, x)
	fmt.Printf("%v\t%v\t%T\n", &y, y, y)

	*y++

	fmt.Printf("%v\t%v\t%T\n", &x, x, x)
	fmt.Printf("%v\t%v\t%T\n", &y, y, y)
}
