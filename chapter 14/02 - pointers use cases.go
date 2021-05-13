package main

import (
	"fmt"
)

func main() {

	x := 11

	//incValor(x)
	incValorPonteiro(&x)

	fmt.Println(x)

}

func incValor(x int) {
	x++
	fmt.Println("Na função:", x)

}

func incValorPonteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}
