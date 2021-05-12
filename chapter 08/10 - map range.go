package main

import "fmt"

func main() {
	x := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		18:  "idade de ir pra festa",
	}

	fmt.Println(x)

	for key, value := range x {
		fmt.Println(key, value)
	}

	delete(x, 123)
	fmt.Println(x)
}
