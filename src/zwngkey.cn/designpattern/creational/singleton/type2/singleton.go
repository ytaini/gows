/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 23:04:53
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-06 23:46:27
 */
package type2

import "sync"

// 优点: Lazy Loading
// 缺点: 存在并发安全问题.只能单线程使用

// 懒汉式(通过包级全局变量+私有类型+一个公有函数)
type singleton1 struct{}

var instance1 *singleton1

func GetInstance1() *singleton1 {
	if instance1 == nil { //多个协程可能同时进入这个if语句块
		instance1 = new(singleton1)
	}
	return instance1
}

// ----------------------------------------------------------------------------------------

// 懒汉式(并发安全, 激进的加锁)
type singleton2 struct{}

var instance2 *singleton2

var lock sync.Mutex

// 这样加锁带来其他潜在的严重问题，把对该函数的并发调用变成了串行。
func GetInstance2() *singleton2 {
	lock.Lock()
	defer lock.Unlock()
	if instance2 == nil {
		instance2 = new(singleton2)

	}
	return instance2
}

// ----------------------------------------------------------------------------------------

// 懒汉式 (并发安全)
type singleton3 struct{}

var instance3 *singleton3

var lock1 sync.Mutex

// Check-Lock-Check模式
// 该模式背后的思想是，你应该首先进行检查，以最小化任何主动锁定，因为IF语句的开销要比加锁小。
// 其次，我们希望等待并获取互斥锁，这样在同一时刻在那个块中只有一个执行。
// 但是，在第一次检查和获取互斥锁之间，可能有其他goroutine获取了锁，
// 因此，我们需要在锁的内部再次进行检查，以避免用另一个实例覆盖了实例。
func GetInstance3() *singleton3 {
	if instance3 == nil {
		lock1.Lock()
		defer lock1.Unlock()
		if instance3 == nil {
			instance3 = new(singleton3)

		}
	}
	return instance3
}
