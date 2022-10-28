/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 22:40:24
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-29 01:07:11
 * @Description:
	Josephu 问题
		设编号为 1，2，… n 的 n 个人围坐一圈，约定编号为 k（1<=k<=n）的人从 1开始报数，数到 m 的那个人出列，
		它的下一位又从 1 开始报数，数到 m 的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列。

	分析:
		用一个不带头结点的循环链表来处理 Josephu 问题：先构成一个有 n 个结点的单循环链表，然后由 k 结点起从 1 开始计数，计到 m 时，
			对应结点从链表中删除，然后再从被删除结点的下一个结点又从 1 开始计数，直到最后一个结点从链表中删除算法结束。

	注意:当只剩一个结点时,node.next = node;
*/
package problem

import (
	"fmt"
	"testing"
)

type (
	person struct {
		no   int
		next *person
	}
	PersonList struct {
		head *person
	}
)

func New() *PersonList {
	return &PersonList{}
}

func (p *PersonList) addPerson(num int) {
	if num < 1 {
		return
	}
	var cur *person
	for i := 1; i <= num; i++ {
		person := &person{no: i}
		if i == 1 {
			p.head = person
			p.head.next = p.head
			cur = p.head
		}
		cur.next = person
		person.next = p.head
		cur = person
	}
}

func (p *PersonList) Print() {
	if p.head == nil {
		fmt.Println("没有人")
		return
	}
	cur := p.head
	for ; cur.next != p.head; cur = cur.next {
		fmt.Printf("编号:%v\n", cur.no)
	}
	fmt.Printf("编号:%v\n", cur.no)
}

func Josephu(n, k, m int) {

	if n <= 1 || k > n || k < 0 || m < 0 {
		fmt.Println("参数非法")
		return
	}
	l := New()
	l.addPerson(n)

	// 初始位置--链尾
	helper := l.head
	for helper.next != l.head {
		helper = helper.next
	}

	// 找到第k-1个人
	for i := 0; i < k-1; i++ {
		helper = helper.next
	}
	flag := true
	for flag {
		if l.head == l.head.next {
			flag = false
		}
		//找到第m-1个人
		for i := 0; i < m-1; i++ {
			helper = helper.next
		}
		fmt.Println(helper.next.no)
		if helper.next == l.head {
			l.head = l.head.next
		}
		helper.next = helper.next.next
	}
}

func Josephu1(n, k, m int) {

	if n <= 1 || k > n || k < 0 || m < 0 {
		fmt.Println("参数非法")
		return
	}
	l := New()
	l.addPerson(n)

	// 初始位置--链尾
	helper := l.head
	for helper.next != l.head {
		helper = helper.next
	}

	// 找到第k-1个人
	for i := 0; i < k-1; i++ {
		helper = helper.next
		l.head = l.head.next
	}

	flag := true
	for flag {
		if l.head == helper {
			flag = false
		}
		//找到第m-1个人
		for i := 0; i < m-1; i++ {
			helper = helper.next
			l.head = l.head.next
		}
		fmt.Println(helper.next.no)
		l.head = l.head.next
		helper.next = l.head
	}
}

func Test(t *testing.T) {
	// Josephu(41, 1, 3)
	Josephu1(41, 1, 3)
}
