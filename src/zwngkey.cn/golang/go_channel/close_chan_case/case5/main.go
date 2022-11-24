/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-24 14:38:46
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 15:58:55
 * @Description:
 */

package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// 情形五：“N个发送者”的一个变种：用来传输数据的通道必须被关闭以通知各个接收者数据发送已经结束了
// 在上面的提到的“N个发送者”情形中，为了遵守通道关闭原则，我们避免了关闭数据通道（dataCh）。
// 但是有时候，数据通道（dataCh）必须被关闭以通知各个接收者数据发送已经结束。
// 对于这种“N个发送者”情形，我们可以使用一个中间通道将它们转化为“一个发送者”情形，
// 	然后继续使用上一节介绍的技巧来关闭此中间通道，从而避免了关闭原始的dataCh数据通道。

func main() {

	// rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 1000
	const NumReceivers = 10
	const NumSenders = 20
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int) // 将被关闭

	// 中间通道
	middleCh := make(chan int) // 不会被关闭

	// notifyCh是一个额外的信号通道。它的发送者为第三方协程和发送者。
	// 它的接收者为 中间层。
	notifyCh := make(chan string) // 用来通知中间层关闭dataCh,closed.

	// closed是一个额外的信号通道。它的发送者为中间层。
	// 它的接收者为 第三方协程和发送者。
	closed := make(chan struct{}) // 用来通知发送者dataCh已关闭.

	var stoppedBy string

	// 中间层
	go func() {
		exit := func(v int, needSend bool) {
			close(closed)
			if needSend {
				dataCh <- v
			}
			close(dataCh)
		}

		for {
			select {
			case stoppedBy = <-notifyCh:
				exit(0, false)
				return
			case v := <-middleCh:
				select {
				case stoppedBy = <-notifyCh:
					exit(v, true)
					return
				// 然后一个中间通道向dataCH发.
				case dataCh <- v:
				}
			}
		}
	}()

	stop := func(by string) {
		select {
		case notifyCh <- by:
			<-closed
		case <-closed:
		}
	}

	// 一些第三方协程
	for i := 0; i < NumThirdParties; i++ {
		go func(id string) {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop("3rd-party#" + id)
		}(strconv.Itoa(i))
	}

	// 发送者
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					stop("sender#" + id)
					return
				}

				select {
				case <-closed:
					return
				default:
				}

				select {
				case <-closed:
					return
				// 通过中间通道,N个发送者向中间通道发
				case middleCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

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
	log.Println("stopped by", stoppedBy)
}
