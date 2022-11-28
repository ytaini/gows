/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-27 12:16:58
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-27 15:09:16
 * @Description:
 */

package code

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 超时控制

// 通常健壮的程序都是要设置超时时间的，避免因为服务端长时间响应消耗资源，
// 所以一些web框架或rpc框架都会采用withTimeout或者withDeadline来做超时控制，
// 当一次请求到达我们设置的超时时间，就会及时取消，不在往下执行。
// withTimeout和withDeadline作用是一样的，就是传递的时间参数不同而已，
// 他们都会通过传入的时间来自动取消Context，这里要注意的是他们都会返回一个cancelFunc方法，
// 通过调用这个方法可以达到提前进行取消，不过在使用的过程还是建议在自动取消后也调用cancelFunc去停止定时减少不必要的资源浪费。

// withTimeout、WithDeadline不同在于WithTimeout将持续时间作为参数输入而不是时间对象，
// 这两个方法使用哪个都是一样的，看业务场景和个人习惯了，因为本质withTimout内部也是调用的WithDeadline。
func Test(t *testing.T) {
	HttpHandler()
}

// 达到超时时间终止接下来的执行
func HttpHandler() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel()
	deal(ctx)
}

func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

func deal(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
		}
	}
}

// 没有达到超时时间终止接下来的执行
func Test1(t *testing.T) {
	HttpHandler1()
}

func HttpHandler1() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel()
	done := make(chan struct{})
	go func() {
		deal1(ctx, cancel)
		done <- struct{}{}
	}()
	<-done
}

func NewContextWithTimeout1() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

func deal1(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
			cancel()
		}
	}
}

// 这里大家要记的一个坑，就是我们往从请求入口透传的调用链路中的context是携带超时时间的，
// 如果我们想在其中单独开一个goroutine去处理其他的事情并且不会随着请求结束后而被取消的话，
// 那么传递的context要基于context.Background或者context.TODO重新衍生一个传递，否决就会和预期不符合了
