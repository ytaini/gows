/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-29 15:03:08
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-29 15:15:16
 * @Description:

	基于数组实现栈
	1.使用数组储存底层数据
	2.定义一个top,来表示栈顶,初始值为-1,记录栈顶所在的下标.
	 	定义一个maxSize,记录栈中的元素个数.
	3.入栈: top++;stack[top]=data
	4.出栈: value:=stack[top];top--;
	5.判空:top == -1
	6.判满:top == maxSize -1
*/
package sby

import (
	"fmt"
	"testing"
)

type Stack struct {
	datas   [5]any
	maxSize int
	top     int //初始值-1 记录栈顶所在的下标.
}

func New() *Stack {
	return &Stack{maxSize: 5, top: -1}
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.maxSize-1
}
func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("栈空")
	}
	val := s.datas[s.top]
	s.top--
	return val, nil
}
func (s *Stack) Push(data any) error {
	if s.IsFull() {
		return fmt.Errorf("栈满")
	}
	s.top++
	s.datas[s.top] = data
	return nil
}
func (s *Stack) Print() {
	for i := s.top; i >= 0; i-- {
		fmt.Printf("↓->%v\n", s.datas[i])
	}
}

func Test1(t *testing.T) {
	s := New()
	for i := 0; i < 6; i++ {
		err := s.Push(i)
		if err != nil {
			fmt.Println(err)
		}
	}
	s.Print()

	for i := 0; i < 6; i++ {
		data, err := s.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
	}

}
