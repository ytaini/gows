/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 16:05:01
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 21:32:30
 * @Description:
	使用切片实现环形队列
*/
package circlequeue

import (
	"fmt"
)

// 定义队列
type CircleQueue struct {
	maxSize, //切片总长度, 实际容量为maxSize -1
	front, // 指向队列的第一个元素
	rear int //指向队列的最后一个元素的后一个位置.所以rear指向的位置永远不存值.
	arr []any
}

func NewCQueue(cap int) *CircleQueue {
	return &CircleQueue{
		maxSize: cap + 1,
		arr:     make([]any, cap+1),
	}
}

func (q *CircleQueue) IsFull() bool {
	return (q.rear+1)%q.maxSize == q.front
}

func (q *CircleQueue) IsEmpty() bool {
	return q.front == q.rear
}

// 队尾入队
func (q *CircleQueue) InQueue(data any) {
	if q.IsFull() {
		fmt.Println("队列满")
		return
	}
	q.arr[q.rear] = data
	q.rear = (q.rear + 1) % q.maxSize
}

// 队头出队
func (q *CircleQueue) OutQueue() (any, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("队列空")
	}
	data := q.arr[q.front]
	q.front = (q.front + 1) % q.maxSize
	return data, nil
}

// 获取队列头数据
func (q *CircleQueue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("队列空")
	}
	return q.arr[q.front], nil
}

// 打印队列
func (q *CircleQueue) PrintQueue() {
	if q.IsEmpty() {
		fmt.Println("空队列")
		return
	}
	for i := q.front; i < q.front+q.Len(); i++ {
		fmt.Printf("arr[%d]=%d\n", i%q.maxSize, q.arr[i%q.maxSize])
	}
}

// 获取队列有效数据个数
func (q *CircleQueue) Len() int {
	return (q.rear - q.front + q.maxSize) % q.maxSize
}
