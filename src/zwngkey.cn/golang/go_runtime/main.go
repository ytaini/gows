/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 15:55:17
 * @Description:
	runtime.Caller函数.
*/

package main

import (
	"fmt"
	"runtime"
)

// Caller函数原型:
// func runtime.Caller(skip int) (pc uintptr, file string, line int, ok bool)
//
//	 参数：skip是要提升的堆栈帧数，0-当前函数，1-上一层函数，....
//		skip=0：Caller()会报告Caller()的调用者的信息
//		skip=1：Caller()回报告Caller()的调用者的调用者的信息
//		skip=2：...
//	 返回值：
//		pc:是函数指针
//		file:是函数所在文件名目录
//		line:所在行号
//		ok:是否可以获取到信息
func main() {
	f1()
}
func f1() {
	// pc, file, line, ok := runtime.Caller(0)

	// funcName := runtime.FuncForPC(pc).Name()
	// fmt.Println(funcName) // main.f1
	// fmt.Println(file)     // /Users/imzw/gows/src/zwngkey.cn/golang/go_runtime/logger.go
	// fmt.Println(line)     // 30
	// fmt.Println(ok)       // true

	// pc, file, line, ok := runtime.Caller(1)

	// funcName := runtime.FuncForPC(pc).Name()
	// fmt.Println(funcName) // main.main
	// fmt.Println(file)     // /Users/imzw/gows/src/zwngkey.cn/golang/go_runtime/logger.go
	// fmt.Println(line)     // 27
	// fmt.Println(ok)       // true

	f2()
}

func f2() {
	pc, file, line, ok := runtime.Caller(0)
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName) // main.f2
	fmt.Println(file)     // /Users/imzw/gows/src/zwngkey.cn/golang/go_runtime/logger.go
	fmt.Println(line)     // 50
	fmt.Println(ok)       // true

	// pc, file, line, ok := runtime.Caller(1)
	// funcName := runtime.FuncForPC(pc).Name()
	// fmt.Println(funcName) // main.f1
	// fmt.Println(file)     // /Users/imzw/gows/src/zwngkey.cn/golang/go_runtime/logger.go
	// fmt.Println(line)     // 46
	// fmt.Println(ok)       // true

	// pc, file, line, ok := runtime.Caller(2)
	// funcName := runtime.FuncForPC(pc).Name()
	// fmt.Println(funcName) // main.main
	// fmt.Println(file)     // /Users/imzw/gows/src/zwngkey.cn/golang/go_runtime/logger.go
	// fmt.Println(line)     // 27
	// fmt.Println(ok)       // true
}
