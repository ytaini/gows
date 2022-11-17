/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-17 00:35:01
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-17 04:13:42
 * @Description:
	二叉排序树的创建和遍历,删除
*/
package binarysorttree

import "fmt"

type TreeNode struct {
	data           int
	lchild, rchild *TreeNode
}

func (t TreeNode) String() string {
	return fmt.Sprintf("Node [data=%v]", t.data)
}

type Tree struct {
	root *TreeNode
}

func NewTree(data int) *Tree {
	return &Tree{&TreeNode{data: data}}
}

// 通过数组创建成对应的二叉排序树
func CreateBST(seq []int) *Tree {
	t := NewTree(seq[0])
	for i := 1; i < len(seq); i++ {
		t.Add(&TreeNode{data: seq[i]})
	}
	return t
}

// 向二叉排序树中添加节点
func (t *Tree) Add(node *TreeNode) *TreeNode {
	var add func(node, root *TreeNode)
	add = func(node, root *TreeNode) {
		if node == nil {
			return
		}
		//先与当前结点进行比较
		if node.data < root.data {
			if root.lchild == nil { //若比当前结点小且当前的左节点为空.
				root.lchild = node //那么当前结点的左节点就是node
			} else {
				add(node, root.lchild) //否则递归比较
			}
		} else {
			if root.rchild == nil { //若比当前结点大且当前的右节点为空.
				root.rchild = node //那么当前结点的右节点就是node
			} else {
				add(node, root.rchild)
			}
		}
	}
	add(node, t.root)
	return node
}

// 中序遍历二叉排序树
// 中序遍历的结果刚好是递增的
func (t *Tree) InOrderTraverse() {
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node != nil {
			f(node.lchild)
			fmt.Printf("%v->", node.data)
			f(node.rchild)
		}
	}
	f(t.root)
	fmt.Println()
}

// 前序遍历二叉排序树
func (t *Tree) PreOrderTraverse() {
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node != nil {
			fmt.Printf("%v->", node.data)
			f(node.lchild)
			f(node.rchild)
		}
	}
	f(t.root)
	fmt.Println()
}

// 二叉排序树的删除情况比较复杂，有下面三种情况需要考虑：
// 	删除叶子节点 (比如：2,5,9,12)
// 	删除只有一颗子树的节点 (比如：1)
// 	删除有两颗子树的节点. (比如：7,3,10)

// 1.删除叶子节点 (比如：2,5,9,12)
// 	需求先去找到要删除的结点 targetNode
// 	找到 targetNode 的 父结点 parent
// 	确定 targetNode 是 parent 的左子结点 还是右子结点
// 	根据前面的情况来对应删除 左子结点 parent.left = null 右子结点 parent.right = null;

// 2.删除有两颗子树的节点，比如：(7, 3，10 )
// 	需求先去找到要删除的结点 targetNode
// 	找到 targetNode 的 父结点 parent
// 	从 targetNode 的右子树找到最小的结点
// 	用一个临时变量，将最小结点的值保存 temp
// 	删除该最小结点
// 	targetNode.value = temp

// 3.删除只有一颗子树的节点，比如(1)
// 	需求先去找到要删除的结点 targetNode
// 	找到 targetNode 的 父结点 parent
// 	确定 targetNode 的子结点是左子结点还是右子结点
// 	targetNode 是 parent 的左子结点还是右子结点
// 	如果 targetNode 有左子结点
// 		如果 targetNode 是 parent 的左子结点 parent.left = targetNode.left
// 		如果 targetNode 是 parent 的右子结点 parent.right = targetNode.left
// 	如果 targetNode 有右子结点
// 		如果 targetNode 是 parent 的左子结点 parent.left = targetNode.right
// 		如果 targetNode 是 parent 的右子结点 parent.right = targetNode.right
func (t *Tree) Delete(data int) bool {
	curNode, parentNode := t.Find(data)
	if curNode == nil {
		return false
	}
	if curNode.lchild == nil && curNode.rchild == nil {
		if parentNode == nil {
			t.root = nil
		} else {
			SetNil(curNode, parentNode)
		}
		return true
	}
	if curNode.lchild != nil && curNode.rchild != nil {
		minNode, minParentNode := curNode.findMinAndParent()
		curNode.data = minNode.data
		SetNil(minNode, minParentNode)
		return true
	}
	//说明curNode此时为根节点,且其只有一个子节点.
	if parentNode == nil {
		if curNode.lchild != nil {
			curNode.data = curNode.lchild.data
			curNode.lchild = nil
		} else {
			curNode.data = curNode.rchild.data
			curNode.rchild = nil
		}
		return true
	}
	//说明curNode此时为非根节点,且其只有一个子节点.
	if parentNode.lchild == curNode {
		if curNode.lchild != nil {
			parentNode.lchild = curNode.lchild
		} else {
			parentNode.lchild = curNode.rchild
		}
	} else {
		if curNode.lchild != nil {
			parentNode.rchild = curNode.lchild
		} else {
			parentNode.rchild = curNode.rchild
		}
	}
	return true
}

//递归算法 寻找node.data = data 的结点及其父节点.
func (t *Tree) Find(data int) (cur *TreeNode, parent *TreeNode) {
	var find func(root *TreeNode, data int) (*TreeNode, *TreeNode)
	find = func(root *TreeNode, data int) (*TreeNode, *TreeNode) {
		if root != nil {
			if root.data == data {
				return root, parent
			}
			parent = root
			if root.data < data {
				return find(root.rchild, data)
			} else {
				return find(root.lchild, data)
			}
		}
		return nil, nil
	}
	return find(t.root, data)
}

//非递归算法 寻找node.data = data 的结点及其父节点.
func (t *Tree) Find2(data int) (cur *TreeNode, parent *TreeNode) {
	cur = t.root
	for cur != nil {
		if cur.data == data {
			return
		}
		parent = cur
		if cur.data < data {
			cur = cur.rchild
		} else {
			cur = cur.lchild
		}
	}
	return nil, nil
}

// 将parent的左或右孩子置为nil
func SetNil(cur *TreeNode, parent *TreeNode) {
	if parent.lchild == cur {
		parent.lchild = nil
	} else {
		parent.rchild = nil
	}
}

// 寻找以t为根节点的树中的最小值节点及其父节点
func (t *TreeNode) findMinAndParent() (cur *TreeNode, parent *TreeNode) {
	parent = t
	var min func(t *TreeNode) *TreeNode
	min = func(t *TreeNode) *TreeNode {
		if t.lchild == nil {
			return t
		}
		parent = t
		return min(t.lchild)
	}
	cur = min(t.rchild)
	return
}
