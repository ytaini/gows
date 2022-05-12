/*
 * @Author: zwngkey
 * @Date: 2022-05-12 17:33:36
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-12 17:57:15
 * @Description:
	Go语言new和make的使用区别和最佳实践
*/
package gobuiltin

/*
	new和make都可以用来分配内存，那他们有什么区别呢？在写代码过程中，对于new和make的最佳实践又是什么呢？

	从官方定义里可以看到，new有以下几个特点：
		1.分配内存。内存里存的值是对应类型的零值。
		2.只有一个参数。参数是分配的内存空间中所存储的值的类型，Go语言里的任何类型都可以是new的参数，比如int， 数组，结构体，甚至函数类型都可以。
		3.返回的是指针。

	注意：Go里的new和C++的new是不一样的：
		Go的new分配的内存可能在栈(stack)上，可能在堆(heap)上。C++ new分配的内存一定在堆上。
		Go的new分配的内存里的值是对应类型的零值，不能显示初始化指定要分配的值。C++ new分配内存的时候可以显示指定要存储的值。
		Go里没有构造函数，Go的new不会去调用构造函数。

	从官方定义里可以看到，make有如下几个特点：
		1.分配和初始化内存。
		2.只能用于slice, map和chan这3个类型，不能用于其它类型。
			如果是用于slice类型，make函数的第2个参数表示slice的长度，这个参数必须给值。
		3.返回的是原始类型，也就是slice, map和chan，不是返回指向slice, map和chan的指针。


	为什么针对slice, map和chan类型专门定义一个make函数？
		这是因为slice, map和chan的底层结构上要求在使用slice，map和chan的时候必须初始化，如果不初始化，
		  那slice，map和chan的值就是零值，也就是nil。我们知道：
			1.map如果是nil，是不能往map插入元素的，插入元素会引发panic
			2.chan如果是nil，往chan发送数据或者从chan接收数据都会阻塞
			3.slice会有点特殊，理论上slice如果是nil，也是没法用的。但是append函数处理了nil slice的情况，
				可以调用append函数对nil slice做扩容。但是我们使用slice，总是会希望可以自定义长度或者容量，这个时候就需要用到make。


	为什么slice是nil也可以直接append？
		对于nil slice，append会对slice的底层数组做扩容，通过调用mallocgc向Go的内存管理器申请内存空间，再赋值给原来的nil slice。


	最佳实践
		1.尽量不使用new
		2.对于slice, map和chan的定义和初始化，优先使用make函数
*/
