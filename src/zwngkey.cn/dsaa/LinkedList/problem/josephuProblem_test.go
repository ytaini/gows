/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 22:40:24
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-29 10:58:59
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
	// 定义结点
	person struct {
		no   int
		next *person
	}
	//用一个不带头结点的循环链表来处理
	//定义单向循环链表
	PersonList struct {
		head *person
	}
)

func New() *PersonList {
	return &PersonList{}
}

// 生成num个结点的不带头的单向循环链表
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

/*
	1.创建n个结点,不带头结点的链表
	2.创建一个辅助指针(变量)helper,初始指向最后一个结点
	3.找到第k个人,让helper指向第k-1个人
	4.判断当前链表是否只有一个结点.如果是flag置为false.
	5.再找到第m个人,让helper指向第m-1个人.
		判断头指针是否指向第m个人
			如果是,头指针指向下一个节点.删除第m个人
			如果不是,删除第m个人.
	6.重复第4-5步.
*/
func Josephu(n, k, m int) {

	if n <= 1 || k > n || k < 0 || m < 0 {
		fmt.Println("参数非法")
		return
	}
	l := New()
	l.addPerson(n)

	head := l.head

	// 初始位置--链尾
	helper := head
	for helper.next != head {
		helper = helper.next
	}

	// 找到第k-1个人
	for i := 0; i < k-1; i++ {
		helper = helper.next
	}
	flag := true
	for flag {
		if l.head == head.next {
			flag = false
		}
		//找到第m-1个人
		for i := 0; i < m-1; i++ {
			helper = helper.next
		}
		fmt.Println(helper.next.no) //打印第m个人的编号
		if helper.next == head {
			head = head.next
		}
		helper.next = helper.next.next //删除第m个人
	}
}

/*
*表示核心操作.

	定义头指针head:指向第1人
	*定义辅助指针(变量)helper:指向最后一个人.
	1.找到第k个人,head指向它,helper指向第k-1个人
	*2.判断头指针是否等于heleper (helper.next 始终等于head,除了只剩一个结点时)
	3.找到第m个人,head指向它,helper指向第m-1个人
	*4.删除第m个人,并打印其编号.
		删除操作: 先让head指向第m+1个人,然后helper.next=head
	5.重复2-4步.
*/
func Josephu1(n, k, m int) {

	if n <= 1 || k > n || k < 0 || m < 0 {
		fmt.Println("参数非法")
		return
	}
	// 创建不带头结点的单向循环链表
	l := New()

	// 添加n个结点
	l.addPerson(n)

	// 头指针
	head := l.head

	// 辅助指针--使其指向最后一个结点.
	helper := head
	for helper.next != head {
		helper = helper.next
	}

	for i := 0; i < k-1; i++ {
		helper = helper.next //helper指向第k-1个人
		head = head.next     //head指向第k个人
	}

	flag := true
	for flag {
		// 判断是否只有一个结点了
		if head == helper {
			flag = false
		}

		for i := 0; i < m-1; i++ {
			helper = helper.next //helper指向第m-1个人
			head = head.next     //head指向第m个人
		}
		fmt.Println(helper.next.no) //打印第m个人的编号
		head = head.next            //head 指向第m+1个人
		helper.next = head          //删除第m个人
	}
}

func Test(t *testing.T) {
	// Josephu(41, 1, 3)
	Josephu1(41, 1, 3)
}
