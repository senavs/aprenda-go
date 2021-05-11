package main

import "fmt"

func main() {
	x := 1
	y := x << 1
	z := y >> 1

	fmt.Printf("%b %b %b\n", x, y, z)
}
