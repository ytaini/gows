/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 16:58:35
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-01 12:07:56
 * @Description:
	使用双向链表实现队列
*/
package circlequeue

import (
	"fmt"

	list "zwngkey.cn/dsaa/LinkedList/doublyLinkList"
)

type Queue struct {
	l *list.DList
}

func NewQueue() *Queue {
	return &Queue{l: list.NewList()}
}

func (q *Queue) IsEmpty() bool {
	return q.l.Len() == 0
}
func (q *Queue) Len() int {
	return q.l.Len()
}

func (q *Queue) Push(data any) {
	q.l.InsertBack(data)
}

func (q *Queue) Pop() (any, error) {
	if q.Len() == 0 {
		return nil, fmt.Errorf("无数据")
	}
	return q.l.DeleteByIndex(1), nil
}

func (q *Queue) Peek() any {
	return q.l.First()
}

func (q *Queue) Print() {
	q.l.PrintList()
}
