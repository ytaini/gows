/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-27 08:04:44
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-27 17:55:57
 * @Description:
	go实现单链表
*/
package sll

import (
	"errors"
	"fmt"
	"testing"
)

// 结点指针
type PNode = *LNode

// 结点结构体
type LNode struct {
	data any
	next PNode
}

// 链表类型
type LLinkList = PNode

// 创建一个单链表
func NewLLinkList() LLinkList {
	return new(LNode)
}

func NewLNode(data any) PNode {
	return &LNode{data: data}
}

// 判空
func (l LLinkList) IsEmpty() bool {
	return l.next == nil
}

// 销毁
func (l LLinkList) Destory() {
	l.next = nil
	// l = nil //在方法体内对属主参数的直接部分的修改将不会反映到方法体外
}

// 重置
func (l LLinkList) Reset() {
	l.next = nil
}

// 返回首元结点
func (l LLinkList) Frist() PNode {
	return l.next
}

// 返回尾节点
func (l LLinkList) Last() PNode {
	if l.IsEmpty() {
		return nil
	}
	cur := l.Frist()
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

// 返回长度
func (l LLinkList) Len() int {
	len := 0
	cur := l.Frist()
	for cur != nil {
		len++
		cur = cur.next
	}
	return len
}

// 通过位置查找值
func (l LLinkList) FindValueByIndex(i int) (data any, err error) {
	if l.IsEmpty() {
		return nil, errors.New("链表为空")
	}
	if i > l.Len() {
		return nil, fmt.Errorf("第%v个元素不存在", i)
	}
	cur := l.Frist()
	for j := 1; j < i; j++ {
		cur = cur.next
	}
	return cur.data, nil
}

// 通过值查询位置
func (l LLinkList) FindIndexByValue(data any) (i int, err error) {
	if l.IsEmpty() {
		return -1, errors.New("链表为空")
	}
	cur := l.Frist()
	for j := 1; cur != nil; j++ {
		if cur.data == data {
			return j, nil
		}
		cur = cur.next
	}
	return -1, errors.New("链表中不存在该值")
}

// 头插结点
func (l LLinkList) InsertBeforeNode(newNode PNode) PNode {
	if l.IsEmpty() {
		l.next = newNode
		return newNode
	}
	newNode.next = l.next
	l.next = newNode
	return newNode
}

// 头插值
func (l LLinkList) InsertBeforeValue(data any) PNode {
	return l.InsertBeforeNode(NewLNode(data))
}

// 尾插结点
func (l LLinkList) InsertBackNode(newNode PNode) PNode {
	if l.IsEmpty() {
		return l.InsertBeforeNode(newNode)
	}
	cur := l.Frist()
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
	return newNode
}

// 尾插值
func (l LLinkList) InsertBackValue(data any) PNode {
	return l.InsertBackNode(NewLNode(data))
}

// 在第i个结点后插入新节点
func (l LLinkList) InsertNodeByIndex(i int, newNode PNode) PNode {
	if l.IsEmpty() || i == 1 {
		return l.InsertBeforeNode(newNode)
	}
	if i > l.Len() {
		return l.InsertBackNode(newNode)
	}

	cur := l.Frist()

	// 找到第i个结点前一个结点
	for j := 1; j != i; j++ {
		cur = cur.next
	}
	newNode.next = cur.next
	cur.next = newNode

	return newNode
}

// 在第i个后插入值data
func (l LLinkList) InsertValueByIndex(i int, data any) PNode {
	return l.InsertNodeByIndex(i, NewLNode(data))
}

// 头删
func (l LLinkList) DeleteBefore() (oldNode PNode, err error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("空链表,删除失败")
	}
	cur := l.Frist()
	l.next = cur.next
	return cur, nil
}

// 尾删
func (l LLinkList) DeleteBack() (oldNode PNode, err error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("空链表,删除失败")
	}
	if l.Len() == 1 {
		return l.DeleteBefore()
	}
	cur := l.Frist()
	for cur.next.next != nil {
		cur = cur.next
	}
	oldNode = cur.next
	cur.next = nil
	return oldNode, nil
}

// 删除第i个结点
func (l LLinkList) DeleteByIndex(i int) (oldNode PNode, err error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("空链表,删除失败")
	}
	if i <= 0 {
		return nil, fmt.Errorf("删除失败,i值非法")
	}
	if i == 1 {
		return l.DeleteBefore()
	}
	if i > l.Len() {
		return l.DeleteBack()
	}
	//寻找第i-1个结点
	cur := l.Frist()
	for j := 1; j != i-1; j++ {
		cur = cur.next
	}
	oldNode = cur.next
	cur.next = cur.next.next
	return oldNode, nil
}

// 反转单链表
func (l LLinkList) Reverse() {
	if l.IsEmpty() || l.Len() == 1 {
		return
	}
	cur := l.Frist() //当前节点
	var pre PNode    //前一个节点
	for cur != nil {
		// temp := cur.next
		// cur.next = pre
		// pre = cur
		// cur = temp

		/*
			实际工作中我们还是尽量不要写出像👇🏻这样复杂、难以让人理解的语句。
			必要的话，拆分成多行就好了，还可以增加些代码量（如果你的公司是以代码量为评价绩效指标之一的），
			得饶人处且饶人啊，烧脑的语句还是尽量避免为好。
		*/
		//在同一步操作中，改变各个变量对应的值，可以省去中间变量
		cur.next, pre, cur = pre, cur, cur.next
	}
	l.next = pre
}

// 显示链表
func (l LLinkList) Show() {
	if l.IsEmpty() {
		fmt.Println("空链表")
		return
	}
	current := l.Frist()
	i := 1
	for current != nil {
		fmt.Printf("第%d的节点是%d\n", i, current.data)
		current = current.next
		i++
	}
}

func Test4(t *testing.T) {
	l := NewLLinkList()
	l.InsertBackValue(1)
	l.InsertBackValue(2)
	l.InsertBackValue(3)
	l.InsertBackValue(5)
	l.InsertBackValue(6)
	l.Show()
	l.Reverse()
	l.Show()
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
