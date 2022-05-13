/*
 * @Author: zwngkey
 * @Date: 2022-05-13 20:36:26
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 21:39:36
 * @Description:
	通道用例大全

	通道同步技术比被很多其它语言采用的其它同步方案（比如角色模型和async/await模式）有着更多的应用场景和更多的使用变种。
*/
package channelcase

import (
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
