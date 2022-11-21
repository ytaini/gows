package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine())

}
