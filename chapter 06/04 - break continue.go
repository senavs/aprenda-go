package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		if x == 5 {
			continue
		}
		fmt.Println(x, "meio do loop")
		if x == 7 {
			break
		}

	}
}
