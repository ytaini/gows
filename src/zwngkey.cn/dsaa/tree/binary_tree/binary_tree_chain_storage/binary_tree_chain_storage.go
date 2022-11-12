/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-12 01:30:55
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-13 03:49:04
 * @Description:
	二叉树的二叉链表存储表示
*/
package binarytree

import (
	"fmt"

	"zwngkey.cn/dsaa/queue"
	"zwngkey.cn/dsaa/stack"
)

// 二叉树的二叉链表存储表示
// 二叉链表，由于缺乏父链的指引，在找回父节点时需要重新扫描树得知父节点的节点地址。
type BiTNode struct {
	data           any
	lchild, rchild *BiTNode //左右孩子指针
}

type BiTree struct {
	root *BiTNode
}

// 前序遍历(递归实现)
func (t *BiTree) PreOrderTreverseRecus() {
	var preOrder func(node *BiTNode)
	preOrder = func(node *BiTNode) {
		if node != nil {
			fmt.Printf("%v\n", node.data)
			preOrder(node.lchild)
			preOrder(node.rchild)
		}
	}
	preOrder(t.root)
}

// 前序遍历(非递归实现)
func (t *BiTree) PreOrderTreverseNonRecur() {
	s := stack.New[*BiTNode]()
	p := t.root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			fmt.Printf("%v\n", p.data)
			s.Push(p)
			p = p.lchild
		} else {
			q, _ := s.Pop()
			p = q.rchild
		}
	}
}

// 中序遍历(递归实现)
func (t *BiTree) InOrderTreverseRecur() {
	var inOrder func(node *BiTNode)
	inOrder = func(node *BiTNode) {
		if node != nil {
			inOrder(node.lchild)
			fmt.Printf("%v\n", node.data)
			inOrder(node.rchild)
		}
	}
	inOrder(t.root)
}

// 中序遍历(非递归实现)
func (t *BiTree) InOrderTreverseNonRecur() {
	s := stack.New[*BiTNode]()
	p := t.root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			s.Push(p)
			p = p.lchild
		} else {
			q, _ := s.Pop()
			fmt.Printf("%v\n", q.data)
			p = q.rchild
		}
	}
}

// 后序遍历(递归实现)
func (t *BiTree) PostOrderTreverseRecur() {
	var postOrder func(node *BiTNode)
	postOrder = func(node *BiTNode) {
		if node != nil {
			postOrder(node.lchild)
			postOrder(node.rchild)
			fmt.Printf("%v\n", node.data)
		}
	}
	postOrder(t.root)
}

// 后序遍历1(非递归实现)
// 算法: 看图.
func (t *BiTree) PostOrderTreverseNonRecur1() {
	s := stack.New[*BiTNode]() //用于记录根和子根结点
	cur := t.root              //临时变量记录当前访问的结点.
	var pre *BiTNode           //临时变量记录上一个访问到的结点,

	for cur != nil || !s.IsEmpty() {
		if cur != nil { //如果当前结点不为nil,就一直找它的左孩子.
			s.Push(cur)
			cur = cur.lchild
		} else { //如果当前结点为nil,说明其父结点没有左节点或已经访问完了.
			cur, _ = s.Peek() //拿出其父节点.
			if cur.rchild != nil /*判断其父节点有无右孩子*/ && cur.rchild != pre /*  是否访问过其右孩子 */ {
				cur = cur.rchild //接着判断其右孩子是否有左孩子.
			} else { //如果没有,那么直接访问其父节点.
				v, _ := s.Pop()
				visit(v)  //对p进行访问.
				pre = v   //记录访问过的结点.
				cur = nil //将cur置空,进入下一层循环,直到栈内无元素.
			}
		}
	}
}

// 访问结点函数.
func visit(node *BiTNode) {
	fmt.Println(node.data)
}

// 后序遍历2(非递归实现)
// 算法: 看图
func (t *BiTree) PostOrderTreverseNonRecur2() {
	if t.root == nil {
		fmt.Println("空树")
		return
	}
	s := stack.New[*BiTNode]()
	var cur *BiTNode //临时变量记录当前访问的结点.
	var pre *BiTNode //临时变量记录上一个访问到的结点,
	s.Push(t.root)   //先将 根节点入栈.

	for !s.IsEmpty() {
		cur, _ = s.Peek()
		if (cur.lchild == nil && cur.rchild == nil) || (pre != nil && (pre == cur.lchild || pre == cur.rchild)) {
			visit(cur)
			s.Pop()
			pre = cur
		} else {
			if cur.rchild != nil {
				s.Push(cur.rchild)
			}
			if cur.lchild != nil {
				s.Push(cur.lchild)
			}
		}
	}
}

// 层次遍历BFS
func (t *BiTree) LevelOrderTreverse() {
	if t.root == nil {
		fmt.Println("空树")
		return
	}
	q := queue.New[*BiTNode]()
	q.Push(t.root)

	for !q.IsEmpty() {
		n, _ := q.Pop()
		fmt.Println(n.data)
		if n.lchild != nil {
			q.Push(n.lchild)
		}
		if n.rchild != nil {
			q.Push(n.rchild)
		}
	}
}

// 递归实现前序查找
// 算法: 看图
func (t *BiTree) PreOrderSearchRecur(data any) bool {
	var preOrderSearch func(p *BiTNode, data any) bool
	preOrderSearch = func(p *BiTNode, data any) bool {
		if p != nil {
			if fmt.Println(1); p.data == data {
				return true
			}
			if is := preOrderSearch(p.lchild, data); is {
				return is
			}
			if is := preOrderSearch(p.rchild, data); is {
				return is
			}
			// return preOrderSearch(p.lchild, data) || preOrderSearch(p.rchild, data)
		}
		return false
	}
	return preOrderSearch(t.root, data)
}

// 非递归实现前序查找
func (t *BiTree) PreOrderSearchNonRecur(data any) bool {
	s := stack.New[*BiTNode]()
	p := t.root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			if fmt.Println(1); p.data == data {
				return true
			}
			s.Push(p)
			p = p.lchild
		} else {
			q, _ := s.Pop()
			p = q.rchild
		}
	}
	return false
}

// 递归实现中序查找
// 算法: 看图
func (t *BiTree) InOrderSearchRecur(data any) bool {
	var inOrderSearch func(p *BiTNode, data any) bool
	inOrderSearch = func(p *BiTNode, data any) bool {
		if p != nil {
			if is := inOrderSearch(p.lchild, data); is {
				return is
			}
			if fmt.Println(1); p.data == data {
				return true
			}
			if is := inOrderSearch(p.rchild, data); is {
				return is
			}
		}
		return false
	}
	return inOrderSearch(t.root, data)
}

// 非递归实现中序查找
func (t *BiTree) InOrderSearchNonRecur(data any) bool {
	s := stack.New[*BiTNode]()
	p := t.root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			s.Push(p)
			p = p.lchild
		} else {
			q, _ := s.Pop()
			if fmt.Println(1); q.data == data {
				return true
			}
			p = q.rchild
		}
	}
	return false
}

// 递归实现后序查找
// 算法: 看图
func (t *BiTree) PostOrderSearchRecur(data any) bool {
	var postOrderSearch func(p *BiTNode, data any) bool
	postOrderSearch = func(p *BiTNode, data any) bool {
		if p != nil {
			if is := postOrderSearch(p.lchild, data); is {
				return is
			}
			if is := postOrderSearch(p.rchild, data); is {
				return is
			}
			if fmt.Println(1); p.data == data {
				return true
			}
		}
		return false
	}
	return postOrderSearch(t.root, data)
}

// 非递归实现后序查找
func (t *BiTree) PostOrderSearchNonRecur(data any) bool {
	s := stack.New[*BiTNode]()
	var cur *BiTNode
	var pre *BiTNode
	s.Push(t.root)
	for !s.IsEmpty() {
		cur, _ = s.Peek()
		if (cur.lchild == nil && cur.rchild == nil) || (pre != nil && (pre == cur.lchild || pre == cur.rchild)) {
			if fmt.Println(1); cur.data == data {
				return true
			}
			s.Pop()
			pre = cur
		} else {
			if cur.rchild != nil {
				s.Push(cur.rchild)
			}
			if cur.lchild != nil {
				s.Push(cur.lchild)
			}
		}
	}
	return false
}

// 前序递归实现删除节点
// 	如果删除的节点是叶子节点，则删除该节点
// 	如果删除的节点是非叶子节点，则删除该子树
func (t *BiTree) PreOrderRecurDelete1(data any) (bool, *BiTNode) {
	if t.root == nil {
		return false, nil
	}
	if t.root.data == data {
		node := t.root
		t.root = nil
		return true, node
	}
	var del func(b *BiTNode, data any) (bool, *BiTNode)

	del = func(b *BiTNode, data any) (bool, *BiTNode) {
		if b != nil {
			if b.lchild != nil && b.lchild.data == data {
				node := b.lchild
				b.lchild = nil
				return true, node
			}
			if is, node := del(b.lchild, data); is {
				return is, node
			}
			if b.rchild != nil && b.rchild.data == data {
				node := b.rchild
				b.rchild = nil
				return true, node
			}
			if is, node := del(b.rchild, data); is {
				return is, node
			}
		}
		return false, nil
	}

	return del(t.root, data)
}

// 前序递归实现删除节点
func (t *BiTree) PreOrderRecurDelete2(data any) (is bool, bt *BiTNode) {
	var del func(b *BiTNode, data any) *BiTNode
	del = func(b *BiTNode, data any) *BiTNode {
		if b != nil {
			if b.data == data {
				bt = b
				is = true
				return nil
			}
			b.lchild = del(b.lchild, data)
			b.rchild = del(b.rchild, data)
		}
		return b
	}
	t.root = del(t.root, data)
	return
}

// 前序递归实现删除节点
// 如果要删除的节点是非叶子节点，现在我们不希望将该非叶子节点为根节点的子树删除，需要指定规则如下:
// - 如果该非叶子节点A只有一个子节点B，则子节点B替代节点A
// - 如果该非叶子节点A有左子节点B和右子节点C，则让左子节点B替代节点A
func (t *BiTree) PreOrderRecurDelete3(data any) (is bool, bt *BiTNode) {
	var del func(b *BiTNode, data any) *BiTNode
	del = func(b *BiTNode, data any) *BiTNode {
		if b != nil {
			if b.data == data {
				bt = b
				is = true
				if b.lchild != nil {
					return b.lchild
				} else {
					return b.rchild
				}
			}
			b.lchild = del(b.lchild, data)
			b.rchild = del(b.rchild, data)
		}
		return b
	}
	t.root = del(t.root, data)
	return
}
