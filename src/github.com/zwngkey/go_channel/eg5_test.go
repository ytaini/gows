/*
 * @Author: zwngkey
 * @Date: 2022-05-11 16:06:40
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-11 16:09:18
 * @Description:go select
 */
package gochannel

import (
	"fmt"
	"testing"
	"time"
)

func TestEg51(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)
		case y := <-ch2:
			fmt.Println(y)
		}
	}
}
