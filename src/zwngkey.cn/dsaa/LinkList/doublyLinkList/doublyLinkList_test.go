/*
 * @Author: imzw
 * @Date: 2022-10-26 11:14:55
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-26 13:39:11
 * @Description:
	go实现双向链表
*/
package dll

import (
	"fmt"
	"testing"
)

type (
	Node struct {
		preNode, nextNode *Node
		data              any
	}

	DList struct {
		header *Node
		length int
	}
)

func newNode(data any) *Node {
	return &Node{nil, nil, data}
}

func NewList() *DList {
	return &DList{newNode(nil), 0}
}

//向链表头增加一个结点
func (d *DList) Add(data any) {
	defer func() {
		d.length++
	}()
	newNode := newNode(data)
	if d.length == 0 {
		d.header = newNode
		return
	}
	newNode.nextNode = d.header
	d.header.preNode = newNode
	d.header = newNode
}

func l_append(node *Node, data any) {
	if node.nextNode == nil {
		newNode := newNode(data)
		node.nextNode = newNode
		newNode.preNode = node
		return
	}
	l_append(node.nextNode, data)
}

// 向链表尾增加一个结点
func (d *DList) Append(data any) {
	defer func() {
		d.length++
	}()
	newNode := newNode(data)
	if d.length == 0 {
		d.header = newNode
		return
	}
	// cur := d.header
	// for cur.nextNode != nil {
	// 	cur = cur.nextNode
	// }
	// cur.nextNode = newNode
	// newNode.preNode = cur
	l_append(d.header, data)
}

// 向第i个位置增加一个结点
func (d *DList) Insert(i int, data any) {
	defer func() {
		d.length++
	}()
	if i > d.length {
		d.Append(data)
		return
	}

	cur := d.header

	for j := 1; j != i; j++ {
		cur = cur.nextNode
	}
	newNode := newNode(data)

	cur.preNode.nextNode = newNode
	newNode.preNode = cur.preNode
	newNode.nextNode = cur
	cur.preNode = newNode
}

//删除第i个节点
func (d *DList) Delete(i int) {
	if i > d.length || i <= 0 {
		fmt.Printf("不存在第%d个结点\n", i)
		return
	}

	defer func() {
		d.length--
	}()

	if i == 1 {
		d.header = d.header.nextNode
		return
	}

	cur := d.header
	for j := 1; j != i; j++ {
		cur = cur.nextNode
	}
	cur.preNode.nextNode = cur.nextNode
	cur.nextNode.preNode = cur.preNode
}

func lookupNode(node *Node, data any) bool {
	if node.data == data {
		return true
	}
	if node.nextNode == nil {
		return false
	}
	return lookupNode(node.nextNode, data)
}

// 查找节点是否存在
func (l *DList) LookupNode(data any) bool {
	return lookupNode(l.header, data)
}

//重置
func (l *DList) Reset() {
	l.header = nil
	l.length = 0
}

// 显示链表
func (l DList) ShowList() {
	if l.length == 0 {
		fmt.Println("空链表")
		return
	}
	current := l.header
	i := 1
	for current.nextNode != nil {
		fmt.Printf("第%v的节点是%v\n", i, current.data)
		current = current.nextNode
		i++
	}
	fmt.Printf("第%v的节点是%v\n", i, current.data)
}

func Test2(t *testing.T) {
	l := NewList()
	l.Append(11)
	l.Append(22)
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)
	l.Insert(3, 10)
	l.ShowList()
	l.Reset()
	l.ShowList()
}

func Test1(t *testing.T) {
	l := NewList()
	l.Append(11)
	l.Append(22)
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)
	l.Insert(3, 10)
	l.Delete(5)
	l.Delete(5)
	l.Delete(6)
	fmt.Println(l.LookupNode(10))
	fmt.Println(l.LookupNode(7))
	fmt.Println(l.LookupNode(1))
	fmt.Println(l.length)
	l.ShowList()
}

func Test(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.ShowList()
	fmt.Println(l.header.nextNode.preNode.data)
	fmt.Println(l.header.nextNode.data)
	fmt.Println(l.header.nextNode.nextNode.data)
}
