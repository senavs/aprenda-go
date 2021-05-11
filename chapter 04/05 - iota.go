package main

import "fmt"

const (
	a = iota
	b = iota
	_ = iota
	c = iota
)

const (
	x = iota * 10
	y
	z
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
}
