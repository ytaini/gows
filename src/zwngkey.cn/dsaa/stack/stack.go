/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-30 00:10:03
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 00:21:33
 * @Description:
	链表实现栈 + 泛型
*/
package stack

import (
	"fmt"
)

type (
	Stack[T any] struct {
		top    *node[T]
		length int
	}
	node[T any] struct {
		data T
		prev *node[T]
	}
)

func New[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

func (s *Stack[T]) Len() int {
	return s.length
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack[T]) Peek() (t T, error error) {
	if s.length == 0 {
		return t, fmt.Errorf("空栈")
	}
	return s.top.data, nil
}

func (s *Stack[T]) Pop() (t T, err error) {
	if s.length == 0 {
		return t, fmt.Errorf("空栈")
	}
	v := s.top
	s.top = v.prev
	s.length--
	return v.data, nil
}

func (s *Stack[T]) Push(data T) {
	newNode := &node[T]{data: data, prev: s.top}
	s.top = newNode
	s.length++
}

func (s *Stack[T]) Print() {
	if s.length == 0 {
		fmt.Println("空栈")
		return
	}
	for cur := s.top; cur != nil; cur = cur.prev {
		fmt.Printf("↓->%v\n", cur.data)
	}
}
