/*
 * @Author: zwngkey
 * @Date: 2022-05-14 06:33:28
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-14 06:39:01
 * @Description:
	runtime.Caller函数.
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)

	funcName := runtime.FuncForPC(pc).Name()
	_ = funcName
}
