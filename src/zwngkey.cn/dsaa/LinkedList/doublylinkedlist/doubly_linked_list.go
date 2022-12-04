/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 18:23:43
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-01 12:05:16
 * @Description:
	go实现双向链表
*/
package dll

import (
	"fmt"
)

type (
	Node struct {
		preNode, nextNode *Node
		data              any
	}

	DList struct {
		header *Node //这是首元结点
		length int
	}
)

func newNode(data any) *Node {
	return &Node{nil, nil, data}
}

func NewList() *DList {
	return &DList{newNode(nil), 0}
}

func (d *DList) First() any {
	return d.header.data
}

func (d *DList) FirstNode() *Node {
	return d.header.nextNode
}

func (d *DList) Len() int {
	return d.length
}

//向链表头增加一个结点
func (d *DList) InsertBefore(data any) {
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
func (d *DList) InsertBack(data any) {
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
func (d *DList) InsertByIndex(i int, data any) {
	defer func() {
		d.length++
	}()
	if i > d.length {
		d.InsertBack(data)
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
func (d *DList) DeleteByIndex(i int) any {
	if i > d.length || i <= 0 {
		fmt.Printf("不存在第%d个结点\n", i)
		return nil
	}

	defer func() {
		d.length--
	}()

	if i == 1 {
		data := d.First()
		d.header = d.header.nextNode
		return data
	}

	cur := d.header
	for j := 1; j != i; j++ {
		cur = cur.nextNode
	}
	data := cur.data
	if cur.nextNode == nil {
		cur.preNode.nextNode = cur.nextNode
		return data
	}
	cur.preNode.nextNode = cur.nextNode
	cur.nextNode.preNode = cur.preNode
	return data
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
func (l DList) PrintList() {
	if l.length == 0 {
		fmt.Println("无数据")
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

func (n *Node) Next() *Node {
	return n.nextNode
}
