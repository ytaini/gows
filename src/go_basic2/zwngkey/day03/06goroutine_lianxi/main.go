package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(20)
	go a()
	go b()
	wg.Wait()
}

func a() {
	for i := 0; i < 10; i++ {
		wg.Done()
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		wg.Done()
		fmt.Println("B:", i)
	}
}
