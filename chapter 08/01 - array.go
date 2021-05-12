package main

import (
	"fmt"
)

// array com tamanho 5
var x [5]int

// array com valores iniciais
// [1, 2, 3, 0, 0]
var y [5]int = [5]int{1, 2, 3}

func main() {
	x[0] = 1
	x[1] = 10
	// x[7] = 20  // error
	fmt.Println(x[0], x[1])
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	fmt.Println(len(x))

	fmt.Println(y)

}
