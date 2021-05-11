package main

import (
	"fmt"
)

func main() {
	if x := 500; x > 100 {
		fmt.Println("chis é maior que cem")
	} else if x < 10 {
		fmt.Println("chis é menor que déis")
	} else {
		fmt.Println("chis não é menor que déis nem maior que cem")
	}
}
