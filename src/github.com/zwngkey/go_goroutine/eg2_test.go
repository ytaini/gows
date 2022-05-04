/*
 * @Author: zwngkey
 * @Date: 2022-04-30 13:32:57
 * @LastEditTime: 2022-04-30 21:28:52
 * @Description: Go 并发：一些有趣的现象和要避开的 “坑”
 */

package godefergopanicrecover

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
	预想: count : 10000
	结果: count 会 <=10000
	原因:
		count++ 并不是一个原子操作
		可能会存在多个 goroutine 同时读取到 count=x 的情况,并各自自增 1，再将其写回。
		与此同时也会有其他的 goroutine 可能也在其自增时读到了值，形成了互相覆盖的情况，这是一种并发访问共享数据的错误。
	如何避免:
		这类竞争问题可以通过 Go 语言所提供的的 race 检测（Go race detector）来进行分析和发现：
			编译器会通过探测所有的内存访问，监听其内存地址的访问（读或写）。
			在应用运行时就能够发现对共享变量的访问和操作，进而发现问题并打印出相关的警告信息。
*/
var wg = sync.WaitGroup{}

func Test1(t *testing.T) {
	var count int32
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				// count++
				// 解决:
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

/*
	特点:
		即使所有的 goroutine 都创建完了，但 goroutine 不一定已经开始运行了。
			其整个程序扭转实质上分为了多个阶段，也就是各自运行的时间线并不同，可以其拆分为：
				先创建：for-loop 循环创建 goroutine。
				再调度：协程goroutine 开始调度执行。
				才执行：开始执行 goroutine 内的输出。

		同时 goroutine 的调度存在一定的随机性（建议了解一下 GMP 模型），那么其输出的结果就势必是无序且不稳定的。
*/
func Test2(t *testing.T) {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

/*
	通道实现并发同步
*/
func Test21(t *testing.T) {
	ch := make(chan struct{}, 5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < 5; i++ {
		<-ch
	}
}

/**
 * @description:
 * @param index (int)
 * @param ch (chan<-int) 只发送通道
 * @param wg (*sync.WaitGroup)  不能为sync.WaitGroup值类型.
 * @return ()
 */
func doSomething(index int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start index: ", index)
	time.Sleep(20 * time.Millisecond)
	ch <- index
}

func Test3(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 100)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go doSomething(i, ch, &wg)
	}
	wg.Wait()
	close(ch)
	fmt.Println("all done")
	for v := range ch {
		fmt.Println("from ch:", v)
	}
}
