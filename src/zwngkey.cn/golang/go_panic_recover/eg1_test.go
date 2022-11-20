/*
 * @Author: zwngkey
 * @Date: 2022-05-11 20:03:17
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-20 15:16:52
 * @Description:
 */
package gopanicrecover

import (
	"fmt"
	"testing"
)

/*
	恐慌（panic）和恢复（recover）
		Go不支持异常抛出和捕获，而是推荐使用返回值显式返回错误。
			不过，Go支持一套和异常抛出/捕获类似的机制。此机制称为恐慌/恢复（panic/recover）机制。

		可以调用内置函数panic来产生一个恐慌以使当前协程进入恐慌状况

		进入恐慌状况是另一种使当前函数调用开始返回的途径。 一旦一个函数调用产生一个恐慌，
			此函数调用将立即进入它的退出阶段，在此函数调用中被推入堆栈的延迟调用将按照它们被推入的顺序逆序执行。

		通过在一个延迟函数调用之中调用内置函数recover，当前协程中的一个恐慌可以被消除，从而使得当前协程重新进入正常状况。

		在一个处于恐慌状况的协程退出之前，其中的恐慌不会蔓延到其它协程。 如果一个协程在恐慌状况下退出，它将使整个程序崩溃。

		内置函数panic和recover的声明原型如下：
			func panic(v interface{})
			func recover() interface{}

		recover函数的返回值为其所恢复的恐慌的参数值.

		一般说来，恐慌用来表示正常情况下不应该发生的逻辑错误。
			如果这样的一个错误在运行时刻发生了，则它肯定是由于某个bug引起的。
			另一方面，非逻辑错误是现实中难以避免的错误，它们不应该导致恐慌。 我们必须正确地对待和处理非逻辑错误。

		一些致命性错误不属于恐慌
			对于官方标准编译器来说，很多致命性错误（比如堆栈溢出和内存不足）不能被恢复。它们一旦产生，程序将崩溃。

*/

func TestEg1(t *testing.T) {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
