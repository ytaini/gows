/*
 * @Author: imzw
 * @Date: 2022-10-26 09:59:45
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-27 09:27:30
 * @Description:
	go 实现单链表
	Node: 包含一个数据域，一个指针域（指向下一个节点）
	LList : 包含头指针 (指向第一个节点)，链表长度
	链表的特点：不能随机访问，只能根据链一个一个查找，查找的时间复杂度是 O (n)

*/
package sll

import (
	"fmt"
	"testing"
)

type (
	Node struct {
		data any
		next *Node
	}

	LList struct {
		header *Node
		length int
	}
)

func NewNode(v any) *Node {
	return &Node{v, nil}
}

func NewList() *LList {
	return &LList{NewNode(nil), 0}
}

//向链表头增加一个结点
func (l *LList) Add(data any) {
	newNode := NewNode(data)
	if l.length == 0 {
		l.header = newNode
	} else {
		newNode.next = l.header
		l.header = newNode //头指针指向新加的结点
	}

	l.length++
}

//向链表尾增加一个结点
func (l *LList) Append(data any) {
	newNode := NewNode(data)
	if l.length == 0 {
		l.header = newNode
	}
	cur := l.header
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
	l.length++
}

//往第i位插入一个节点
func (l *LList) Insert(i int, data any) {
	if i > l.length {
		l.Append(data)
		return
	}
	newNode := NewNode(data)
	cur := l.header
	for j := 1; j != i-1; j++ {
		cur = cur.next
	}
	after := cur.next
	cur.next = newNode
	newNode.next = after
	l.length++
}

//删除第i个节点
func (l *LList) Delete(i int) {
	if i == 1 {
		l.header = l.header.next
		return
	}
	cur := l.header
	for j := 1; j != i-1; j++ {
		cur = cur.next
	}
	after := cur.next.next
	cur.next = after
	l.length--
}

// 获取链表长度
func (l *LList) GetLength() int {
	return l.length
}

func lookupNode(node *Node, data any) bool {
	if node.data == data {
		return true
	}
	if node.next == nil {
		return false
	}
	return lookupNode(node.next, data)
}

// 查找节点是否存在
func (l *LList) LookupNode(data any) bool {
	return lookupNode(l.header, data)
}

//显示
func (l *LList) Scan() {
	if l.length == 0 {
		fmt.Println("空链表")
		return
	}
	current := l.header
	i := 1
	for current.next != nil {
		fmt.Printf("第%d的节点是%d\n", i, current.data)
		current = current.next
		i++
	}
	fmt.Printf("第%d的节点是%d\n", i, current.data)
}

func Test(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Append(5)
	l.Append(6)
	l.Append(7)
	l.Insert(3, 9)
	l.Delete(4)
	l.Scan()
	fmt.Println(l.LookupNode(1))
}
