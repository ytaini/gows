/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-26 13:27:54
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-26 21:18:53
 * @Description:
	go实现双向循环链表及基本操作.

*/
package dll

import (
	"errors"
	"fmt"
	"testing"
)

// 定义链表中的节点
type Node struct {
	next, prev *Node
	data       any
	list       *LinkList
}

// 定义链表
type LinkList struct {
	//root.pre和root.next分别是链表的最后一个节点和第一个节点
	root   *Node //根节点，链表的起点，不存值。
	length int   // 当前链表长度，不包含根节点
}

// 初始化链表/清空链表
func (l *LinkList) Init() *LinkList {
	l.root.next = l.root
	l.root.prev = l.root
	l.length = 0
	return l
}

// 创建节点
func newNode(data any) *Node {
	return &Node{data: data}
}

// 创建链表且初始化
func NewLinkList() *LinkList {
	l := &LinkList{newNode(nil), 0}
	l.Init()
	return l
}

// 获取长度
func (l *LinkList) GetLen() int {
	return l.length
}

// 返回链表首节点
// 首节点就是根节点指向的后继，注意链表为空的情况处理。
func (l *LinkList) Front() *Node {
	if l.IsEmpty() {
		return nil
	}
	return l.root.next
}

// 返回链表末节点方法
func (l *LinkList) Back() *Node {
	if l.IsEmpty() {
		return nil
	}
	return l.root.prev
}

// 向链表特定元素后面插入新元素
func (l *LinkList) insertNode(newNode, atNode *Node) *Node {
	newNode.prev = atNode
	newNode.next = atNode.next
	atNode.next.prev = newNode
	atNode.next = newNode
	newNode.list = l
	l.length++
	return newNode
}

// 直接向链表插入一个值方法
func (l *LinkList) insertValue(data any, atNode *Node) *Node {
	return l.insertNode(newNode(data), atNode)
}

// 删除链表指定节点
func (l *LinkList) removeNode(node *Node) *Node {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
	node.list = nil
	l.length--
	return node
}

// 删除节点并返回元素值方法
func (l *LinkList) Remove(n *Node) (val any, err error) {
	if n.list != l {
		return nil, errors.New("删除节点失败")
	}
	l.removeNode(n)
	return n.data, nil
}

// 将一个节点n1移到另一节点n2后面方法
func (l *LinkList) move(n1, n2 *Node) *Node {
	if n1 == n2 {
		return n1
	}

	n1.prev.next = n1.next
	n1.next.prev = n1.prev

	n1.prev = n2
	n1.next = n2.next
	n2.next = n1
	n1.next.prev = n1

	return n1
}

// 插入一个值到链表头
// 链表头即首节点，也就是插入链表的root后面一个节点。
func (l *LinkList) PushFront(data any) *Node {
	return l.insertValue(data, l.root)
}

// 插入一个值到链表末尾
// 末节点也就是root的前驱。这里使用环形双向链表操作比较方便。
func (l *LinkList) PushBack(data any) *Node {
	return l.insertValue(data, l.root.prev)
}

// 插入一个值到某个节点前面
// 插入值到元素n到前面，也就是插入到n.pre的后面。
func (l *LinkList) InsertBefore(data any, n *Node) *Node {
	if n.list != l {
		return nil
	}
	return l.insertValue(data, n.prev)
}

//插入一个值到某节点后面
func (l *LinkList) InsertAfter(data any, n *Node) *Node {
	if n.list != l {
		return nil
	}
	return l.insertValue(data, n)
}

// 将一个节点移到链表首位置
func (l *LinkList) MoveToFront(node *Node) {
	if node.list != l || l.root.next == node {
		return
	}
	l.move(node, l.root)
}

// 将一个节点移动到链表末尾
func (l *LinkList) MoveToBack(node *Node) {
	if node.list != l || l.root.prev == node {
		return
	}
	l.move(node, l.root.prev)
}

// 将一个节点移到另一节点前面
// 先判断特殊情况，然后等同于移到node2前面一个元素的后面。
func (l *LinkList) MoveBefore(node1, node2 *Node) {
	if node1.list != l || node2.list != l || node1 == node2 {
		return
	}
	l.move(node1, node2.prev)
}

// 移动一个节点到另一个节点的后面
func (l *LinkList) MoveAfter(node1, node2 *Node) {
	if node1.list != l || node2.list != l || node1 == node2 {
		return
	}
	l.move(node1, node2)
}

// 按值查找指定结点
func (l *LinkList) FindByValue(data any) *Node {
	cur := l.Front()

	for cur != l.root {
		if data == cur.data {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// 按值删除指定结点
func (l *LinkList) DeleteByValue(data any) *Node {
	cur := l.Front()

	for cur != l.root {
		if cur.data == data {
			return l.removeNode(cur)
		}
		cur = cur.next
	}
	return nil
}

// 把值为x的元素的值修改为y
func (l *LinkList) ModifyValue(data any) *Node {
	node := l.FindByValue(data)
	if node != nil {
		node.data = data
	}
	return node
}

// 反转链表
// 交换各个结点的prev,next
func (l *LinkList) Reverse() {
	if l.length < 2 {
		return
	}
	firstNode := l.Front()
	lastNode := l.Back()

	var tempNext *Node
	cur := l.Front()
	//根节点与其他节点分开交换.
	for cur != l.root {
		tempNext = cur.next
		cur.next = cur.prev
		cur.prev = tempNext
		cur = tempNext
	}

	l.root.next = lastNode
	l.root.prev = firstNode
}

// 判断链表是否为空
func (l *LinkList) IsEmpty() bool {
	return l.length == 0
}

// 显示链表
func (l *LinkList) ShowLinkList() {
	if l.IsEmpty() {
		fmt.Println("空链表")
		return
	}
	cur := l.Front()
	i := 1
	for cur != l.root {
		fmt.Printf("第%v个结点的值为: %v\n", i, cur.data)
		cur = cur.next
		i++
	}
}

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
