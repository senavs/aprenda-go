package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	// add(total de funções)
	wg.Add(1)

	go func1()
	func2()

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	// espera
	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
	}

	// deu
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20)

	}
}
