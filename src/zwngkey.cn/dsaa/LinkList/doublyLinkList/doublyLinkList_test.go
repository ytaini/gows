/*
 * @Author: imzw
 * @Date: 2022-10-26 11:14:55
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 19:40:42
 * @Description:
	测试
*/
package dll

import (
	"fmt"
	"testing"
)

func Test2(t *testing.T) {
	l := NewList()
	l.InsertBack(11)
	l.InsertBack(22)
	l.InsertBefore(1)
	l.InsertBefore(2)
	l.InsertBefore(3)
	l.InsertBack(4)
	l.InsertBack(5)
	l.InsertBack(6)
	l.InsertByIndex(3, 10)
	l.PrintList()
	l.Reset()
	l.PrintList()
}

func Test1(t *testing.T) {
	l := NewList()
	l.InsertBack(11)
	l.InsertBack(22)
	l.InsertBefore(1)
	l.InsertBefore(2)
	l.InsertBefore(3)
	l.InsertBack(4)
	l.InsertBack(5)
	l.InsertBack(6)
	l.InsertByIndex(3, 10)

	l.PrintList()
	fmt.Println("--------")

	l.DeleteByIndex(5)
	l.DeleteByIndex(5)
	l.DeleteByIndex(6)
	l.DeleteByIndex(6)

	l.PrintList()
	fmt.Println("--------")

	fmt.Println(l.LookupNode(10))
	fmt.Println(l.LookupNode(7))
	fmt.Println(l.LookupNode(1))
	fmt.Println(l.length)
}

func Test(t *testing.T) {
	l := NewList()
	l.InsertBefore(1)
	l.InsertBefore(2)
	l.InsertBefore(3)
	l.PrintList()
	fmt.Println(l.header.nextNode.preNode.data)
	fmt.Println(l.header.nextNode.data)
	fmt.Println(l.header.nextNode.nextNode.data)
}
