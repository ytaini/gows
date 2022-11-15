/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-14 02:00:37
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-14 06:47:49
 * @Description:
	二叉树的二叉线索存储表示：
		在线索链表上添加一个头结点，并令其lchild域的指针指向二叉树的根结点，其rchild域的指针指向中序遍历时访问的最后一个结点。
		令二叉树中序序列中的第一个结点的lchild域指针和最后一个结点的rchild域的指针均指向头结点，这样就创建了一个双向线索链表。
		二叉树常采用二叉链表方式存储。
*/
package threadbinarytree

import "fmt"

type TBiTNode struct {
	data           any
	ltag, rtag     int
	lchild, rchild *TBiTNode
}

const LINK = 0   //指针
const THREAD = 1 //线索

type TBiTree struct {
	head *TBiTNode
}

func newTNode(data any) *TBiTNode {
	return &TBiTNode{data: data}
}

func newTree() *TBiTree {
	h := newTNode(nil)
	h.ltag = LINK
	h.rtag = THREAD
	h.rchild = h
	h.lchild = h
	return &TBiTree{head: h}
}

// 通过数组构造一个二叉树
func CreateTBiTree(s []rune) *TBiTree {
	var create func(index int) *TBiTNode
	create = func(index int) *TBiTNode {
		var node *TBiTNode
		if index < len(s) {
			node = newTNode(s[index])
			node.lchild = create(index*2 + 1)
			node.rchild = create(index*2 + 2)
		}
		return node
	}
	tree := newTree()
	tree.head.lchild = create(0)
	return tree
}

// 通过中序遍历进行中序线索化
func (t *TBiTree) InThreading() {
	var inThreading func(n *TBiTNode)
	var pre *TBiTNode = t.head //记录当前访问过的结点.
	inThreading = func(n *TBiTNode) {
		if n != nil { //线索二叉树不为空
			inThreading(n.lchild) // 递归左子树线索化
			if n.lchild == nil {  //没有左孩子
				n.ltag = THREAD //左标志为线索
				n.lchild = pre  //左孩子指向前驱
			}
			if pre != nil && pre.rchild == nil { //前驱不为空且前驱没有右孩子.
				pre.rtag = THREAD //前驱的右标志为线索
				pre.rchild = n    //前驱的右孩子指向其后继
			}
			pre = n               //保持pre指向n的前驱.
			inThreading(n.rchild) // 递归右子树线索化
		}
	}
	inThreading(t.head.lchild)
	pre.rchild = t.head //最后一个结点的右指标指向头结点
	pre.rtag = THREAD   //最后一个结点的右标志为线索
	t.head.rchild = pre //头结点的右指标指向中序遍历的最后一个结点
}

// 中序遍历线索二叉树(头结点)的非递归算法
func (t *TBiTree) InOrderTraverse() {
	p := t.head.lchild //p指向头结点
	for p != t.head {
		for p.ltag == LINK { //由根节点一直找到二叉树的最左节点.
			p = p.lchild
		}
		fmt.Printf("%c\n", p.data)
		for p.rtag == THREAD && p.rchild != t.head {
			p = p.rchild
			fmt.Printf("%c\n", p.data)
		}
		p = p.rchild
	}
}

// 通过前序遍历进行前序线索化
func (t *TBiTree) PreThreading() {
	var preThreading func(n *TBiTNode)
	var pre *TBiTNode = t.head //全局变量
	preThreading = func(n *TBiTNode) {
		if n != nil {
			if n.lchild == nil { //n没有左孩子
				n.ltag = THREAD
				n.lchild = pre
			}
			if pre.rchild == nil { //n的前驱没有右孩子
				pre.rtag = THREAD
				pre.rchild = n
			}
			pre = n             //保持pre指向n的前驱.
			if n.ltag == LINK { //n有左孩子
				preThreading(n.lchild)
			}
			if n.rtag == LINK { //n有右孩子
				preThreading(n.rchild)
			}
		}
	}
	preThreading(t.head.lchild)
	pre.rchild = t.head
	pre.rtag = THREAD
	t.head.rchild = pre
}

// 前序遍历线索二叉树(头结点)的非递归算法
func (t *TBiTree) PreOrderTraverse() {
	p := t.head.lchild // p指向根节点
	for p != t.head {  //p不能指向头结点
		fmt.Printf("%c\n", p.data) //访问
		if p.ltag == LINK {        //p有左孩子
			p = p.lchild //p指向其左孩子(后继)
		} else { //p无左孩子
			p = p.rchild //p指向右孩子(后继)
		}
	}
}

// 通过后序遍历进行后序线索化
func (t *TBiTree) PostThreading() {
	var postThreading func(n *TBiTNode)
	var pre *TBiTNode = t.head //全局变量
	postThreading = func(n *TBiTNode) {
		if n != nil {
			postThreading(n.lchild)
			postThreading(n.rchild)
			if n.lchild == nil { //n没有左孩子
				n.ltag = THREAD
				n.lchild = pre
			}
			if pre.rchild == nil { //n的前驱没有右孩子
				pre.rtag = THREAD
				pre.rchild = n
			}
			pre = n
		}
	}
	postThreading(t.head.lchild)
	t.head.rchild = t.head.lchild
	if pre.rtag != LINK { //如果最后一个结点没有右孩子
		pre.rchild = t.head
		pre.rtag = THREAD
	}
}

// 后序遍历线索二叉树(头结点)的非递归算法
// 二叉链表下不能进行后序遍历.需要三叉链表.
