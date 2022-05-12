package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		x := <-ch
		fmt.Printf("x: %v\n", x)
	}(ch)

	ch <- 10
	fmt.Println("发送")

	// time.Sleep(time.Second)

}
