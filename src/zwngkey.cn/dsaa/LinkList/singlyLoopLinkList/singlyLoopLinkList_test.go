/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-27 14:12:55
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-27 21:16:16
 * @Description:
	go实现单向循环链表
*/
package singlylooplinklist

import (
	"fmt"
	"testing"
)

type PNode = *Node

// 定义结点
type Node struct {
	data any
	next PNode
}

// 定义单向循环链表
type List struct {
	head   PNode
	tail   PNode
	length int
}

func NewNode(data any) PNode {
	return &Node{data: data}
}

func (l *List) Init() *List {
	l.head.next = l.head
	l.tail = l.head
	l.length = 0
	return l
}

func NewList() *List {
	l := &List{head: NewNode(nil), tail: NewNode(nil)}
	l.Init()
	return l
}

func (l *List) IsEmpty() bool {
	return l.length == 0
}

func (l *List) Reset() {
	l.Init()
}

func (l *List) Frist() PNode {
	return l.head.next
}

func (l *List) Last() PNode {
	return l.tail
}

func (l *List) Len() int {
	return l.length
}

func (l *List) FindValueByIndex(i int) (data any, err error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("空链表")
	}
	if i <= 0 || i > l.length {
		return nil, fmt.Errorf("i 非法")
	}
	cur := l.Frist()
	for j := 1; j != i; j++ {
		cur = cur.next
	}
	return cur.data, nil
}

func (l *List) FindIndexByValue(data any) (i int, err error) {
	if l.IsEmpty() {
		return -1, fmt.Errorf("空链表")
	}
	cur := l.Frist()
	for j := 1; cur != l.head; j++ {
		if data == cur.data {
			return j, nil
		}
		cur = cur.next
	}
	return -1, fmt.Errorf("不存在")
}

func Test3(t *testing.T) {
	l := createList()
	l.Show()
	// i, err := l.FindIndexByValue(3)
	// i, err := l.FindIndexByValue(7)
	i, err := l.FindIndexByValue(19)
	fmt.Println(i, err)
}

func (l *List) InsertBeforeNode(newNode PNode) PNode {
	defer func() {
		l.length++
	}()
	if l.IsEmpty() {
		l.head.next = newNode
		newNode.next = l.head
		l.tail = newNode
		return newNode
	}
	newNode.next = l.Frist()
	l.head.next = newNode
	return newNode
}

func (l *List) InsertBeforeValue(data any) PNode {
	return l.InsertBeforeNode(NewNode(data))
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

func (l *List) InsertAfterNode(newNode PNode) PNode {
	defer func() {
		l.length++
	}()
	if l.IsEmpty() {
		l.head.next = newNode
		newNode.next = l.head
		l.tail = newNode
		return newNode
	}
	l.tail.next = newNode
	newNode.next = l.head
	l.tail = newNode
	return newNode
}

func (l *List) InsertAfterValue(data any) PNode {
	return l.InsertAfterNode(NewNode(data))
}

// 在第i个结点后插入新节点
func (l *List) InsertNodeByIndex(newNode PNode, i int) PNode {
	if i <= 0 {
		return l.InsertBeforeNode(newNode)
	}
	if i > l.length {
		return l.InsertAfterNode(newNode)
	}
	cur := l.Frist()
	for j := 1; j != i; j++ {
		cur = cur.next
	}
	newNode.next = cur.next
	cur.next = newNode

	if l.tail.next != l.head {
		l.tail = newNode
	}
	return newNode
}

func (l *List) InsertValueByIndex(data any, i int) PNode {
	return l.InsertNodeByIndex(NewNode(data), i)
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

func (l *List) DeleteBefore() (PNode, error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("空链表")
	}
	oldNode := l.Frist()
	l.head.next = oldNode.next
	l.length--
	if l.IsEmpty() {
		l.tail = l.head
	}
	return oldNode, nil
}

func Test6(t *testing.T) {
	l := createList()
	l.DeleteBefore()
	l.Show()
	l.DeleteBefore()
	l.Show()
}

func (l *List) DeleteBack() (PNode, error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("空链表")
	}
	if l.length == 1 {
		return l.DeleteBefore()
	}
	cur := l.Frist()
	for cur.next.next != l.head {
		cur = cur.next
	}
	oldNode := cur.next
	cur.next = l.head
	l.tail = cur
	l.length--
	return oldNode, nil
}

func Test7(t *testing.T) {
	l := createList()
	l.DeleteBack()
	l.Show()
	l.DeleteBack()
	l.Show()
}

func (l *List) Reverse() {
	if l.IsEmpty() || l.length == 1 {
		return
	}
	cur := l.Frist()
	pre := l.head
	for cur != l.head {
		// temp := cur.next
		// cur.next = pre
		// pre = cur
		// cur = temp
		cur.next, pre, cur = pre, cur, cur.next
	}
	l.tail = l.Frist()
	l.head.next = pre
}
func Test8(t *testing.T) {
	l := createList()
	l.Show()
	fmt.Println("-------")
	l.Reverse()
	l.Show()
	fmt.Println(l.tail.data)
}

func (l *List) Join(l1 *List) *List {
	l.tail.next = l1.Frist()
	l.tail = l1.tail
	l.tail.next = l.head
	return l
}
func Test9(t *testing.T) {
	l1 := createList()
	l2 := createList()
	l1.Join(l2)
	l1.Show()
	fmt.Println(l1.tail.data)
}

func (l *List) Show() {
	if l.length == 0 {
		fmt.Println("空链表")
		return
	}
	current := l.head.next
	i := 1
	for current != l.head {
		fmt.Printf("第%d的节点是%d\n", i, current.data)
		current = current.next
		i++
	}
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

func Test(t *testing.T) {
	l := createList()
	l.Show()
	fmt.Println(l.tail.data)
}

func createList() *List {
	l := NewList()
	l.InsertAfterNode(NewNode(1))
	l.InsertAfterNode(NewNode(2))
	l.InsertAfterNode(NewNode(3))
	l.InsertAfterNode(NewNode(4))
	l.InsertAfterNode(NewNode(7))
	return l
}
