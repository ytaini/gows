package lc02

// AddTwoNumbers1 新建一个单向链表.
func AddTwoNumbers1(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

// AddTwoNumbers 自己写的.
// 在l2的基础上修改.
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	head = l2
	var pre *ListNode // 记录l2的父节点
	for l1 != nil && l2 != nil {
		v := l1.Val + l2.Val
		if v > 9 {
			l2.Val = v % 10
			if l2.Next == nil {
				l2.Next = &ListNode{Val: 1}
			} else {
				l2.Next.Val = l2.Next.Val + 1
			}
		} else {
			l2.Val = v
		}
		pre = l2
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 == nil { // l1位数较多
		pre.Next = l1
	}
	for ; l2 != nil; l2 = l2.Next { //l2位数较多
		if l2.Val > 9 {
			l2.Val = l2.Val % 10
			if l2.Next == nil {
				l2.Next = &ListNode{Val: 1}
			} else {
				l2.Next.Val = l2.Next.Val + 1
			}
		} else {
			break
		}
	}
	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}
