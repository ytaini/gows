/*
 * @Author: zwngkey
 * @Date: 2021-12-25 17:21:09
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 09:44:01
 * @Description:
 */
package eggot1

import (
	"fmt"
	"testing"
)

// 为什么B类型只继承A类型的方法,没有继承*A类型的方法?
// 对于B而言,它继承*A的方法是没有意义的.
type A int

func (a A) Value() int {
	return int(a)
}

func (a *A) Set(n int) {
	*a = A(n)
}

// A 类型有的方法: Value(a A)
//*A 类型有的方法: Value(a *A),Set(a *A,n int)

type B struct {
	A
	b int
}

func Test1(te *testing.T) {
	var a A = 1
	var b B = B{a, 2}
	a.Set(3)
	fmt.Printf("a.Value(): %v\n", a.Value()) //3
	fmt.Printf("b.Value(): %v\n", b.Value()) //1
}

type C struct {
	*A
	c int
}

func Test3(t *testing.T) {
	var a1 A = 1
	var c C = C{&a1, 2}
	a1.Set(3)
	fmt.Printf("a1.Value(): %v\n", a1.Value()) // 3
	fmt.Printf("c.Value(): %v\n", c.Value())   // 3
}
