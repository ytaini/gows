/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 18:42:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 18:42:04
 * @Description:
 */

package singlylooplinklist

import (
	"fmt"
	"testing"
)

func createList() *List {
	l := NewList()
	l.InsertAfterNode(NewNode(1))
	l.InsertAfterNode(NewNode(2))
	l.InsertAfterNode(NewNode(3))
	l.InsertAfterNode(NewNode(4))
	l.InsertAfterNode(NewNode(7))
	return l
}

func Test1(t *testing.T) {
	l := createList()
	l.Show()
	fmt.Println(l.tail.data)
}

func Test2(t *testing.T) {
	l := createList()
	l.Show()
	// data, err := l.FindValueByIndex(-1)
	// data, err := l.FindValueByIndex(5)
	// data, err := l.FindValueByIndex(7)
	data, err := l.FindValueByIndex(2)
	fmt.Println(data, err)
}

func Test4(t *testing.T) {
	l := createList()
	l.InsertBeforeValue(8)
	l.InsertBeforeValue(9)
	l.InsertBeforeValue(5)
	l.InsertBeforeValue(6)
	l.InsertAfterValue(10)
	l.Show()
	fmt.Println(l.tail.data)
	fmt.Println(l.tail.next.next.data)
}

func Test3(t *testing.T) {
	l := createList()
	l.Show()
	// i, err := l.FindIndexByValue(3)
	// i, err := l.FindIndexByValue(7)
	i, err := l.FindIndexByValue(19)
	fmt.Println(i, err)
}

func Test5(t *testing.T) {
	l := createList()
	// l.InsertValueByIndex(5, 7)
	// l.InsertValueByIndex(6, -1)
	// l.InsertValueByIndex(8, 0)
	// l.InsertValueByIndex(9, 2)
	l.InsertValueByIndex(2, 1)

	l.Show()
	fmt.Println(l.tail.data)
}

func Test6(t *testing.T) {
	l := createList()
	l.DeleteBefore()
	l.Show()
	l.DeleteBefore()
	l.Show()
}

func Test7(t *testing.T) {
	l := createList()
	l.DeleteBack()
	l.Show()
	l.DeleteBack()
	l.Show()
}

func Test8(t *testing.T) {
	l := createList()
	l.Show()
	fmt.Println("-------")
	l.Reverse()
	l.Show()
	fmt.Println(l.tail.data)
}

func Test9(t *testing.T) {
	l1 := createList()
	l2 := createList()
	l1.Join(l2)
	l1.Show()
	fmt.Println(l1.tail.data)
}
