/*
* @Author: zwngkey
* @Date: 2022-05-15 03:01:17
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 19:17:45
* @Description:
*/
package channelcase

import (
	"fmt"
	"os"
	"testing"
	"time"
)

/*
	使用通道传送传输通道
*/
// 一个通道类型的元素类型可以是另一个通道类型。
// 在下面这个例子中， 单向发送通道类型chan<- int是另一个通道类型chan chan<- int的元素类型。
var counter = func(n int) chan<- chan<- int {
	requests := make(chan chan<- int)
	go func() {
		for request := range requests {
			if request == nil {
				n++ // 递增计数
			} else {
				request <- n // 返回当前计数
			}
		}
	}()
	return requests // 隐式转换到类型chan<- (chan<- int)
}(0)

func Test16(t *testing.T) {

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			counter <- nil
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done

	request := make(chan int, 1)
	counter <- request
	fmt.Println(<-request) // 2000
}

/*
	尝试发送和尝试接收
		含有一个default分支和一个case分支的select代码块可以被用做一个尝试发送或者尝试接收操作，
			取决于case关键字后跟随的是一个发送操作还是一个接收操作。

		尝试发送和尝试接收代码块永不阻塞。
*/
/*
	无阻塞地检查一个通道是否已经关闭
		假设我们可以保证没有任何协程会向一个通道发送数据，则我们可以使用下面的代码来（并发安全地）检查此通道是否已经关闭，
			此检查不会阻塞当前协程。

		此方法常用来查看某个期待中的通知是否已经来临。此通知将由另一个协程通过关闭一个通道来发送。
*/
func IsClosed(c chan struct{}) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}

/*
	对话（或称乒乓）
		两个协程可以通过一个通道进行对话，整个过程宛如打乒乓球一样。
			下面是一个这样的例子，它将打印出一系列斐波那契（Fibonacci）数。
*/

type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		ball := <-table // 接球
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue { // 溢出结束
			os.Exit(0)
		}
		lastValue = ball
		table <- ball // 回球
		time.Sleep(s / 2)
	}
}

func Test15(t *testing.T) {
	table := make(chan Ball)
	go func() {
		table <- 1 // （裁判）发球
	}()
	go Play("A:", table)
	Play("B:", table)
}
