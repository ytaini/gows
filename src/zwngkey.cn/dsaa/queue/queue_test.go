/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-01 12:17:43
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-01 12:45:18
 * @Description:

 */

package queue

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	q := New[int]()
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	fmt.Println(q.Peek())
	for i := 0; i < 6; i++ {
		fmt.Println(q.Pop())
	}

	// q1 := q.Copy()
	// fmt.Println(q.l)
	// fmt.Println(q1.l)
	// fmt.Println(q.l.Front())
	// fmt.Println(q1.l.Front())
}
