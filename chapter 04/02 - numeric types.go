package main

import (
	"fmt"
	"runtime"
)

func main() {
	a, b, c := "a", "á", "癃"
	fmt.Println(a, b, c)

	d, e, f := []byte(a), []byte(b), []byte(c)
	fmt.Println(d, e, f)

	x, y := 10, 10.5
	fmt.Printf("%T, %T\n", x, y)

	fmt.Println(runtime.GOOS)   // Operational System
	fmt.Println(runtime.GOARCH) // Architecture (32 or 64)

	// int == int64 in 64 architectures
	// int == int32 in 32 architectures

	// byte == int8
	// rune == int32
}
