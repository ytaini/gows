/*
 * @Author: zwngkey
 * @Date: 2022-05-13 20:36:26
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-15 05:11:15
 * @Description:
	通道用例大全

	通道同步技术比被很多其它语言采用的其它同步方案（比如角色模型和async/await模式）有着更多的应用场景和更多的使用变种。
*/
package channelcase

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
	将通道用做future/promise
		很多其它流行语言支持future/promise来实现异步（并发）编程。 Future/promise常常用在请求/回应场合。
*/

// 返回单向接收通道做为函数返回结果
func longTimeRequests() RecvChInt {
	ch1 := make(ChInt)
	go func() {
		time.Sleep(3 * s)
		ch1 <- rand.Intn(100)
	}()
	return ch1
}

func sumSquares(num1, num2 int) int {
	return num1*num1 + num2*num2
}

// 在下面这个例子中，sumSquares函数调用的两个实参的请求是并发进行的。 每个通道读取操作将阻塞到请求返回结果为止。
//  两个实参总共需要大约3秒钟（而不是6秒钟）准备完毕（以较慢的一个为准）。
func Test(t *testing.T) {

	ch1, ch2 := longTimeRequests(), longTimeRequests()
	//3s
	t.Log(sumSquares(<-ch1, <-ch2))
	//6s
	// fmt.Println(sumSquares(<-longTimeRequests(), <-longTimeRequests()))
}

// 将单向发送通道类型用做函数实参
func longTimeRequests2(ch SendChInt) {
	time.Sleep(3 * s)
	ch <- rand.Intn(100)
}

// 和上例一样，在下面这个例子中，sumSquares函数调用的两个实参的请求也是并发进行的。
// 和上例不同的是longTimeRequest2函数接收一个单向发送通道类型参数而不是返回一个单向接收通道结果。
func Test1(t *testing.T) {
	ch1, ch2 := make(ChInt), make(ChInt)

	go longTimeRequests2(ch1)
	go longTimeRequests2(ch2)

	t.Log(sumSquares(<-ch1, <-ch2))
}

// 对于上面这个特定的例子，我们可以只使用一个通道来接收回应结果，因为两个参数的作用是对等的。
// 这可以看作是数据聚合的一个应用。
func Test2(t *testing.T) {
	ch1 := make(ChInt, 2)

	go longTimeRequests2(ch1)
	go longTimeRequests2(ch1)

	t.Log(sumSquares(<-ch1, <-ch1))
}

/*
	采用最快回应
*/

// 本用例可以看作是上例中只使用一个通道变种的增强。
// 	 	有时候，一份数据可能同时从多个数据源获取。这些数据源将返回相同的数据。
//	  	因为各种因素，这些数据源的回应速度参差不一，甚至某个特定数据源的多次回应速度之间也可能相差很大。
//		同时从多个数据源获取一份相同的数据可以有效保障低延迟。我们只需采用最快的回应并舍弃其它较慢回应。

// 注意：如果有N个数据源，为了防止被舍弃的回应对应的协程永久阻塞，则传输数据用的通道必须为一个容量至少为N-1的缓冲通道
func source(ch SendChInt) {
	ra, rb := rand.Intn(1000), rand.Intn(3)+1

	time.Sleep(time.Duration(rb) * s)

	ch <- ra
}
func Test3(t *testing.T) {
	ch := make(ChInt, 5)

	stime := time.Now()

	for i := 0; i < cap(ch); i++ {
		go source(ch)
	}
	rnd := <-ch //只有最快的回应被使用了

	etime := time.Since(stime)
	t.Log(etime)
	t.Log(rnd)
}

/*
	做为函数参数和返回结果使用的通道可以是缓冲的，从而使得请求协程不需阻塞到它所发送的数据被接收为止。

	有时，一个请求可能并不保证返回一份有效的数据。对于这种情形，我们可以使用一个形如struct{v T; err error}的结构体类型
		或者一个空接口类型做为通道的元素类型以用来区分回应的值是否有效。

	有时，一个请求可能需要比预期更长的用时才能回应，甚至永远都得不到回应。 我们可以使用本文后面将要介绍的超时机制来应对这样的情况。

	有时，回应方可能会不断地返回一系列值，这也同时属于后面将要介绍的数据流的一个用例。
*/

/*
	另一种“采用最快回应”的实现方式
		我们也可以使用选择机制来实现“采用最快回应”用例。 每个数据源协程只需使用一个缓冲为1的通道并向其尝试发送回应数据即可。

		注意，使用选择机制来实现“采用最快回应”的代码中使用的通道的容量必须至少为1，以保证最快回应总能够发送成功。
			 否则，如果数据请求者因为种种原因未及时准备好接收，则所有回应者的尝试发送都将失败，从而所有回应的数据都将被错过。
*/
// 示例代码如下：

func source1(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// 休眠1秒/2秒/3秒
	time.Sleep(time.Duration(rb) * time.Second)
	select {
	case c <- ra:
	default:
	}
}
func Test18(t *testing.T) {
	c := make(chan int32, 1) // 此通道容量必须至少为1
	for i := 0; i < 5; i++ {
		go source1(c)
	}
	rnd := <-c // 只采用第一个成功发送的回应数据
	fmt.Println(rnd)
}

/*
	第三种“采用最快回应”的实现方式
		如果一个“采用最快回应”用例中的数据源的数量很少，比如两个或三个，
			我们可以让每个数据源使用一个单独的缓冲通道来回应数据，然后使用一个select代码块来同时接收这三个通道。

		注意：如果下例中使用的通道是非缓冲的，未被选中的case分支对应的两个source函数调用中开辟的协程将处于永久阻塞状态，从而造成内存泄露。
*/
// 示例代码如下：
func source2() <-chan int32 {
	c := make(chan int32, 1) // 必须为一个缓冲通道
	go func() {
		ra, rb := rand.Int31(), rand.Intn(3)+1
		time.Sleep(time.Duration(rb) * time.Second)
		c <- ra
	}()
	return c
}
func Test19(t *testing.T) {

	var rnd int32
	// 阻塞在此直到某个数据源率先回应。
	select {
	case rnd = <-source2():
	case rnd = <-source2():
	case rnd = <-source2():
	}
	fmt.Println(rnd)
}

/*
	超时机制（timeout）
		在一些请求/回应用例中，一个请求可能因为种种原因导致需要超出预期的时长才能得到回应，有时甚至永远得不到回应。
			对于这样的情形，我们可以使用一个超时方案给请求者返回一个错误信息。 使用选择机制可以很轻松地实现这样的一个超时方案。

		func requestWithTimeout(timeout time.Duration) (int, error) {
			c := make(chan int)
			go doRequest(c) // 可能需要超出预期的时长回应

			select {
			case data := <-c:
				return data, nil
			case <-time.After(timeout):
				return 0, errors.New("超时了！")
			}
		}
*/

/*
	脉搏器（ticker）
		我们可以使用尝试发送操作来实现一个每隔一定时间发送一个信号的脉搏器。

*/
func Tick(d time.Duration) <-chan void {
	ch := make(chan void, 1)
	go func() {
		for {
			time.Sleep(d)
			select {
			case ch <- void{}:
			default:
			}
		}
	}()
	return ch
}
func Test20(t *testing.T) {
	ti := time.Now()
	for range Tick(s) {
		fmt.Println(time.Since(ti))
	}
}

/*
	time标准库包中的Tick函数提供了同样的功能，但效率更高。 我们应该使用标准库包中的实现。
*/
func Test21(t *testing.T) {
	ti := time.Now()
	for range time.Tick(s) {
		fmt.Println(time.Since(ti))
	}
}

/*
	速率限制（rate limiting）
		同样地，我们也可以使用尝试机制来实现速率限制，但需要前面刚提到的定时器实现的配合。
			速率限制常用来限制吞吐和确保在一段时间内的资源使用不会超标。
*/
//  在此例中，任何一分钟时段内处理的请求数不会超过200。
type Request interface{}

func handle(r Request) {
	// fmt.Println(r.(int))
}

const RateLimitPeriod = time.Minute
const RateLimit = 200 // 任何一分钟内最多处理200个请求

func handleRequests(requests <-chan Request) {
	// quotas := make(chan time.Time, RateLimit)

	// go func() {
	// 	tick := time.NewTicker(RateLimitPeriod / RateLimit)
	// 	defer tick.Stop()
	// 	for t := range tick.C {
	// 		select {
	// 		case quotas <- t:
	// 		default:
	// 		}
	// 	}
	// }()

	for r := range requests {
		// <-quotas
		go handle(r)
	}
}

func Test22(t *testing.T) {
	requests := make(chan Request)
	go handleRequests(requests)
	// time.Sleep(time.Minute)
	for i := 0; ; i++ {
		requests <- i
	}
}
