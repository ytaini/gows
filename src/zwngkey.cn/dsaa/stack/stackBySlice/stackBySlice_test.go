/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-27 23:47:15
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-29 15:02:18
 * @Description:
	基于slice 实现栈1
*/
package sbs

import (
	"fmt"
	"testing"
)

//基于切片
type Stack struct {
	datas []any
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Print() {
	for i := len(s.datas) - 1; i >= 0; i-- {
		fmt.Printf("↓->%d\n", s.datas[i])
	}
}

func (s *Stack) Push(data any) {
	s.datas = append(s.datas, data)
}

func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	data := s.datas[len(s.datas)-1]
	s.datas = s.datas[0 : len(s.datas)-1]
	return data
}

func (s *Stack) IsEmpty() bool {
	return len(s.datas) == 0
}

func Test1(t *testing.T) {
	s := NewStack()
	for i := 0; i < 100; i++ {
		s.Push(i)
	}
	s.Print()

	for i := 0; i < 101; i++ {
		s.Pop()
	}
}
