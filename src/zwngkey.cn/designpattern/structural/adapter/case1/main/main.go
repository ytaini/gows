/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 02:10:34
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 02:15:00
 */
package main

import (
	"fmt"

	. "zwngkey.cn/designpattern/structural/adapter/case1"
)

func main() {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()

	want := "adaptee method"
	if res != want {
		fmt.Println("err")
	}
}
