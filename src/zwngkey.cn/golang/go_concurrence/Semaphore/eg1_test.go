/*
 * @Author: zwngkey
 * @Date: 2022-05-15 01:27:15
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-15 02:47:36
 * @Description:
	go并发原语:啥是Semaphore?
*/
package goconc

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"testing"

	"golang.org/x/sync/semaphore"
)

/*
	信号量 是 并发编程中常见的同步机制.在标准库的并发原语中使用频繁,比如Mutex,WaitGroup等.

	维基百科上是这样解释信号量的：
		系统中，会给每一个进程一个信号量，代表每个进程当前的状态，未得到控制权的进程，
			会在特定的地方被迫停下来，等待可以继续进行的信号到来。

		通俗点解释就是，信号量通常使用一个整型变量 S 表示一组资源，当 G 完成对此信号量的等待（wait）时，S 就减 1，
			当 G 完成对此信号量的释放（release）时，S 就加 1。当计数值为 0 的时候，G 调用 wait 等待该信号量会阻塞，
			除非 S 又大于 0，等待的 G 才会解除阻塞，成功返回。

		举个例子，假如图书馆有 10 本《Go 语言编程之旅》，有 1 万个人都想读这本书，“僧多粥少”。
			所以，图书馆管理员先会让这 1 万个人进行登记，按照登记的顺序，借阅此书。
			如果书全部被借走，那么，其他想看此书的人就需要等待，如果有人还书了，图书馆管理员就会通知下一位同学来借阅这本书。
			这里的资源是《Go 语言编程之旅》这十本书，想读此书的同学就是 goroutine，图书管理员就是信号量。

	信号量是什么?
		信号量就是一种变量或者抽象数据类型,用于控制并发系统中多个进程对公共资源的访问,访问具有原子性.

	信号量主要分两类:
		1.计数信号量
		2.二元信号量.一种特殊的计数信号量,其值只有0或1,相当于互斥量,当值为1时资源可用,当值为0时,资源被锁住,进程阻塞无法继续执行.

	PV 操作
		信号量定义有两个操作 P 和 V，P 操作是减少信号量的计数值，而 V 操作是增加信号量的计数值。
		通常初始化时，将信号量 S 指定数值为 n，就像是一个有 n 个资源的池子。
		P 操作相当于请求资源，如果资源可用，就立即返回；如果没有资源或者不够，那么，G 会阻塞等待。
		V 操作会释放持有的资源，把资源返还给信号量。
		信号量的值除了初始化的操作以外，只能由 P/V 操作改变。
		我们一般用信号量保护一组资源，比如数据库连接池、几个打印机资源等等。
		如果信号量蜕变成二值信号量，那么，它的 P/V 就和互斥锁的 Lock/Unlock 一样了。


*/

/*
	信号量的实现-- 官方扩展包 Semaphore
		Go 在它的扩展包中提供了信号量 semaphore，不过这个信号量的类型名并不叫 Semaphore，而是叫 Weighted。这是一个带权重的信号量

		Weighted 的实现思路：使用互斥锁 + List 实现的。互斥锁实现其它字段的保护，而 List 实现了一个等待队列，
			等待者的通知是通过 Channel 的通知机制实现的。


		Weighted 主要包括两个结构体和几个常用方法。
			type Weighted struct {
				size    int64       // 最大资源个数，初始化的时候指定
				cur     int64       // 计数器，当前已使用资源数
				mu      sync.Mutex  // 互斥锁，对字段保护
				waiters list.List   // 等待者列表，当前处于阻塞等待的请求者 goroutine
			}

		其中 waiters 存储的数据是 waiter 对象，waiter 数据结构如下：
			type waiter struct {
				n     int64        		// 调用者申请的资源数
				ready chan<- struct{}   // 当调用者可以获取到信号量资源时, close chan，调用者便会收到通知，成功返回
			}

		方法
			1.阻塞获取资源的方法 -- Acquire()，源码如下：
				func (s *Weighted) Acquire(ctx context.Context, n int64) error {
					s.mu.Lock()

					// 有可用资源，直接返回
					if s.size-s.cur >= n && s.waiters.Len() == 0 {
						s.cur += n
						s.mu.Unlock()
						return nil
					}

					// 程序执行到这里说明无足够资源使用

					if n > s.size {
						s.mu.Unlock()
						<-ctx.Done()
						return ctx.Err()
					}

					// 资源不足，构造 waiter，将其加入到等待队列
					// ready channel 用于通知阻塞的调用者有资源可用，由释放资源的 goroutine 负责 close，起到消息通知的作用
					ready := make(chan struct{})
					w := waiter{n: n, ready: ready}
					elem := s.waiters.PushBack(w)      // 加入到等待队列
					s.mu.Unlock()

					// 调用者陷入 select 阻塞，除非收到外部 ctx 的取消信号或者被通知有资源可用
					select {
						case <-ctx.Done():     // 收到外面的控制信号
							err := ctx.Err()
							s.mu.Lock()
							select {
								case <-ready:    // 再次确认是否可能是被唤醒的，如果被唤醒了则忽略控制信号，返回 nil 表示成功
									err = nil
								default:      // 收到控制信息且还没有获取到资源，就直接将原来添加的 waiter 删除掉
									isFront := s.waiters.Front() == elem     // 当前 waiter 是否是链表头元素
									s.waiters.Remove(elem)     // 删除 waiter
									if isFront && s.size > s.cur {    // 如果是链表头元素且有资源可用则尝试唤醒链表第一个等待的 waiter
										s.notifyWaiters()
									}
							}
							s.mu.Unlock()
							return err
						case <-ready:      // 消息通知，请求资源的 goroutine 被释放资源的 goroutine 唤醒了
							return nil
					}
				}

			Acquire() 相当于 P 操作，可以一次获取多个资源，如果没有足够多的资源，调用者就会被阻塞。
				可以通过第一个参数 Context 增加超时或者 cancel 的机制。如果正常获取了资源，就返回 nil；
					否则，就返回 ctx.Err()，信号量不改变。


			2.非阻塞获取资源的方法 -- TryAcquire，源码如下：
				func (s *Weighted) TryAcquire(n int64) bool {
					s.mu.Lock()
					success := s.size-s.cur >= n && s.waiters.Len() == 0
					if success {
						s.cur += n
					}
					s.mu.Unlock()
					return success
				}

			这个方法比较简单，非阻塞地获取指定数量的资源，如果当前没有空闲资源，就直接返回 false。


			3.通知等待者 notifyWaiters，源码如下：
				func (s *Weighted) notifyWaiters() {
					for {
						next := s.waiters.Front()     // 获取队头元素
						if next == nil {        // 队列里没有元素
							break
						}

						w := next.Value.(waiter)
						if s.size-s.cur < w.n {       // 资源不满足请求者的要求
							break
						}

						s.cur += w.n              // 增加已用资源
						s.waiters.Remove(next)
						close(w.ready)         // 关闭 ready channel，用于通知调用者 goroutine 已经获取到资源，继续运行
					}
				}

			通过 for 循环从链表头部开始依次遍历链表中的所有 waiter，并更新计数器 weighted.cur，
			同时将其从链表中移除，直到遇到空闲资源小于 waiter.n 为止。
			仔细分析，我们会发现，notifyWaiters 方法是按照 FIFO 方式唤醒调用者。
			这样做的目的是为了避免调用者出现“饿死”的情况，当释放 10 个资源的时候，如果第一个等待者需要 11 个资源，
			那么，队列中的所有等待者都会继续等待，即使有的等待者只需要 1 个资源，
			否则的话，资源可能总是被那些请求资源数小的调用者获取，这样一来，请求资源数巨大的调用者，就没有机会获得资源了。

			4.释放占用的资源 -- Release()，源码如下：
				func (s *Weighted) Release(n int64) {
					s.mu.Lock()
					s.cur -= n       // 释放占用资源数
					if s.cur < 0 {
						s.mu.Unlock()
						panic("semaphore: released more than held")
					}
					s.notifyWaiters()   // 唤醒等待请求资源的 goroutine
					s.mu.Unlock()
				}

			Release() 相当于 V 操作，可以将 n 个资源释放，返还给信号量。
*/

/*
	怎么用？
		我们举个 worker pool 的例子，也是官网提供的：考拉兹猜想。

		“考拉兹猜想”说的是：对于每一个正整数，如果它是奇数，则对它乘 3 再加 1，如果它是偶数，则对它除以 2，
			如此循环，最终都能够得到 1。
*/
//我们的例子需要实现的是，对于给出的正整数，计算循环多少次之后能得到 1，代码如下：
func Test1(t *testing.T) {
	/*
		下面的代码创建数量与 CPU 核数相同的 worker，假设是 4， 相当于池子里只有 4 个资源可用，
			每个 worker 处理完一个整数，才能处理下一个，相当于控制住了并发数量。
	*/
	var (
		maxWorkers = runtime.GOMAXPROCS(0)                    // worker 数量
		sem        = semaphore.NewWeighted(int64(maxWorkers)) // 信号量
		out        = make([]int, 32)                          // 任务数
	)
	ctx := context.TODO()

	for i := range out {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(i int) {
			defer sem.Release(1)
			out[i] = collatzSteps(i + 1)
		}(i)
	}

	// 等待所有的任务执行完成，也可以通过 WaitGroup 实现
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}

	fmt.Println(out)

}
func collatzSteps(n int) (steps int) {
	if n <= 0 {
		panic("nonpositive input")
	}

	for ; n > 1; steps++ {
		if steps < 0 {
			panic("too many steps")
		}

		if n%2 == 0 {
			n /= 2
			continue
		}

		const maxInt = int(^uint(0) >> 1)
		if n > (maxInt-1)/3 {
			panic("overflow")
		}
		n = 3*n + 1
	}

	return steps
}

/*
	如何正确使用信号量？
		阅读完源码之后，会发现使用 semaphore 过程中一不小心就会导致错误，
			比如：如果请求的资源数比最大的资源数还大，那么，调用者可能永远被阻塞；
			调用 Release() 方法时，可以传递任意的整数。但如果传递一个比请求到的数量大的错误的数值，程序就会 panic；
			如果传递一个负数，会导致资源永久被持有，等等。

		使用时有哪些常犯的错误：
			请求的资源数大于最大的资源数；
			请求了资源，但是忘记释放；
			长时间持有资源，即使不需要它；
			释放了未请求过的资源；

		使用一项技术，保证不出错的前提是正确地使用它，对于信号量来说也是一样，所以使用信号量是应该格外小心，确保正确地传递参数，请求多少资源，就释放多少资源。
*/
