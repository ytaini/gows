/*
 * @Author: zwngkey
 * @Date: 2022-05-13 21:38:35
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-14 02:36:04
 * @Description:
 */
package channelcase

import (
	"crypto/rand"
	"fmt"
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
		done <- void{} // 通知排序已完成
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
	done := make(chan void)
	// 此信号通道也可以缓冲为1。如果这样，则在下面
	// 这个协程创建之前，我们必须向其中写入一个值。

	go func() {
		fmt.Print("Hello")
		// 模拟一个工作负载。
		time.Sleep(time.Second * 2)

		// 使用一个接收操作来通知主协程。
		<-done
	}()

	done <- void{} // 阻塞在此，等待通知
	fmt.Println(" world!")

}

/*
	另一个事实是，上面的两种单对单通知方式其实并没有本质的区别。 它们都可以被概括为较快者等待较慢者发出通知。
*/

// 多对单和单对多通知
