/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-27 23:47:15
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 21:59:57
 * @Description:
	基于slice 实现栈1
*/
package sbs

import (
	"fmt"
	"sync"
	"testing"
)

type Stack struct {
	datas []any
	lock  sync.RWMutex
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Print() {
	fmt.Println(s.datas...)
}

func (s *Stack) Push(data any) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.datas = append(s.datas, data)
}

func (s *Stack) Pop() any {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.datas) == 0 {
		return nil
	}
	data := s.datas[len(s.datas)-1]
	s.datas = s.datas[0 : len(s.datas)-1]
	return data
}

func Test1(t *testing.T) {
	s := NewStack()
	fmt.Println(s.datas)
}
