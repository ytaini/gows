/*
 * @Author: zwngkey
 * @Date: 2022-04-30 11:42:18
 * @LastEditTime: 2022-11-24 11:23:39
 * @Description:
 */

package gogoroutine

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
	协程（goroutine）
		并发计算是指若干计算可能在某些时间片段内同时运行.
		并行计算是指多个计算在任何时间点都在同时运行

		协程有时也被称为绿色线程。	绿色线程是由程序的运行时（runtime）维护的线程。
			一个绿色线程的内存开销和情景转换（context switching）时耗比一个系统线程常常小得多。
			 只要内存充足，一个程序可以轻松支持上万个并发协程。

		Go不支持创建系统线程，所以协程是一个Go程序内部唯一的并发实现方式。

		每个Go程序启动的时候只有一个对用户可见的协程，我们称之为主协程。
			一个协程可以开启更多其它新的协程。
			在Go中，开启一个新的协程是非常简单的。
			我们只需在一个函数调用之前使用一个go关键字，即可让此函数调用运行在一个新的协程之中。
			 当此函数调用退出后，这个新的协程也随之结束了。
			 我们可以称此函数调用为一个协程调用（或者为此协程的启动调用）。
			 一个协程调用的所有返回值（如果存在的话）必须被全部舍弃。

		当一个程序的主协程退出后，此程序也就退出了，即使还有一些其它协程在运行。
*/

func SayGreetings(greet string, count int) {
	for i := 0; i < count; i++ {
		d := time.Duration(rand.Intn(5)) / 2
		log.Println(greet)
		time.Sleep(time.Second * d)
	}
}

func TestSayGreetings(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.Ldate | log.Ltime)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		SayGreetings("hi", 10)
	}()
	go func() {
		defer wg.Done()
		go SayGreetings("go", 10)
	}()
	wg.Wait()
}

/*
并发同步:(concurrency synchronization）

	不同的并发计算可能共享一些资源，其中共享内存资源最为常见。

	在一个并发程序中，常常会发生下面的情形:
		在一个计算向一段内存写数据的时候，另一个计算从此内存段读数据，结果导致读出的数据的完整性得不到保证。
		在一个计算向一段内存写数据的时候，另一个计算也向此段内存写数据，结果导致被写入的数据的完整性得不到保证。

	这些情形被称为数据竞争（data race）。
	并发编程的一大任务就是要调度不同计算，控制它们对资源的访问时段，以使数据竞争的情况不会发生。
	此任务常称为并发同步（或者数据同步）。Go支持几种并发同步技术.

并发编程中的其它任务包括：

	决定需要开启多少计算；
	决定何时开启、阻塞、解除阻塞和结束哪些计算；
	决定如何在不同的计算中分担工作负载。

Go支持几种并发同步技术。 其中， 通道是最独特和最常用的。
但是，为了简单起见，这里我们将使用sync标准库包中的WaitGroup来同步上面这个程序中的主协程和两个新创建的协程。

	WaitGroup类型有三个方法：Add、Done和Wait。
		Add方法用来注册新的需要完成的任务数。
		Done方法用来通知某个任务已经完成了。
		一个Wait方法调用将阻塞（等待）到所有任务都已经完成之后才继续执行其后的语句。

	使用time.Sleep调用来做并发同步不是一个好的方法。
*/
var wg sync.WaitGroup

func SayGreetings2(greet string, count int) {
	defer wg.Done() // 通知当前任务已经完成。
	for i := 0; i < count; i++ {
		d := time.Duration(rand.Intn(5)) / 2
		log.Println(greet)
		time.Sleep(time.Second * d)
	}
}

func TestSayGreetings2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.Ldate | log.Ltime)
	wg.Add(2) // 注册两个新任务。
	go SayGreetings2("hi", 5)
	go SayGreetings2("go", 5)
	wg.Wait() // 阻塞在这里，直到所有任务都已完成。
}

/*
	协程的状态
		从上面这个的例子，我们可以看到一个活动中的协程可以处于两个状态：
			运行状态和阻塞状态。一个协程可以在这两个状态之间切换。
			 比如上例中的主协程在调用wg.Wait方法的时候，将从运行状态切换到阻塞状态；
			  当两个新协程完成各自的任务后，主协程将从阻塞状态切换回运行状态。

		注意，一个处于睡眠中的（通过调用time.Sleep）或者在等待系统调用返回的协程被认为是处于运行状态，而不是阻塞状态。

		当一个新协程被创建的时候，它将自动进入运行状态，一个协程只能从运行状态而不能从阻塞状态退出。
		 如果因为某种原因而导致某个协程一直处于阻塞状态，则此协程将永远不会退出。
		  除了极个别的应用场景，在编程时我们应该尽量避免出现这样的情形。

		一个处于阻塞状态的协程不会自发结束阻塞状态，它必须被另外一个协程通过某种并发同步方法来被动地结束阻塞状态。
		 如果一个运行中的程序 当前所有的协程 都出于阻塞状态，则这些协程将永远阻塞下去，程序将被视为死锁了。
		  当一个程序死锁后，官方标准编译器的处理是让这个程序崩溃。
*/
