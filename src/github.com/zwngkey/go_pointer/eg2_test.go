/*
 * @Author: zwngkey
 * @Date: 2022-05-01 21:25:36
 * @LastEditTime: 2022-05-03 17:29:11
 * @Description: go 中取地址符&
 */
package gopointer

import (
	"fmt"
	"testing"
)

/*
	哪些值可以被取地址，哪些值不可以被取地址？
		以下的值是不可以寻址的：
			字符串的字节元素
			映射元素
			接口值的动态值（类型断言的结果）
			常量（包括有名常量和字面量）
			声明的包级别函数
			方法（用做函数值）
			中间结果值:
				函数调用
				显式值转换
				各种操作，不包含指针解引用（dereference）操作，但是包含：
					通道接收操作
					子字符串操作
					子切片操作
					加法、减法、乘法、以及除法等等。

		注意：&T{}在Go里是一个语法糖，它是tmp := T{}; (&tmp)的简写形式。 所以&T{}是合法的并不代表字面量T{}是可寻址的。
*/

/*
	go 中一些不可寻址的值.
*/
func Test1(t *testing.T) {
	q1 := &map[string]int{
		"zs": 1,
	}
	_ = q1

	var s string = "string中"
	var a byte = s[5]
	// 字符串的字节元素 不可寻址
	// var b = &s[1]
	_ = a

	var b []byte = []byte("string")
	var c *byte = &b[1]
	_ = c

	type name = string
	type age = int

	var d = map[name]age{
		"李四": 18,
		"张三": 20,
	}
	var e = d["张三"]
	_ = e
	//映射元素 不可寻址
	// var f = &d["张三"]

	var f any = "string"
	g, ok := f.(string)
	_, _ = g, ok
	// 接口值的动态值（类型断言的结果）不可寻址
	// var h = &f.(string)

	const h = 1
	_ = h
	//常量（包括有名常量和字面量）不可寻址
	// var i = &h
	// var i = &1

	var j = f1
	_ = j
	//声明的包级别函数 不可寻址
	// var k = &f1

	//匿名函数可以
	f2 := func() int { fmt.Println("f2"); return 1 }
	var l (*func() int) = &f2
	// _ = l
	(*l)()

	var m = MyInt(1)
	var n = m.print
	_ = n
	// 方法（用做函数值）不可寻址
	// var o = &m.print

	var o = f2()
	_ = o
	// 函数调用 不可寻址
	// var p = &f2()

	//显式值转换 不可寻址
	// var p = &MyInt(1)

	var p = make(chan int)
	p <- 1
	var q = <-p
	_ = q
	// 通道接收操作 不可寻址
	// var r = &<-p

	var r = "hi,zhangsan"
	var s1 = r[2:]
	_ = s1
	// 子字符串操作 不可寻址
	// var s2 = &r[2:]

	var b1 = b[1:]
	_ = b1
	// 子切片操作 不可寻址
	// var b2 = &b[1:]

	// 加法、减法、乘法、以及除法等等。不可寻址
	var a1 = 1
	var a2 = 2
	var a3 = a1 + a2
	_ = a3
	// var a4 = &(a1 + a2)
	// var a4 = &(1 + 1)
	// var a4 = &(a1 - a2)
	// var a4 = &(1 - 1)
	// var a4 = &(a1 * a2)
	// var a4 = &(1 * 1)
	// var a4 = &(1/2)
	// var a4 = &(a1 / a2)
	// var a4 = &(3 % 2)
	// var a4 = &(a1 % a2)
	// var a4 = &(a1 | a2)
	// var a4 = &(a1 & a2)
	// var a4 = (&a1 &^ a2)
	// var a4 = &(^a2)
	// var a4 = &(true || false)
	// var a4 = &(true && false)
	// var a4 = &(!true)
}

func f1() { fmt.Println("f1") }

type MyInt int

func (m MyInt) print() {
	fmt.Println(m)
}

/*
	以下的值是可寻址的，因此可以被取地址：
		变量
		可寻址的结构体的字段
		可寻址的数组的元素
		任意切片的元素（无论是可寻址切片或不可寻址切片）
		指针解引用（dereference）操作
*/

func Test2(t *testing.T) {
	//变量可寻址
	var a = 1
	var b = &a
	_ = b

	//可寻址的结构体的字段
	var p Person = Person{12, "lisi"}
	age := &p.Age
	_ = age

	//语法糖
	var per *Person = &Person{12, "lisi"}
	// age2 := &((*per).Age)
	// age3 := &(*per).Age
	age1 := &per.Age
	_ = age1

	//可寻址的数组的元素
	var arr = [...]int{1, 2, 3}
	p1 := &arr[0]
	_ = p1

	//语法糖
	var parr = &[...]int{1, 2, 3}
	// ppp := &((*parr)[0])
	// ppp := &(*parr)[0]
	p2 := &parr[0]
	_ = p2

	//任意切片的元素（无论是可寻址切片或不可寻址切片）
	var sli = []int{1, 2, 3}
	fmt.Println(&sli[0])
	fmt.Println(&sli[1])
	fmt.Println(&sli[2])

	var sli1 = &[]int{1, 2, 3}
	fmt.Println(&(*sli1)[0])

	//指针解引用（dereference）操作
	var person = &Person{12, "a"}
	p3 := &(*person)
	_ = p3
	p4 := &*&*&*&*&*&*person
	_ = p4

}

type Person struct {
	Age  int
	Name string
}
