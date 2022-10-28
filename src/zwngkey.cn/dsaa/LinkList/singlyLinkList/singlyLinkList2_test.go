/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 18:38:44
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 22:34:37
 * @Description:
	测试 singlyLinkList.go
*/
package sll

import (
	"fmt"
	"testing"
)

func Test4(t *testing.T) {
	l := NewLLinkList()
	l.InsertBackValue(1)
	l.InsertBackValue(2)
	l.InsertBackValue(3)
	l.InsertBackValue(5)
	l.InsertBackValue(6)
	l.Show()

	fmt.Println("----")
	l.Reverse()
	l.Show()

	fmt.Println("----")
	l.Reverse2()
	l.Show()

	fmt.Println("----")
	l.ReversePrint2()

	fmt.Printf("l.Last().data: %v %v\n", l.Last().data, l.Last().next)
}

func Test3(t *testing.T) {
	l := NewLLinkList()
	l.InsertBackValue(1)
	l.InsertBackValue(2)
	l.InsertBackValue(3)
	l.InsertBeforeValue(4)
	l.InsertBeforeValue(5)
	l.InsertBeforeValue(6)
	l.InsertValueByIndex(4, 7)
	l.Show()
	fmt.Println("----------------")
	oldNode, _ := l.DeleteByIndex(1)
	fmt.Println(oldNode.data)
	l.Show()
	oldNode, _ = l.DeleteByIndex(2)
	fmt.Println(oldNode.data)
	l.Show()

	oldNode, _ = l.DeleteByIndex(4)
	fmt.Println(oldNode.data)
	l.Show()

	oldNode, _ = l.DeleteByIndex(5)
	fmt.Println(oldNode.data)
	l.Show()

	oldNode, _ = l.DeleteByIndex(5)
	fmt.Println(oldNode.data)
	l.Show()

	oldNode, err := l.DeleteByIndex(-1)
	fmt.Println(oldNode, err)
	l.Show()
}

func Test2(t *testing.T) {
	l := NewLLinkList()
	l.InsertBackValue(1)
	l.InsertBackValue(2)
	l.InsertBackValue(3)
	l.InsertBeforeValue(4)
	l.InsertBeforeValue(5)
	l.InsertBeforeValue(6)
	l.InsertValueByIndex(4, 7)
	l.Show()
	fmt.Println("----------------")
	for i := 0; i < 8; i++ {
		// oldNode, err := l.DeleteBefore()
		oldNode, err := l.DeleteBack()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(oldNode.data)
		}
	}
	l.Show()
}

func Test1(t *testing.T) {
	l := NewLLinkList()
	l.InsertBackValue(1)
	l.InsertBackValue(2)
	l.InsertBackValue(3)
	l.InsertBeforeValue(4)
	l.InsertBeforeValue(5)
	l.InsertBeforeValue(6)
	l.InsertValueByIndex(4, 7)
	l.Show()
	fmt.Println("----------------")

	// l.Reset()
	// fmt.Println(l)

	fmt.Printf("l.Len(): %v\n", l.Len())

	index := 3
	data, err := l.FindValueByIndex(index)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("第%v个位置的值为%v\n", index, data)
	}

	data = 7
	i, err := l.FindIndexByValue(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("值%v的位置为%v\n", data, i)
	}

}
