package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")
	// timer := time.Tick(time.Second)
	timer := time.NewTicker(time.Second)

	for i := 10; i > 0; i-- {
		fmt.Println(i)
		// <-timer
		select {
		case <-timer.C:
		// case <-timer:
		// case <-time.After(1 * time.Second):
		case <-abort:
			fmt.Println("launch abort")
			return
		}
	}
	timer.Stop()
	fmt.Println("Lift off!")
}
