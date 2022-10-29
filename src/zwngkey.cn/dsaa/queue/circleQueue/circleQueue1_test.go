/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 19:19:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 03:34:47
 * @Description:

 */
package circlequeue

import (
	"fmt"
	"testing"
)

func Test10(t *testing.T) {
	q := NewQueue()
	q.Push(0)
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Len())
	fmt.Println(q.Peek())
	q.Print()

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Peek())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	q.Print()
}
