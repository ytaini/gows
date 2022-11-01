/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-01 12:08:24
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-01 12:46:02
 * @Description:
	双向循环链表实现队列
*/
package queue

import (
	"container/list"
	"fmt"
)

type Queue[T any] struct {
	l *list.List
}

func New[T any]() *Queue[T] {
	return &Queue[T]{l: list.New()}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.l.Len() == 0
}
func (q *Queue[T]) Len() int {
	return q.l.Len()
}

func (q *Queue[T]) Push(data T) {
	q.l.PushBack(data)
}

func (q *Queue[T]) Pop() (t T, err error) {
	if q.Len() == 0 {
		return t, fmt.Errorf("无数据")
	}
	t = q.l.Remove(q.l.Front()).(T)
	return t, nil
}

func (q *Queue[T]) Peek() T {
	first := q.l.Front()
	return first.Value.(T)
}

func (q *Queue[T]) Copy() *Queue[T] {
	q1 := New[T]()
	q1.l.PushBackList(q.l)
	return q1
}
