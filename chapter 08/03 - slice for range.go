package main

import "fmt"

func main() {
	slice := []string{"banana", "maçã", "jaca"}

	for index, value := range slice {
		fmt.Println("Index:", index, "there is the value:", value)
	}

	// slice[3] = "pêssego"  // error array subjacente é de 3 elementos
	// []string{"banana", "maçã", "jaca"} from [3]string{"banana", "maçã", "jaca"}
	// use o append

	slice = append(slice, "pêssego")

	for _, value := range slice {
		fmt.Println("Value:", value)
	}
}
