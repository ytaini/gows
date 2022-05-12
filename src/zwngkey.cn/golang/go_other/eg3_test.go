/*
 * @Author: zwngkey
 * @Date: 2022-05-12 18:00:00
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-12 18:38:40
 * @Description:
 */
package goother

import (
	"fmt"
	"testing"
)

/*
	在Go中如果使用过map和channel，就会发现把map和channel作为函数参数传递，
		不需要在函数形参里对map和channel加指针标记*就可以在函数体内改变外部map和channel的值。
*/

func changeMap(data map[string]any) {
	data["c"] = 3
}

// 下面的例子里，函数changeMap改变了外部的map类型counter的值。
func Test31(t *testing.T) {
	counter := map[string]any{"a": 1, "b": 2}
	fmt.Println("begin:", counter)
	changeMap(counter)
	fmt.Println("after:", counter)
}

/*
	什么是引用变量(reference variable)和引用传递(pass-by-reference)
		引用变量和引用传递的特点如下：
			引用变量和原变量的内存地址一样。就像上面的例子里引用变量b和原变量a的内存地址相同。
			函数使用引用传递，可以改变外部实参的值。就像上面的例子里，changeValue函数使用了引用传递，改变了外部实参a的值。
			对原变量的值的修改也会改变引用变量的值。就像上面的例子里，changeValue函数对a的修改，也改变了引用变量b的值


	Go有引用变量(reference variable)和引用传递(pass-by-reference)么？
		先给出结论：Go语言里没有引用变量和引用传递。

		在Go语言里，不可能有2个变量有相同的内存地址，也就不存在引用变量了。

		注意：这里说的是不可能2个变量有相同的内存地址，但是2个变量指向同一个内存地址是可以的，这2个是不一样的。
*/
/*
	可以看出，变量p1和p2的值相同，都指向变量a的内存地址。但是变量p1和p2自己本身的内存地址是不一样的。
		而C++里的引用变量和原变量的内存地址是相同的。

	因此，在Go语言里是不存在引用变量的，也就自然没有引用传递了。
*/
func Test32(t *testing.T) {
	a := 10
	var p1 *int = &a
	var p2 *int = &a
	fmt.Printf("a address: %p\n", &a)
	fmt.Println("p1 value:", p1, " address:", &p1)
	fmt.Println("p2 value:", p2, " address:", &p2)
}

/*
	从下面的代码中可以看出，函数initMap并不会改变外部实参data的值，这就证明了map并不是引用变量。

	那问题来了，为啥map作为函数参数不是使用的引用传递，但是在本文最开头举的例子里，却可以改变外部实参的值呢？
*/
func initMap(data map[string]int) {
	data = make(map[string]int)
	fmt.Println("in function initMap, data == nil:", data == nil) //false
}

func Test33(t *testing.T) {
	var data map[string]int
	fmt.Println("before init, data == nil:", data == nil) //true
	initMap(data)
	fmt.Println("after init, data == nil:", data == nil) //true
}

/*
	map究竟是什么？
		结论是：map变量是指向runtime.hmap的指针

		当我们使用下面的代码初始化map的时候
			data := make(map[string]int)
		Go编译器会把make调用转成对runtime.makemap的调用,从源代码中可以看出，runtime.makemap返回的是一个指向runtime.hmap结构的指针。

		如果map是指针，那make返回的不应该是*map[string]int么，为啥官方文档里说的是not a pointer to it.

		这里其实也有Go语言历史上的一个演变过程，
			在Go语言早期，的确对于map是使用过指针形式的，但是最后Go设计者们发现，几乎没有人使用map不加指针，
				因此就直接去掉了形式上的指针符号*。
*/
/*
	通过下面的例子，来验证map变量到底是不是指针。
*/

func Test34(t *testing.T) {
	m1 := make(map[string]int)
	fmt.Printf("%p\n", m1)

	var s = make([]int, 2)
	fmt.Printf("%p\n", s)

	var ch = make(chan int)
	fmt.Printf("%p\n", ch)
}
