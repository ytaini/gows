/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 18:45:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 18:45:05
 * @Description:
 */
package circlequeue

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	q := NewCQueue(5)
	for i := 0; i < 6; i++ {
		q.InQueue(i)
	}
	q.PrintQueue()

	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	q.PrintQueue()

	fmt.Println("----------")
	q.InQueue(5)
	q.PrintQueue()
	fmt.Println("----------")
	q.InQueue(6)
	q.PrintQueue()
}

func Test(t *testing.T) {
	q := NewCQueue(5)
	for i := 0; i < 6; i++ {
		q.InQueue(i)
	}
	q.PrintQueue()

	for i := 0; i < 5; i++ {
		fmt.Println(q.OutQueue())
	}
	q.PrintQueue()
}
