/*
 * @Author: zwngkey
 * @Date: 2022-04-21 10:25:49
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-12 17:35:01
 * @Description:	go中的内置函数
 */
package gobuiltin

import (
	"fmt"
	"testing"
)

func TestEg21(t *testing.T) {

	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Println(len(s), cap(s))
	// 作为一种特殊的情况，将字符追加到字节数组之后是合法的，
	a := append([]byte{'h', 'e', 'l', 'l', 'o'}, "world"...)
	b := append([]byte("hello"), "world"...)
	c := make([]byte, 2)
	// 特殊情况是，它也能将字节从字符串复制到字节切片中
	copy(c, "hello")
	fmt.Printf("a: %q\n", a)
	fmt.Printf("b: %q\n", b)
	fmt.Printf("c: %s\n", c)
	fmt.Printf("c: %q\n", c)

	// len 内建函数返回 v 的长度，这取决于具体类型：
	//	数组：v 中元素的数量。
	//	数组指针：*v 中元素的数量（即使 v 为 nil）。
	//	切片或映射：v 中元素的数量；若 v 为 nil，len(v) 即为零。
	//	字符串：v 中字节byte的数量。而不是像其他语言中计算好的unicode字符串中字符的数量。
	// 要在Go中得到相同的结果，可以使用“unicode/utf8”包中的 RuneCountInString()函数。
	//	信道：信道缓存中队列（未读取）元素的数量；若 v 为 nil，len(v) 即为零。
	fmt.Printf("len(\"123\"): %v\n", len("你好啊!"))

	// cap 内建函数返回 v 的容量，这取决于具体类型：
	//	数组：v 中元素的数量（与 len(v) 相同）。
	//	数组指针：*v 中元素的数量（与 len(v) 相同）。
	//	切片：在重新切片时，切片能够达到的最大长度；若 v 为 nil，len(v) 即为零。
	//	信道：按照元素的单元，相应信道缓存的容量；若 v 为 nil，len(v) 即为零。

	// make 内建函数分配并初始化一个类型为切片、映射、或（仅仅为）信道的对象。
	// 与 new 相同的是，其第一个实参为类型，而非值。不同的是，make 的返回类型
	// 与其参数相同，而非指向它的指针。其具体结果取决于具体的类型：
	//	切片：size 指定了其长度。该切片的容量等于其长度。第二个整数实参可用来指定
	//		不同的容量；它必须不小于其长度，因此 make([]int, 0, 10) 会分配一个长度为0，
	//		容量为10的切片。
	//	映射：初始分配的创建取决于 size，但产生的映射长度为0。size 可以省略，这种情况下
	//		就会分配一个小的起始大小。
	//	信道：信道的缓存根据指定的缓存容量初始化。若 size 为零或被省略，该信道即为无缓存的。

	// new 内建函数分配内存。
	// 其第一个实参为类型，而非值，其返回值为指向该类型的新分配的零值的指针。

	//close():
	// close 内建函数关闭信道，该信道必须为双向的或只发送的。
	// 它应当只由发送者执行，而不应由接收者执行，其效果是在最后发送的值被接收后停止该信道。
	// 在最后一个值从已关闭的信道 c 中被接收后，任何从 c 的接收操作都会无阻塞成功，
	// 它会返回该信道元素类型的零值。对于已关闭的信道，形式
	//	x, ok := <-c
	// 还会将 ok 置为 false。

	// panic()

	//recover():recover()的调用仅当它在defer函数中 被直接调用 时才有效。

	// error 内建接口类型是表示错误情况的约定接口，nil 值即表示没有错误。
	// type error interface {
	// 		Error() string
	// 	}

}

func TestEg22(t *testing.T) {

	a := []int{1, 2}
	b := append(a, 3)
	c := append(b, 4)
	d := append(b, 5)
	fmt.Println(a, b, c[3], d[3])
	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(b), cap(b), b)
	fmt.Println(len(c), cap(c), c)
	fmt.Println(len(d), cap(d), d)
}
