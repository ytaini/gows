/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-08 01:25:30
 */
package main

import "fmt"

//T和*T的方法集是啥关系?
//为什么要限制T和*T不能定义同名方法?

type T struct{}

func (t T) A() { fmt.Println("ABC") }

func (pt *T) B() {}

func main() {
	var t T
	var pt *T

	pt.A()
	(&t).A()
}
