/*
 * @Author: imzw
 * @Date: 2022-10-26 09:05:33
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-27 23:06:09
 * @Description: 	基于数组实现一次性队列.
 */
package singlequeue

import (
	"errors"
	"fmt"
	"testing"
)

type Queue struct {
	maxSize int
	buf     [5]int
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

func Test2(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 4; i++ {
		err := q.Push(i)
		if err != nil {
			fmt.Println(err)
		}
	}
	for i := 0; i < 2; i++ {
		_, err := q.Pop()
		if err != nil {
			fmt.Println(err)
		}
	}
}

//获取
func (q *Queue) Pop() (val int, err error) {
	if q.front == q.rear {
		return -1, errors.New("队列为空")
	}
	q.front++
	return q.buf[q.front], err
}

func Test1(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 4; i++ {
		err := q.Push(i)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(q.buf)
	}
	for i := 0; i < 5; i++ {
		val, err := q.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(val)
	}
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

//测试Push方法
func Test(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 6; i++ {
		err := q.Push(i)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(q.buf)
	}
}
