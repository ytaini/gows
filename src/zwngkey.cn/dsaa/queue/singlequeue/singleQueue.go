/*
 * @Author: imzw
 * @Date: 2022-10-26 09:05:33
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 18:43:30
 * @Description:
 	基于数组模拟顺序队列.
		缺点: 当前现实的队列只能使用一次,不能复用.
*/
package singlequeue

import (
	"errors"
	"fmt"
)

type Queue struct {
	maxSize int
	buf     [5]int // 数组
	front   int
	rear    int
}

func NewQueue() *Queue {
	return &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
}

//显示
func (q Queue) ShowQueue() {
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("%d ", q.buf[i])
	}
	fmt.Println()
}

//获取
func (q *Queue) Pop() (val int, err error) {
	if q.front == q.rear {
		return -1, errors.New("队列为空")
	}
	q.front++
	return q.buf[q.front], err
}

//添加
func (q *Queue) Push(elem int) (err error) {
	if q.rear == q.maxSize-1 {
		return errors.New("队列已满")
	}
	q.rear++
	q.buf[q.rear] = elem
	return
}
