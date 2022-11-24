/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-24 13:46:38
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 15:34:47
 * @Description:
 */
package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// 情形四：“M个接收者和一个发送者”情形的一个变种：用来传输数据的通道的关闭请求由第三方发出
// 	有时，数据通道（dataCh）的关闭请求需要由某个第三方协程发出。
// 	对于这种情形，我们可以使用一个额外的信号通道来通知唯一的发送者关闭数据通道（dataCh）。

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	notifyCh := make(chan struct{}) // 信号通道

	// 一些第三方协程
	for i := 0; i < NumThirdParties; i++ {
		go func() {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			select {
			case notifyCh <- struct{}{}:
			default:
			}
		}()
	}

	//发送者
	go func() {
		defer func() {
			close(dataCh)
		}()
		for {
			select {
			case <-notifyCh:
				return
			default:
			}
			select {
			case <-notifyCh:
				return
			case dataCh <- rand.Intn(Max):
			}
		}
	}()
	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
