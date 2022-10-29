/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-29 23:22:22
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 00:43:54
 * @Description:
	完成一个逆波兰计算器，要求完成如下任务:
		输入一个逆波兰表达式(后缀表达式)，使用栈(Stack), 计算其结果
		支持小括号和多位数整数，因为这里我们主要讲的是数据结构，因此计算器进行简化，只支持对整数的计算。
*/
package problem

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"zwngkey.cn/dsaa/stack"
)

func Test(t *testing.T) {
	// 逆波兰表达式
	// (3+4)*5-6 => 34+5*6-
	sufExp := "30 40 + 5 * 6 -"
	exp := strings.Split(sufExp, " ")
	res, _ := calcRpn(exp)
	fmt.Println("(3+4)*5-6 :", res)
}

func calcRpn(exp []string) (int, error) {
	s := stack.New[string]()
	for _, v := range exp {
		regex, _ := regexp.Compile(`\d+`)
		//过滤数值
		if regex.MatchString(v) {
			s.Push(v)
		} else {
			v1, _ := s.Pop()
			v2, _ := s.Pop()
			v3, _ := strconv.Atoi(v1)
			v4, _ := strconv.Atoi(v2)

			var res int
			switch v {
			case "+":
				res = v3 + v4
			case "-":
				res = v4 - v3
			case "*":
				res = v4 * v3
			case "/":
				res = v4 / v3
			default:
				return -1, fmt.Errorf("存在不支持操作符")
			}
			s.Push(strconv.Itoa(res))
		}
	}
	v1, _ := s.Pop()
	return strconv.Atoi(v1)
}
