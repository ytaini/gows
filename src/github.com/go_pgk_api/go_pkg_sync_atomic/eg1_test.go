/*
 * @Author: zwngkey
 * @Date: 2022-05-11 21:11:13
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-11 22:25:28
 * @Description:
 */
/*
	Go语言标准库中的sync/atomic包提供了偏底层的原子内存原语(atomic memory primitives)，用于实现同步算法，
		其本质是将底层CPU提供的原子操作指令封装成了Go函数。

	使用sync/atomic提供的原子操作可以确保在任意时刻只有一个goroutine对变量进行操作，避免并发冲突。

	使用sync/atomic需要特别小心，Go官方建议只有在一些偏底层的应用场景里才去使用sync/atomic，
		其它场景建议使用channel或者sync包里的锁

	sync/atomic提供了5种类型的原子操作和1个Value类型。

	5种类型的原子操作
		swap操作：SwapXXX
		compare-and-swap操作：CompareAndSwapXXX
		add操作：AddXXX
		load操作：LoadXXX
		store操作：StoreXXX
	这几种类型的原子操作只支持几个基本的数据类型。

	Value类型
		由于上面5种类型的原子操作只支持几种基本的数据类型，因此为了扩大原子操作的使用范围，
			Go团队在1.4版本的sync/atomic包中引入了一个新的类型Value。Value类型可以用来读取(Load)和修改(Store)任意类型的值。

		Go 1.4版本的Value类型只有Load和Store2个方法，Go 1.17版本又给Value类型新增了CompareAndSwap和Swap这2个新方法。
*/

/*
	swap操作
		func SwapInt32(addr *int32, new int32) (old int32)
			swap操作实现的功能是 把addr 指针指向的内存里的值替换为新值new，然后返回旧值old，是如下伪代码的原子实现：
				old = *addr
				*addr = new
				return old


	compare-and-swap操作
		func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
			compare-and-swap操作实现的功能是先比较addr 指针指向的内存里的值是否为旧值old相等。
				1.如果相等，就把addr指针指向的内存里的值替换为新值new，并返回true，表示操作成功。
				2.如果不相等，直接返回false，表示操作失败。

		伪代码的原子实现：
		if *addr == old {
			*addr = new
			return true
		}
		return false


	add操作
		func AddInt32(addr *int32, delta int32) (new int32)
		  add操作实现的功能是把addr 指针指向的内存里的值和delta做加法，然后返回新值，是如下伪代码的原子实现：
			*addr += delta
			return *addr


	load操作
		func LoadInt32(addr *int32) (val int32)
		  load操作实现的功能是返回addr 指针指向的内存里的值，是如下伪代码的原子实现：
			return *addr

	store操作
		store操作实现的功能是把addr 指针指向的内存里的值修改为val，是如下伪代码的原子实现：
			*addr = val
*/
/*
	Value类型
		Go标准库里的sync/atomic包提供了Value类型，可以用来并发读取和修改任何类型的值。


*/
/*
	CAS操作会有ABA问题

*/
package gopkgsyncatomic

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestEg1(t *testing.T) {
	// config变量用来存放该服务的配置信息
	var config atomic.Value

	// 初始化时从别的地方加载配置文件，并存到config变量里
	config.Store(loadConfig())

	go func() {
		// 每10秒钟定时拉取最新的配置信息，并且更新到config变量里
		for {
			time.Sleep(10 * time.Second)
			// 对应于赋值操作 config = loadConfig()
			config.Store(loadConfig())
		}
	}()

	// 创建协程，每个工作协程都会根据它所读取到的最新的配置信息来处理请求
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				// 对应于取值操作 c := config
				// 由于Load()返回的是一个interface{}类型，所以我们要先强制转换一下
				c := config.Load().(map[string]string)
				// 这里是根据配置信息处理请求的逻辑...
				_, _ = r, c
			}
		}()
	}
}

type configContext = map[string]string

func loadConfig() configContext {
	// 从数据库或者文件系统中读取配置信息，然后以map的形式存放在内存里
	return make(configContext)
}
func requests() chan int {
	// 将从外界中接收到的请求放入到channel里
	return make(chan int)
}
