/*
 * @Author: zwngkey
 * @Date: 2022-05-13 05:34:00
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:34:54
 * @Description:
 */
package main

import (
	"fmt"

	gooperator "zwngkey.cn/golang/go_basic/go_operator"
)

func main() {
	// gooperator.Test1()
	var a int8 = ^0
	fmt.Printf("%b\n", a) // -1
	gooperator.Testeg2()
	gooperator.Testeg3()
}
