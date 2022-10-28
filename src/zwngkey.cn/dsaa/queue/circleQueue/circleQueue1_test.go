/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 19:19:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 21:32:26
 * @Description:

 */
package circlequeue

import (
	"fmt"
	"testing"
)

func Test10(t *testing.T) {
	q := NewQueue()
	q.InQueue(0)
	q.InQueue(1)
	q.InQueue(2)
	fmt.Println(q.Len())
	fmt.Println(q.Peek())
	q.Print()

	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.Peek())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.OutQueue())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.OutQueue())
	q.Print()
}
