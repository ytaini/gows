/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-30 00:45:35
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 18:27:25
 * @Description:

 	注意: 无法处理操作数有负数.
		如: (3+4)*5+6*(-1)

	中缀表达式转逆波兰式

	中缀表达式转逆波兰式算法:
		1. 初始化两个栈：运算符栈s1和储存中间结果的栈s2
		2. 从左至右扫描中缀表达式
		3. 遇到操作数时，将其压s2
		4. 遇到运算符时，比较其与s1栈顶运算符的优先级：
			1. 如果s1为空，或栈顶运算符为左括号“(”，则直接将此运算符入栈
			2. 否则，若优先级比栈顶运算符的高，也将运算符压入s1
			3. 否则，将s1栈顶的运算符弹出并压入到s2中，再次与s1中新的栈顶运算符相比
		5. 遇到括号时：
			1. 如果是左括号“(”，则直接压入s1
			2. 如果是右括号“)”，则依次弹出s1栈顶的运算符，并压入s2，直到遇到左括号为止，此时将这一对括号丢弃
		6. ........
		7. 重复步骤2至5，直到表达式的最右边
		8. 将s1中剩余的运算符依次弹出并压入s2
		9. 依次弹出s2中的元素并输出，结果的逆序即为中缀表达式对应的后缀表达式

*/
package problem

import (
	"math"
	"regexp"

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

// s1 := stack.New[string]()
var q = queue.NewQueue()

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
	// s = Reverse(s)
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

// 操作符为括号时,处理逻辑
// func parenthesesHandle(v1 string, s1, s2 *stack.Stack[string]) {
// 	if v1 == ")" {
// 		cur, _ := s1.Pop()
// 		for cur != "(" {
// 			s2.Push(cur)
// 			cur, _ = s1.Pop()
// 		}
// 	}
// 	if v1 == "(" {
// 		s1.Push(v1)
// 	}
// }

// 操作符为运算符时,处理逻辑
// func operHandle(v1 string, s1, s2 *stack.Stack[string]) {
// 	v3, _ := s1.Peek()
// 	if s1.IsEmpty() || v3 == "(" {
// 		s1.Push(v1)
// 	} else if prio(v1) > prio(v3) {
// 		s1.Push(v1)
// 	} else {
// 		v4, _ := s1.Pop()
// 		s2.Push(v4)
// 		operHandle(v1, s1, s2)
// 	}
// }

// 判断是否为操作符
func IsOp(op string) bool {
	return regexp.MustCompile(`\+|\-|\*|\/|\(|\)`).MatchString(op)
}

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

// 判断操作符是否为括号
func isParentheses(op string) bool {
	return op == RIGHT || op == LEFT
}

// 反转切片
func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 将exp各个操作数,操作符分别加入s中
func Split(exp string) (s []string) {
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
