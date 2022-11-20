/*
  - @Author: zwngkey
  - @Date: 2022-04-26 15:23:16

* @LastEditTime: 2022-11-20 15:09:29
  - @Description:
    defer语句的一些易错点
*/
package godefer

import (
	"fmt"
	"testing"
)

func Test4(t *testing.T) {
	fmt.Println(f1()) //2
	fmt.Println(f2()) //1
	fmt.Println(f3()) //1
	fmt.Println(f4()) //2
	fmt.Println(f5()) //2
}

func f1() (n int) {
	n = 1
	defer func() {
		n++
	}()
	return n
}

func f2() int {
	n := 1
	defer func() {
		n++
	}()
	//没有声明返回值的参数名的情况:
	//真正返回的不是n中的值.
	//此时会创建一个匿名的返回值,假设为a
	//return n 的操作为2步
	// 1.a = n
	// 2.return a
	return n
}
func f3() (n int) {
	n = 1
	defer func(n int) {
		n++ //修改的是局部变量n
	}(n)
	return n
}

func f4() (n int) {
	n = 1
	defer func(n *int) {
		*n++
	}(&n)
	return n
}

type N struct {
	x int
}

func f5() *N {
	n := &N{1}
	defer func() {
		n.x++
	}()
	return n
}
