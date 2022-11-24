/*
 * @Author: zwngkey
 * @Date: 2022-05-13 21:38:35
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 19:17:30
 * @Description:
 */
package channelcase

import (
	"crypto/rand"
	"fmt"
	"log"
	"sort"
	"testing"
	"time"
)

/*
	使用通道实现通知
		通知可以被看作是特殊的请求/回应用例。在一个通知用例中，我们并不关心回应的值，我们只关心回应是否已发生。
			所以我们常常使用空结构体类型struct{}来做为通道的元素类型，因为空结构体类型的尺寸为零，能够节省一些内存（虽然常常很少量）。
*/

// 我们已知道，如果一个通道中无值可接收，则此通道上的下一个接收操作将阻塞到另一个协程发送一个值到此通道为止。
// 所以一个协程可以向此通道发送一个值来通知另一个等待着从此通道接收数据的协程。

// 在下面这个例子中，通道done被用来做为一个信号通道来实现单对单通知。
// 向一个通道发送一个值来实现单对单通知
func Test4(t *testing.T) {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		t.Error(err)
	}

	done := make(ChVoid) // 也可以是缓冲的

	// 排序协程
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 通知排序已完成
	}()

	// 并发地做一些其它事情...

	<-done
	t.Log("排序完成:")
	t.Log(values[0], values[len(values)-1])

}

// 从一个通道接收一个值来实现单对单通知
// 如果一个通道的数据缓冲队列已满（非缓冲的通道的数据缓冲队列总是满的）但它的发送协程队列为空，则向此通道发送一个值将阻塞，
// 直到另外一个协程从此通道接收一个值为止。 所以我们可以通过从一个通道接收数据来实现单对单通知。
// 一般我们使用非缓冲通道来实现这样的通知。
func Test5(t *testing.T) {
	done := make(chan struct{})
	// 此信号通道也可以缓冲为1。如果这样，则在下面
	// 这个协程创建之前，我们必须向其中写入一个值。

	go func() {
		fmt.Print("Hello")
		// 模拟一个工作负载。
		time.Sleep(s * 2)

		// 使用一个接收操作来通知主协程。
		<-done
	}()

	done <- struct{}{} // 阻塞在此，等待通知
	fmt.Println(" world!")

}

/*
	另一个事实是，上面的两种单对单通知方式其实并没有本质的区别。 它们都可以被概括为较快者等待较慢者发出通知。
*/

// 多对单和单对多通知
func worker(id int, ready chan struct{}, done chan<- struct{}) {
	<-ready // 阻塞在此，等待通知
	log.Print("Worker#", id, "开始工作")
	// 模拟一个工作负载。
	time.Sleep(s * time.Duration(id+1))
	log.Print("Worker#", id, "工作完成")
	done <- struct{}{} // 通知主协程（N-to-1）
}

func Test6(t *testing.T) {
	log.SetFlags(log.Ltime)
	ready, done := make(chan struct{}), make(chan struct{})
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// 模拟一个初始化过程
	time.Sleep(s * 3 / 2)

	//单对多通知.
	ready <- struct{}{}
	ready <- struct{}{}
	ready <- struct{}{}

	// 等待被多对单通知
	<-done
	<-done
	<-done
}

/*
	事实上，上例中展示的多对单和单对多通知实现方式在实践中用的并不多。
		在实践中，我们多使用sync.WaitGroup来实现多对单通知，
			使用关闭一个通道的方式来实现单对多通知  。
*/

// 通过关闭一个通道来实现群发通知
/*
	上一个用例中的单对多通知实现在实践中很少用，因为通过关闭一个通道的方式在来实现单对多通知的方式更简单。
		 我们已经知道，从一个已关闭的通道可以接收到无穷个值，我们可以利用这一特性来实现群发通知。

	当然，我们也可以通过关闭一个通道来实现单对单通知。事实上，关闭通道是实践中用得最多通知实现方式。

	从一个已关闭的通道可以接收到无穷个值这一特性也将被用在很多其它在后面将要介绍的用例中。
		实际上，这一特性被广泛地使用于标准库包中。比如，context标准库包使用了此特性来传达操作取消消息。
*/

func Test7(t *testing.T) {
	log.SetFlags(log.Ltime)
	ready, done := make(chan struct{}), make(chan struct{})
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// 模拟一个初始化过程
	time.Sleep(s * 3 / 2)

	//单对多通知.
	close(ready)
	// 等待被多对单通知
	<-done
	<-done
	<-done
}

// 定时通知（timer）

// 用通道实现一个一次性的定时通知器是很简单的。 下面是一个自定义实现：

func afterDuration(d time.Duration) <-chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		ch <- struct{}{}
	}()
	return ch
}

func Test8(t *testing.T) {
	log.SetFlags(log.Ltime)
	log.Println("start")
	<-afterDuration(3 * s)
	log.Println("end")
	<-afterDuration(3 * s)
	log.Println("xiix")
}

/*
	事实上，time标准库包中的After函数提供了和上例中AfterDuration同样的功能。
		在实践中，我们应该尽量使用time.After函数以使代码看上去更干净。

	注意，操作<-time.After(aDuration)将使当前协程进入阻塞状态，而一个time.Sleep(aDuration)函数调用不会如此。

	<-time.After(aDuration)经常被使用在后面将要介绍的超时机制实现中。
*/
