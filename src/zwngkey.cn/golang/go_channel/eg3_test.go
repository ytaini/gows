/*
 * @Author: zwngkey
 * @Date: 2022-05-06 22:19:46
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-06 22:41:38
 * @Description: select 的使用
 */
package gochannel

import (
	"fmt"
	"testing"
)

/*
	select 是 GO 语言中用来提供 IO 复用的机制，它可以检测多个 chan 是否 ready（可读/可写）。
*/
/*
	select 中的 case 执行顺序是随机的，如果某个 case 中的 channel 已经 ready，
	  那么就会执行相应的语句并退出 select 流程，如果所有 case 中的 channel 都未 ready，
	  	那么就会执行 default 中的语句然后退出 select 流程。

	由于启动的协程和 select 语句并不能保证执行的顺序，所以也有可能 select 执行时协程还未向channel中写入数据，
		所以 select 直接执行 default 语句并退出。因此，此程序理论上是有可能产生三种输出的.

*/
func TestEg31(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan2 <- 1
	}()

	go func() {
		chan1 <- 1
	}()

	select {
	case <-chan1:
		fmt.Println("chan1")
	case <-chan2:
		fmt.Println("chan2")
	default:
		fmt.Println("default")
	}
	fmt.Println("main exit")
}

/*
select 会随机检测各 case 语句中 channel 是否 ready，注意已关闭的 channel 也是可读的，

	所以下面程序中select 不会阻塞，具体执行哪个 case 语句具是随机的。
*/
func TestEg32(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()
	go func() {
		close(chan2)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1")
	case <-chan2:
		fmt.Println("chan2")
	}
	fmt.Println("main exit.")

}

// 当一个协程中出现空的 select 语句时,此协程会被永久阻塞.但 Go 自带死锁检测机制，
// 当发现此协程再也没有机会被唤醒时，则会发生 panic.
func TestEg33(t *testing.T) {
	select {}
}
