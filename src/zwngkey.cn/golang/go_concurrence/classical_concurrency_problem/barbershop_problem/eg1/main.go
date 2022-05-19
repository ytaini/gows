/*
 * @Author: zwngkey
 * @Date: 2022-05-18 20:38:59
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-18 22:50:56
 * @Description:
	经典并发问题： 理发店的故事
		Sleeping barber problem是一个经典的goroutine交互和并发控制的问题，可以很好的用来演示多写多读的并发问题(multiple writers multiple readers)。

		这个问题是这样子的。
		有这么一个理发店，有一位理发师和几个让顾客等待的座位：
			1.如果没有顾客，这位理发师就躺在理发椅上睡觉
			2.顾客必须唤醒理发师让他开始理发
			3.如果一位顾客到来，理发师正在理发
				如果还有顾客用来等待的座位，则此顾客坐下
				如果座位都满了，则此顾客离开
			4.理发师理完头发后，需要检查是否还有等待的顾客
				如果有，则请一位顾客起来开始理发
				如果没有，理发师则去睡觉
*/
package main

func main() {
	// go barber()
	// customers()
	go Barber2()
	Customers2()
}
