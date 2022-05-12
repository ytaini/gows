/*
 * @Author: zwngkey
 * @Date: 2022-05-13 05:34:00
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:33:38
 * @Description:
 */
package main

import (
	"fmt"

	goconstvar "zwngkey.cn/golang/go_basic/go_const_var"
)

func main() {
	// goconstvar.Testeg1()
	// goconstvar.Testeg2()
	fmt.Printf("%b\n", goconstvar.MaxUint2)
	fmt.Printf("%v\n", ^int(0))

}
