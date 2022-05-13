/*
 * @Author: zwngkey
 * @Date: 2021-12-25 17:21:09
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:55:38
 * @Description:
 */
package eggot1

import (
	"fmt"
	"testing"
)

//为什么B类型只继承A类型的方法,没有继承*A类型的方法?

type A int

func (a A) Value() int {
	return int(a)
}

func (a *A) Set(n int) {
	*a = A(n)
}

//对于B而言,它继承*A的方法是没有意义的.
type B struct {
	A
	b int
}

type C struct {
	*A
	c int
}

func Test1(te *testing.T) {
	var a A = 1
	var b B = B{a, 2}
	a.Set(3)
	fmt.Printf("a.Value(): %v\n", a.Value()) // 3
	b.Set(2)                                 //(&b).Set(1)
	fmt.Printf("a.Value(): %v\n", a.Value()) // 3
	fmt.Printf("b.Value(): %v\n", b.Value()) //b.Value() --> A.Value(b.A) 此时Value方法操作的是b变量的副本.

	var a1 A = 1
	var c C = C{&a1, 2}
	a1.Set(3)
	fmt.Printf("a1.Value(): %v\n", a1.Value()) // 3
	c.Set(2)
	fmt.Printf("a1.Value(): %v\n", a1.Value()) // 2
	fmt.Printf("c.Value(): %v\n", c.Value())   //c.Value() --> *A.Value(&(b.A))
}
