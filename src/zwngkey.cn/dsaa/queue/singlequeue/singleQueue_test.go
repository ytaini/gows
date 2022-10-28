/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 18:44:09
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 18:44:09
 * @Description:
 */

package singlequeue

import (
	"fmt"
	"testing"
)

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
