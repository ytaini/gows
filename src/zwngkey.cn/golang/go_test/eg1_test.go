/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-08 05:22:08
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-08 05:27:40
 */

package gotest

import "fmt"

func Print() string {
	return "print"
}

// 示例函数的使用
func ExamplePrint() {
	fmt.Println(Print())
	// Output:
	// print1
}

// 输出

// === RUN   ExamplePrint
// --- FAIL: ExamplePrint (0.00s)
// got:
// print
// want:
// print1
// FAIL
// FAIL	zwngkey.cn/golang/go_test	0.302s
