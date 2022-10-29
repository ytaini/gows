/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 22:00:52
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 22:29:31
 * @Description:
	基于链表实现stack
*/
package stack

import (
	"fmt"
)

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		data int
		prev *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) Peek() (int, error) {
	if s.length == 0 {
		return 0, fmt.Errorf("空栈")
	}
	return s.top.data, nil
}

func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return -1, fmt.Errorf("空栈")
	}
	v := s.top
	s.top = v.prev
	s.length--
	return v.data, nil
}

func (s *Stack) Push(data int) {
	newNode := &node{data: data, prev: s.top}
	s.top = newNode
	s.length++
}

func (s *Stack) Print() {
	if s.length == 0 {
		fmt.Println("空栈")
		return
	}
	for cur := s.top; cur != nil; cur = cur.prev {
		fmt.Printf("↓->%v\n", cur.data)
	}
}
