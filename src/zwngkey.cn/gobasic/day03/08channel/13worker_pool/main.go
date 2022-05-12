package main

import (
	"fmt"
	"time"
)

// worker pool  (goroutine 池)

// 在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，
// 控制goroutine的数量，防止goroutine泄漏和暴涨。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}

}

func main() {
	jobs := make(chan int, 100)

	results := make(chan int, 100)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}

	for i := 0; i < 5; i++ {
		<-results
	}

}
