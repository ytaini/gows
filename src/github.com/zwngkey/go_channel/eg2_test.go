/*
 * @Author: zwngkey
 * @Date: 2022-05-06 21:16:44
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-06 22:56:51
 * @Description: channel and goroutine的使用案例
 */
package gochannel

import (
	"fmt"
	"testing"
	"time"
)

func FibonacciNumber(ch chan<- uint64, n uint64) {
	var x, y uint64 = 0, 1
	for ; y < n; ch <- x {
		x, y = y, x+y
	}
	close(ch)
}

func ComputedFibonacciNumber(ch chan uint64, n uint64) <-chan uint64 {
	go FibonacciNumber(ch, n)
	return ch
}

func TestEg21(t *testing.T) {
	ch := make(chan uint64)
	var n uint64 = 2 << 20
	go FibonacciNumber(ch, n)

	for v := range ch {
		time.Sleep(time.Second / 2)
		fmt.Println(v)
	}
}

func TestEg22(t *testing.T) {
	ch := make(chan uint64)
	var n uint64 = 2 << 30
	for v := range ComputedFibonacciNumber(ch, n) {
		time.Sleep(time.Second / 4)
		fmt.Println(v)
	}
}
