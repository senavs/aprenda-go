package main

import "fmt"

func main() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Maior",
		idade: 50,
	}

	fmt.Println(x)
}
