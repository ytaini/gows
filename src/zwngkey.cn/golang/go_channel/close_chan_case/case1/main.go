/*
  - @Author: wzmiiiiii
  - @Date: 2022-11-24 12:10:52

* @LastEditors: wzmiiiiii
* @LastEditTime: 2022-11-24 12:28:37
  - @Description:
    优雅地关闭通道的方法
    在各种情形下使用纯通道操作来关闭通道的方法。
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 情形一：M个接收者和一个发送者。发送者通过关闭用来传输数据的通道来传递发送结束信号
// 这是最简单的一种情形。当发送者欲结束发送，让它关闭用来传输数据的通道即可。
func main() {
	rand.Seed(time.Now().UnixNano())

	const Max = 10000
	const NumReceivers = 20

	wg := sync.WaitGroup{}
	wg.Add(NumReceivers)

	dataCh := make(chan int, 10)
	// 发送者
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				// 此唯一的发送者可以安全地关闭此数据通道。
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// 接收者
	for i := 0; i < NumReceivers; i++ {
		i := i
		go func() {
			defer wg.Done()
			for v := range dataCh {
				// 接收数据直到通道dataCh已关闭
				// 并且dataCh的缓冲队列已空。
				fmt.Printf("第%d个接收者,接受到值: %d\n", i+1, v)
			}
		}()
	}
	wg.Wait()
}
