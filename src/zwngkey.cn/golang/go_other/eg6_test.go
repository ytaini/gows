/*
 * @Author: zwngkey
 * @Date: 2022-05-12 21:32:13
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-12 21:40:37
 * @Description: go 内联函数
 */
package goother

import (
	"fmt"
	"testing"
)

/*
	Go 语言的编译器也会对函数调用进行优化，但是他没有提供任何关键字可以手动声明内联函数，
		不过我们可以在函数上添加//go:noinline注释告诉编译器不要对它进行内联优化。

	我们在编译代码时传入--gcflags=-m 参数可以查看编译器的优化策略，传入--gcflags="-m -m"会查看更完整的优化策略！

	Go在内部维持了一份内联函数的映射关系，会生成一个内联树，我们可以通过-gcflags="-d pctab=pctoinline"参数查看
*/

func Test61(t *testing.T) {
	s := []int{10, 12, 3, 14}
	fmt.Println(GetMaxValue(s))
}

func GetMaxValue(s []int) int {
	max := 0
	for i := 0; i < len(s); i++ {
		max = maxValue(s[i], max)
	}
	return max
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
