/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-29 15:35:13
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 01:57:41
 * @Description:
	使用栈完成中缀表达式计算(不带括号,规范化表达式),运算符包括(加减乘除)如
		"1000*2+2000/200*400+514"
		"33+2+2*6-2"
		"7*22*2-5+1-5+3-4"

	使用栈完成中缀表达式计算的思路
	1. 初始化两个栈，一个【数栈】和一个【符号栈】
	2. 通过一个index值（索引），来遍历我们的表达式
	3. 如果发现是一个数字，就直接入【数栈】
	4. 如果发现是一个运算符，就分情况讨论
	1. 如果当前的【符号栈】为空，那就直接入栈
	2. 如果当前的【符号栈】有操作符，就进行比较
		1. 【当前的操作符】的优先级 <= 【符号栈中操作符】的优先级，就需要从【数栈】中pop出两个数，再从【符号栈】中pop出一个符号，
			然后进行运算，将得到结果入【数栈】，**再递归调用步骤3**
		2. 【当前的操作符】的优先级 > 【符号栈中操作符】的优先级，就直接入符号栈
	5. 当表达式扫描完毕，就顺序的从【数栈】和【符号栈】中pop相应的数和符号，并运行
	6. 最从【数栈】中只有一个数字就是表达式的结果

*/
package problem

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	stack "zwngkey.cn/dsaa/stack"
)

func Test3(t *testing.T) {
	expression := "1000*2+2000/200*400+514"
	expression1 := "33+2+2*6-2"
	expression2 := "7*22*2-5+1-5+3-4"
	expression3 := "4/2*3-4*2-3-99"
	expression4 := "1*1*1*3*2/3"
	expression5 := "11*1*1*3*2/3"
	expression6 := "1000*23"
	fmt.Println(calc2(expression))
	fmt.Println(calc2(expression1))
	fmt.Println(calc2(expression2))
	fmt.Println(calc2(expression3))
	fmt.Println(calc2(expression4))
	fmt.Println(calc2(expression5))
	fmt.Println(calc2(expression6))
}

/*
	当处理多位数时,不能发现是一个数就立即入栈,因为可能是多位数.
	在处理数,需要向expresion的表达式的index后再看一位,如果是数进行扫描,如果是符号再能入栈.
*/
func calc2(expression string) int {
	//存操作数
	numStack := stack.New[int]()
	//存操作符
	operStack := stack.New[int]()

	// 记录字符串下标
	inde := 0
	for inde != len(expression) {
		v := expression[inde]

		if !isOper(rune(v)) {
			//处理数字入栈的情况，包含处理多位数的情况
			j := inde + 1
			for j != len(expression) && !isOper(rune(expression[j])) {
				j++
			}
			v1 := expression[inde:j]
			val, _ := strconv.Atoi(v1)
			numStack.Push(val)
			inde = j
			continue
		} else {
			//处理操作符入栈的情况，
			operSolve(rune(v), numStack, operStack)
		}
		inde++
	}
	v1, _ := numStack.Pop()
	v2, _ := numStack.Pop()
	oper, _ := operStack.Pop()
	v3 := cal(v1, v2, rune(oper))
	numStack.Push(v3)
	val, _ := numStack.Pop()
	return val
}

func operSolve(op rune, numStack, operStack *stack.Stack[int]) {
	if !operStack.IsEmpty() {
		op1, _ := operStack.Peek()
		p1 := priority(op)
		p2 := priority(rune(op1))
		// 当运算符op优先级比栈中运算符优先级低时,
		// 递归比较当前运算符op与栈中所有存在的运算符,直到栈空.
		if p1 <= p2 {
			v1, _ := numStack.Pop()
			v2, _ := numStack.Pop()
			op2, _ := operStack.Pop()
			v3 := cal(v1, v2, rune(op2))
			numStack.Push(v3)

			operSolve(op, numStack, operStack)
		} else {
			operStack.Push(int(op))
		}
	} else {
		operStack.Push(int(op))
	}
}

func Test1(t *testing.T) {
	// expression := "1-2+3"
	expression := "4/2*3-4*2-3-9"
	fmt.Println(calc1(expression))
}

// 只能处理一位数的计算
func calc1(expression string) int {
	//存操作数
	numStack := stack.New[int]()
	//存操作符
	operStack := stack.New[int]()

	for _, ch := range expression {
		// fmt.Printf("%c\n", v)
		if !isOper(ch) {
			val, _ := strconv.Atoi(fmt.Sprintf("%c", ch))
			numStack.Push(val)
			// numStack.Push(int(v - '0'))
		} else {
			operSolve(ch, numStack, operStack)
		}
	}
	v1, _ := numStack.Pop()
	v2, _ := numStack.Pop()
	oper, _ := operStack.Pop()
	v3 := cal(v1, v2, rune(oper))
	numStack.Push(v3)
	val, _ := numStack.Pop()
	return val
}

// 判断(+-*/)运算符优先级
func priority(oper rune) int {
	switch oper {
	case '*', '/':
		return 1
	case '+', '-':
		return 0
	}
	return -1
}

// 判断字符是否为操作符
func isOper(val rune) bool {
	return val == '+' || val == '-' || val == '*' || val == '/'
}

// 计算
func cal(a, b int, oper rune) int {
	var res int
	switch oper {
	case '+':
		res = a + b
	case '-':
		res = b - a // a是先弹出来的数，为被减数
	case '*':
		res = a * b
	case '/':
		res = b / a // a是先弹出来的数，为被除数
	}
	return res
}

// 其他验证
func Test2(t *testing.T) {
	expression := "30+20*6-2"
	for _, v := range expression {
		fmt.Println(int(v), v)
	}
}

// 解析出一个表达式中的操作数与操作符
func ParseString(exp string) (nums, opers [][]string) {
	regp1 := "\\d+"
	reg := regexp.MustCompile(regp1)
	res := reg.FindAllStringSubmatch(exp, -1)

	regp2 := "\\D+"
	reg2 := regexp.MustCompile(regp2)
	res1 := reg2.FindAllStringSubmatch(exp, -1)
	return res, res1
}
