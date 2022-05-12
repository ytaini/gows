/*
 * @Author: zwngkey
 * @Date: 2022-04-27 08:15:17
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-06 03:16:47
 * @Description:
	time.NewTicker与time.Tick函数的使用

*/
package gopkgtime

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestEg11(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)
	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go func(p int) {
			temp := rand.Intn(5)
			time.Sleep(time.Duration(temp) * time.Second)
			fmt.Println("I want to sleep", temp, "seconds")
			c <- p
		}(i)
	}

loop:
	for {
		select {
		case x := <-c:
			fmt.Println("goroutine-", x, "run done")
		case <-ticker.C:
			// case <-time.Tick(5 * time.Second):
			fmt.Println("time out")
			// os.Exit(2)
			break loop
		}
	}
}
