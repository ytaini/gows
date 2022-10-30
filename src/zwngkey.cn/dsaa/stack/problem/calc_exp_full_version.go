/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-30 04:00:36
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 18:53:42
 * @Description:
	完整版的逆波兰计算器，功能包括如下：
		支持 + - * / ( )
		多位数，支持小数
		兼容处理, 过滤任何空白字符，包括空格、制表符、换页符
	无法处理操作数有负数.
*/
package problem

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"testing"

	queue "zwngkey.cn/dsaa/queue/circleQueue"
	stack "zwngkey.cn/dsaa/stack"
)

const LEFT string = "("
const RIGHT string = ")"
const ADD string = "+"
const MINUS string = "-"
const TIMES string = "*"
const DIVISION string = "/"

const DEFAULT_LEVEL int = -1
const LEVEL1 int = 1
const LEVEL2 int = 2
const MAX_LEVEL int = math.MaxInt

var s = stack.New[string]()
var q = queue.NewQueue()

// 计算中缀表达式结果
func CalcExp(exp string) (res float64, err error) {
	// 去空 12.8+(2-3.55)*4+10/5.0+8
	exp = trimSpace(exp)

	// 取操作数,操作符
	s := split(exp)

	// 中缀表达式转后缀表达式
	s = InfixToSuffix(s)

	// 计算后缀表达式
	res, err = calcRPN(s)

	return res, err
}

// 计算后缀表达式的结果
func calcRPN(sufExp []string) (float64, error) {
	// s := stack.New[string]()
	for _, v := range sufExp {
		//过滤数值
		if isNumber(v) {
			s.Push(v)
		} else {
			v1, _ := s.Pop()
			v2, _ := s.Pop()

			v3, _ := strconv.ParseFloat(v1, 64)
			v4, _ := strconv.ParseFloat(v2, 64)

			//注意: v3,v4传入顺序.
			v5 := calc(v4, v3, v)
			s.Push(strconv.FormatFloat(v5, 'f', -1, 64))
		}
	}
	v1, _ := s.Pop()
	return strconv.ParseFloat(v1, 64)
}

// 将中缀表达式转后缀表达式
func InfixToSuffix(infixExp []string) (suffixExp []string) {

	for _, item := range infixExp {
		// 如果v1能转为int,说明v1为操作数,否则为操作符
		// _, err := strconv.Atoi(v1)
		reg, _ := regexp.Compile(`\d+`)

		if reg.MatchString(item) {
			// 如果v1为操作数
			q.Push(item)
		} else {
			// 如果v1为操作符

			// 如果v1为括号
			if isParentheses(item) {
				parenthesesHandle(item)
			} else {
				//如果为运算符
				operHandle(item)
			}
		}
	}
	for !s.IsEmpty() {
		v1, _ := s.Pop()
		q.Push(v1)
	}
	for !q.IsEmpty() {
		v1, _ := q.Pop()
		suffixExp = append(suffixExp, v1.(string))
	}
	return
}

// 操作符为括号时,处理逻辑
func parenthesesHandle(v1 string) {
	if v1 == RIGHT {
		cur, _ := s.Pop()
		for cur != LEFT {
			q.Push(cur)
			cur, _ = s.Pop()
		}
	}
	if v1 == LEFT {
		s.Push(v1)
	}
}

// 操作符为运算符时,处理逻辑
func operHandle(v1 string) {
	v3, _ := s.Peek()
	if s.IsEmpty() || v3 == LEFT {
		s.Push(v1)
	} else if operPriority(v1) > operPriority(v3) {
		s.Push(v1)
	} else {
		v4, _ := s.Pop()
		q.Push(v4)
		operHandle(v1)
	}
}

// 将exp各个操作数,操作符分别加入s中
func split(exp string) (s []string) {
	index := 0
	for index != len(exp) {
		ch := rune(exp[index])
		if ch >= '0' && ch <= '9' {
			j := index + 1
			for j != len(exp) && !IsOp(string(exp[j])) {
				j++
			}
			v1 := exp[index:j]
			index = j
			s = append(s, v1)
			continue
		} else if IsOp(string(ch)) {
			s = append(s, exp[index:index+1])
		} else {
			panic("错误表达式")
		}
		index++
	}
	return s
}

// 去除表达式所有空格
func trimSpace(exp string) string {
	pat := `\s+`
	reg, _ := regexp.Compile(pat)
	return reg.ReplaceAllString(exp, "")
}

// 判断是不是数字 int double long float
func isNumber(str string) bool {
	pat := `^[-+]?\d+([.]\d+)?$`
	reg, _ := regexp.Compile(pat)
	return reg.MatchString(str)
}

// 返回操作符优先级
func operPriority(str string) int {
	switch str {
	case ADD, MINUS:
		return LEVEL1
	case TIMES, DIVISION:
		return LEVEL2
	case RIGHT, LEFT:
		return MAX_LEVEL
	default:
		return DEFAULT_LEVEL
	}
}

// 判断是否为操作符
func IsOp(op string) bool {
	return regexp.MustCompile(`\+|\-|\*|\/|\(|\)`).MatchString(op)
}

// 运算
func calc(a, b float64, op string) float64 {
	switch op {
	case ADD:
		return a + b
	case MINUS:
		return a - b
	case TIMES:
		return a * b
	case DIVISION:
		return a / b
	default:
		return 0
	}
}

// 判断操作符是否为括号
func isParentheses(op string) bool {
	return op == RIGHT || op == LEFT
}

// 无法处理操作数有负数.
func Test11(t *testing.T) {
	exp := "12.8 + (2 - 3.55)*4+10/5.0 + 8*(-1)"
	pat := `[-+]?\d+([.]\d+)?`
	reg := regexp.MustCompile(pat)

	// [12.8 2 3.55 4 +10 5.0 8 -1]
	fmt.Printf("reg.FindAllString(exp, -1): %v\n", reg.FindAllString(exp, -1))

	// [12.8 2 -3.55 4 +10 5.0 +8 -1]
	exp = trimSpace(exp)
	fmt.Printf("reg.FindAllString(exp, -1): %v\n", reg.FindAllString(exp, -1))
}
