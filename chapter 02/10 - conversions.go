package main

import (
	"fmt"
)

type hotdog int

var b hotdog = 10

func main() {
	x := 10

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", b, b)

	// x = b  // error: cannot use b (type hotdog) as type int in assignment
	// b = x  // error: cannot use x (type int) as type hotdog in assignment

	x = int(b)
	b = hotdog(x)
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", b, b)
}
