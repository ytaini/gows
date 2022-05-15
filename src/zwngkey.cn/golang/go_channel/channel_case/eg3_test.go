/*
 * @Author: zwngkey
 * @Date: 2022-05-14 23:01:13
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-15 03:53:03
 * @Description:
 */
package channelcase

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
	将通道用做互斥锁（mutex）
		容量为1的缓冲通道可以用做一次性二元信号量。
		  事实上，容量为1的缓冲通道也可以用做多次性二元信号量（即互斥锁）尽管这样的互斥锁效率不如sync标准库包中提供的互斥锁高效。

	有两种方式将一个容量为1的缓冲通道用做互斥锁：
		通过发送操作来加锁，通过接收操作来解锁；
		通过接收操作来加锁，通过发送操作来解锁。
*/
func Test9(t *testing.T) {
	lock := make(ChVoid, 1) // 容量必须为1

	var wg sync.WaitGroup

	var sum int

	wg.Add(100)

	// 发送操作来加锁
	// for i := 0; i < 100; i++ {
	// 	i := i
	// 	go func() {
	// 		defer wg.Done()
	// 		lock <- void{} //加锁
	// 		sum += i
	// 		<-lock //解锁
	// 	}()
	// }

	// 通过接收操作来加锁
	lock <- void{}
	for i := 0; i < 100; i++ {
		i := i
		go func() {
			defer wg.Done()
			<-lock //加锁
			sum += i
			lock <- void{} //解锁
		}()
	}
	wg.Wait()
	fmt.Println(sum)

}

/*
	将通道用做计数信号量（counting semaphore）
		缓冲通道可以被用做计数信号量。 计数信号量可以被视为多主锁。
			如果一个缓冲通道的容量为N，那么它可以被看作是一个在任何时刻最多可有N个主人的锁。
				上面提到的二元信号量是特殊的计数信号量，每个二元信号量在任一时刻最多只能有一个主人。

		计数信号量经常被使用于限制最大并发数。

		和将通道用做互斥锁一样，也有两种方式用来获取一个用做计数信号量的通道的一份所有权。
			通过发送操作来获取所有权，通过接收操作来释放所有权；
			通过接收操作来获取所有权，通过发送操作来释放所有权。
*/
type Seat int
type Bar chan Seat

// 下面是一个通过接收操作来获取所有权的例子：
func (bar Bar) ServeCustomer(c int) {
	log.Print("顾客#", c, "进入酒吧")
	// 只有获得一个座位的顾客才能开始饮酒。 所以在任一时刻同时在喝酒的顾客数不会超过座位数10。
	seat := <-bar // 需要一个位子来喝酒
	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(10+rand.Intn(6)))
	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
	bar <- seat // 释放座位，离开酒吧
}

func Test10(t *testing.T) {
	bar10 := make(Bar, 10) // 此酒吧有10个座位

	var wg sync.WaitGroup
	wg.Add(200)

	// 摆放10个座位。
	for seatId := 0; seatId < cap(bar10); seatId++ {
		bar10 <- Seat(seatId) // 均不会阻塞
	}

	for customerId := 0; customerId < 200; customerId++ {
		time.Sleep(1 * s)

		go func(customerId int) {
			defer wg.Done()
			bar10.ServeCustomer(customerId)
		}(customerId)
	}
	wg.Wait()
}

/*
	在上例中，尽管在任一时刻同时在喝酒的顾客数不会超过座位数10，但是在某一时刻可能有多于10个顾客进入了酒吧，
		因为某些顾客在排队等位子。 在上例中，每个顾客对应着一个协程。
			虽然协程的开销比系统线程小得多，但是如果协程的数量很多，则它们的总体开销还是不能忽略不计的。
			所以，最好当有空位的时候才创建顾客协程。
*/
func (bar Bar) ServeCustomerAtSeat(c int, seat Seat) {
	log.Print("顾客#", c, "进入酒吧")
	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(10+rand.Intn(6)))
	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
	bar <- seat // 释放座位，离开酒吧
}
func Test11(t *testing.T) {
	bar10 := make(Bar, 10) // 此酒吧有10个座位

	var wg sync.WaitGroup
	wg.Add(200)

	// 摆放10个座位。
	for seatId := 0; seatId < cap(bar10); seatId++ {
		bar10 <- Seat(seatId) // 均不会阻塞
	}

	for customerId := 0; customerId < 200; customerId++ {
		time.Sleep(1 * s)
		// 当有空位的时候才创建顾客协程
		seat := <-bar10
		go func(customerId int) {
			defer wg.Done()
			bar10.ServeCustomerAtSeat(customerId, seat)
		}(customerId)
	}
	wg.Wait()
}

/*
	在上面这个修改后的例子中，在任一时刻最多只有10个顾客协程在运行
		（但是在程序的生命期内，仍旧会有大量的顾客协程不断被创建和销毁）。
*/

/*
	在下面这个更加高效的实现中，在程序的生命期内最多只会有10个顾客协程被创建出来。
*/
func (bar Bar) ServeCustomerAtSeat2(consumers chan int) {
	for c := range consumers {
		seatId := <-bar
		log.Print("++ 顾客#", c, "在第", seatId, "个座位开始饮酒")
		time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
		log.Print("-- 顾客#", c, "离开了第", seatId, "个座位")
		bar <- seatId // 释放座位，离开酒吧
	}
}

func Test12(t *testing.T) {

	bar24x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	consumers := make(chan int)
	for i := 0; i < cap(bar24x7); i++ {
		go bar24x7.ServeCustomerAtSeat2(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}

/*
	如果我们并不关心顾客坐在哪个位置(有个位置就行)（这种情况在编程实践中很常见），则实际上bar24x7计数信号量是完全不需要的：
*/
func ServeCustomer(consumers chan int) {
	for c := range consumers {
		log.Print("++ 顾客#", c, "开始在酒吧饮酒")
		time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
		log.Print("-- 顾客#", c, "离开了酒吧")
	}
}

func Test13(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const BarSeatCount = 10
	consumers := make(chan int)
	for i := 0; i < BarSeatCount; i++ {
		go ServeCustomer(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}

// 通过发送操作来获取所有权的实现相对简单一些，省去了摆放座位的步骤。
type Customer struct{ id int }
type Bar1 chan Customer

func (bar Bar1) ServeCustomer1(c Customer) {
	log.Print("++ 顾客#", c.id, "开始饮酒")
	time.Sleep(time.Second * time.Duration(3+rand.Intn(16)))
	log.Print("-- 顾客#", c.id, "离开酒吧")
	<-bar // 离开酒吧，腾出位子
}

func Test14(t *testing.T) {
	var wg sync.WaitGroup

	const maxCustomerCount = 20

	wg.Add(maxCustomerCount)
	bar24x7 := make(Bar1, 10) // 最多同时服务10位顾客

	for customerId := 0; customerId < maxCustomerCount; customerId++ {
		time.Sleep(time.Second * 2)
		customer := Customer{customerId}
		bar24x7 <- customer // 等待进入酒吧
		go func() {
			defer wg.Done()
			bar24x7.ServeCustomer1(customer)
		}()
	}
	wg.Wait()
}

/*
	峰值限制（peak/burst limiting）
		将通道用做计数信号量用例和通道尝试（发送或者接收）操作结合起来可用实现峰值限制。 峰值限制的目的是防止过大的并发请求数。
*/
// 使得顾客不再等待而是离去或者寻找其它酒吧。
func Test17(t *testing.T) {
	var wg sync.WaitGroup

	const maxCustomerCount = 20

	bar24x7 := make(Bar1, 10) // 最多同时服务10位顾客

	for customerId := 0; customerId < maxCustomerCount; customerId++ {
		time.Sleep(time.Second * 1)
		customer := Customer{customerId}
		select {
		case bar24x7 <- customer: // 试图进入此酒吧
			wg.Add(1)
			go func() {
				defer wg.Done()
				bar24x7.ServeCustomer1(customer)
			}()
		default:
			log.Print("顾客#", customerId, "不愿等待而离去")
		}
	}
	wg.Wait()
}
