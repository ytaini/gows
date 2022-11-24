/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-21 19:39:27
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-23 19:03:15
 * @Description:
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

// worker pool  (goroutine 池)

// 在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，
// 控制goroutine的数量，防止goroutine泄漏和暴涨。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second / 10)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

var wg sync.WaitGroup

func main() {
	jobs := make(chan int, 100)

	results := make(chan int, 100)

	wg.Add(3)

	for i := 0; i < 3; i++ {
		i := i
		go func() {
			defer wg.Done()
			worker(i, jobs, results)
		}()
	}

	go func() {
		for i := 0; i < 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for v := range results {
		fmt.Println("v := ", v)
	}
	// for i := 0; i < 5; i++ {
	// 	<-results
	// }

}
