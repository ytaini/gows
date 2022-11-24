/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-22 22:56:18
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-22 23:15:19
 * @Description:
 */
package main

import (
	"fmt"
)

func f1(a chan int) {
	for i := 0; i < 100; i++ {
		a <- i
	}
	close(a)
}

func f2(a, b chan int) {
	for v := range a {
		b <- v * v
	}
	close(b)
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)

	go f1(a)
	go f2(a, b)

	for v := range b {
		fmt.Println(v)
	}
}
