package main

import (
	"fmt"
)

func main() {
	// Print
	a := 10

	fmt.Print("Um print sem quebra de linha")
	fmt.Println("Um print com quebra de linha")
	fmt.Printf("Um print com formatação. a = %v\n", a)

	// Sprint (String Print)
	b := "Olá"
	c := "Tudo bem com vc?"

	sprint1 := fmt.Sprint(b, c)
	sprint2 := fmt.Sprintf(b, c)
	sprint3 := fmt.Sprintln(b, c)

	fmt.Print(sprint1)
	fmt.Print(sprint2)
	fmt.Print(sprint3)

	// Fprint (File Print)
	// fmt.Fprint
}
