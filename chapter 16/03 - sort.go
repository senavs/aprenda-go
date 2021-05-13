package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"banana", "abacaxi", "pêra", "laranja", "banana"}
	si := []int{7, 5, 2, 5, 8, 5, 1, 7, 9, 2, 4, 6, 1, 2}

	fmt.Println(ss)
	fmt.Println(si)
	fmt.Println(sort.IntsAreSorted(si))

	sort.Strings(ss)
	sort.Ints(si)

	fmt.Println(ss)
	fmt.Println(si)

	fmt.Println(sort.IntsAreSorted(si))
}
