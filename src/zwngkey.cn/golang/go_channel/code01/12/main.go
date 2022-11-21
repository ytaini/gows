package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
// 1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
// 2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 3.主goroutine从resultChan取出结果并打印到终端输出

type Result struct {
	sourceData int64
	sumDate    int64
}

func A(jobChan chan<- int64) {
	for {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Int63()
		jobChan <- randNum
	}
}

func B(jobChan <-chan int64, resultChan chan<- *Result) {
	for i := 0; i < 24; i++ {
		go func() {
			for {
				var sum int64 = 0
				result := &Result{}
				n := <-jobChan
				result.sourceData = n

				for n > 0 {
					sum += n % 10
					n /= 10
				}
				result.sumDate = sum

				resultChan <- result
			}
		}()
	}
}

func main() {
	jobChan := make(chan int64)
	resultChan := make(chan *Result)

	go A(jobChan)

	go B(jobChan, resultChan)

	for v := range resultChan {
		fmt.Printf("源数据:%d,各位数之和:%d\n", v.sourceData, v.sumDate)
	}
}
