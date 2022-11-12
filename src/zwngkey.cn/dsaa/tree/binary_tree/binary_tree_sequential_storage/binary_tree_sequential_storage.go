/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-12 01:45:50
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-12 19:49:02
 * @Description:
	见binary_tree_sequential_storage.md
*/
package binarytree

import (
	"fmt"
	"math"
)

type BiTNode struct {
	data any
}
type BiTree struct {
	bt []BiTNode
}

// #号代表空节点.
// 按满二叉树的结点层次编号,依次存放二叉树中的数据元素
func CreateBiTree(n int) (b *BiTree, err error) {
	var s string
	fmt.Scanln(&s)
	b = &BiTree{make([]BiTNode, len(s))}
	for i := 0; i < len(s); i++ {
		b.bt[i].data = s[i]
	}
	// 判断输入的树是否合法
	for i := 1; i < len(b.bt); i++ {
		ch := b.bt[(i+1)/2-1].data.(byte)
		if ch == '#' && b.bt[i].data.(byte) != '#' { //此非根节点(不空)无双亲.
			return nil, fmt.Errorf("出现无双亲的非根节点")
		}
	}
	return b, nil
}

func (b BiTree) preOrder(n int) {
	fmt.Printf("%c ", b.bt[n].data)
	if 2*n+1 < len(b.bt) {
		b.preOrder(2*n + 1)
	}
	if 2*n+2 < len(b.bt) {
		b.preOrder(2*n + 2)
	}
}

// 顺序存储二叉树的前序遍历
func (b BiTree) PreOrder() {
	if len(b.bt) == 0 {
		return
	}
	fmt.Print("前序遍历:")
	b.preOrder(0)
	fmt.Println()
}

func (b BiTree) inOrder(n int) {
	if 2*n+1 < len(b.bt) {
		b.inOrder(2*n + 1)
	}
	fmt.Printf("%c ", b.bt[n].data)
	if 2*n+2 < len(b.bt) {
		b.inOrder(2*n + 2)
	}
}

// 顺序存储二叉树的中序遍历
func (b BiTree) InOrder() {
	if len(b.bt) == 0 {
		return
	}
	fmt.Print("中序遍历:")
	b.inOrder(0)
	fmt.Println()
}

func (b BiTree) postOrder(n int) {
	if 2*n+1 < len(b.bt) {
		b.postOrder(2*n + 1)
	}
	if 2*n+2 < len(b.bt) {
		b.postOrder(2*n + 2)
	}
	fmt.Printf("%c ", b.bt[n].data)
}

// 顺序存储二叉树的后序遍历
func (b BiTree) PostOrder() {
	if len(b.bt) == 0 {
		return
	}
	fmt.Print("后序遍历:")
	b.postOrder(0)
	fmt.Println()
}

// 返回树的深度
func (b BiTree) BiTreeDepth() int {
	return int(math.Log2(float64(len(b.bt)))) + 1
}
