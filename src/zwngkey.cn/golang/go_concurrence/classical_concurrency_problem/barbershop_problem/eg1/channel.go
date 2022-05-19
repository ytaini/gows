/*
 * @Author: zwngkey
 * @Date: 2022-05-18 22:39:02
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-18 22:58:17
 * @Description:
	使用Channel实现Semaphore
		我们可以使用Channel实现一个Semaphore来解决理发店问题。

		这里为什么不使用Go官方扩展的semaphore.Weighted并发原语呢，
			是因为semaphore.Weighted有个问题，在Accquire之前调用Release方法的话会panic.

	这种channel实现的Semaphore有什么缺陷吗？那就是如果队列的长度太大的话，channel的容量就会很大。
		不过如果类型设置为strcut{}类型的话，就会节省很多的内存，所以一般也不会有什么问题，
			虽然比官方扩展的计数器方式的semaphore.Weighted多占用一些空间，但是占用的空间还是有限的。
*/
package main

import "log"

// 有了这个并发原语，我们就容易解决理发店问题了。
type Semaphore chan struct{}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

// 注意这里我们实现了TryAcquire,就是为了顾客到来的是否检查有没有空闲的座位。
func (s Semaphore) TryAcquire() bool {
	select {
	case s <- struct{}{}: // 还有空位子
		return true
	default: // 没有空位子了,离开
		return false
	}
}
func (s Semaphore) Release() {
	<-s
}

// 我们定义了有三个等待座位的信号量。
// Tony老师先调用Release方法，也就是想从座位上请一位顾客过来理发，以便空出一个等待座位。
// 如果没有顾客，Tony就会无奈的等待和睡觉了。
var seats2 = make(Semaphore, 3)

func Barber2() {
	for {
		// 等待一个用户
		log.Println("Tony老师尝试请求一个顾客")
		seats2.Release()
		log.Println("Tony老师找到一位顾客，开始理发")
		RandomPause(2000)
	}
}

// 模拟顾客陆陆续续的过来
func Customers2() {
	for {
		RandomPause(1000)
		go Customer2()
	}
}

// 顾客
func Customer2() {
	if ok := seats2.TryAcquire(); ok {
		log.Println("一位顾客开始坐下排队理发")
	} else {
		log.Println("没有空闲座位了，一位顾客离开了")
	}
}
