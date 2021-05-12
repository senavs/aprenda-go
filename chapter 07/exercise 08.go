package main

import (
	"fmt"
)

func main() {
	tamanhodocansaço := 2

	switch {

	case tamanhodocansaço == 0:
		fmt.Println("que malandragem")
	case tamanhodocansaço == 1:
		fmt.Println("uma gelada ia bem")
	case tamanhodocansaço == 2:
		fmt.Println("ih já era, só nascendo denovo")
	}

}
