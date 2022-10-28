/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 18:36:39
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 18:36:39
 * @Description:
 */

package dll

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	l := NewLinkList()

	l.PushFront(1)
	l.PushFront(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.ShowLinkList()
	fmt.Println("------")
	l.Reverse()
	l.ShowLinkList()

}

func Test(t *testing.T) {
	l := NewLinkList()

	l.PushFront(1)
	n2 := l.PushFront(2)

	n := l.PushBack(3)
	l.PushBack(4)
	n1 := l.PushBack(5)
	l.Reverse()

	v, err := l.Remove(n2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	l.ShowLinkList()

	fmt.Println(l.Front().data) //2
	fmt.Println(l.Back().data)  //5
	fmt.Println(l.GetLen())     //4

	l.MoveToBack(n)
	fmt.Println(l.Back().data)

	l.MoveToFront(n)
	fmt.Println(l.Front().data)

	l.ShowLinkList()

	fmt.Println("---------------")
	l.MoveAfter(n, n1)
	l.ShowLinkList()

	l.MoveBefore(n, n1)
	fmt.Println("---------------")
	l.ShowLinkList()

	l.InsertAfter(6, n)
	fmt.Println("---------------")
	l.ShowLinkList()

	l.InsertBefore(7, n)
	fmt.Println("---------------")
	l.ShowLinkList()

	l.Init()
	fmt.Println(l.IsEmpty())
	l.ShowLinkList()

}
