package main

import (
	"fmt"
)

func main() {
	s := "Olá, mundo! 坔"
	sb := []byte(s)

	fmt.Printf("%v\n%T\n", s, s)
	fmt.Printf("%v\n%T\n", sb, sb)

	// por caracter
	for _, v := range s {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}

	fmt.Println("")

	// por byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i], s[i])
	}
}
