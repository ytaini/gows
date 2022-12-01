/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-17 04:43:08
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-01 15:35:15
 * @Description:
	AVL树的实现.
*/

package avltree

import (
	"fmt"
	"math"
)

// 使用二叉链表来实现二叉树的存储

// AVLNode 平衡二叉树节点类
type AVLNode struct {
	data           int
	lchild, rchild *AVLNode
	height         int // 以该节点作为根节点对应子树的高度，用于计算平衡因子
}

// AVLTree 平衡二叉树结构体
type AVLTree struct {
	root *AVLNode // 根节点
}

// NewAVLTree 平衡二叉树构造函数
func NewAVLTree(data int) *AVLTree {
	return &AVLTree{
		root: &AVLNode{data: data, height: 1},
	}
}

// Find 查找指定节点
func (tree *AVLTree) Find(data int) *AVLNode {
	// 空树直接返回空
	if tree.root == nil {
		return nil
	}
	return tree.root.Find(data)
}
func (node *AVLNode) Find(data int) *AVLNode {
	if data == node.data {
		return node
	} else if data < node.data {
		// 如果查找的值小于节点值，从节点的左子树开始查找
		if node.lchild == nil {
			return nil
		}
		return node.lchild.Find(data)
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始查找
		if node.rchild == nil {
			return nil
		}
		return node.rchild.Find(data)
	}
}

// Insert 插入节点到平衡二叉树
func (tree *AVLTree) Insert(data int) {
	// 从根节点开始插入数据
	// 根节点在动态变化，所以需要不断刷新
	tree.root = tree.root.Insert(data)
}

func (node *AVLNode) Insert(data int) *AVLNode {
	// 如果节点为空，则初始化该节点
	if node == nil {
		return &AVLNode{data: data, height: 1}
	}
	// 如果值重复，则什么都不做
	if node.data == data {
		return node
	}
	// 辅助变量，用于存储（旋转后）子树根节点
	var newTreeNode *AVLNode
	if data > node.data {
		// 插入的值大于当前节点值，要从右子树插入
		node.rchild = node.rchild.Insert(data)
		// 计算插入节点后当前节点的平衡因子
		// 按照平衡二叉树的特征，平衡因子绝对值不能大于 1
		bf := node.BalanceFactor()
		// 如果右子树高度变高了，导致左子树-右子树的高度从 -1 变成了 -2
		if bf == -2 {
			if data > node.rchild.data {
				// 表示在右子树中插入右子节点导致失衡，需要单左旋
				newTreeNode = LeftRotate(node)
			} else {
				// 表示在右子树中插上左子节点导致失衡，需要先右后左双旋
				newTreeNode = RightLeftRotation(node)
			}
		}
	} else {
		// 插入的值小于当前节点值，要从左子树插入
		node.lchild = node.lchild.Insert(data)
		bf := node.BalanceFactor()
		//  左子树的高度变高了，导致左子树-右子树的高度从 1 变成了 2
		if bf == 2 {
			if data < node.lchild.data {
				// 表示在左子树中插入左子节点导致失衡，需要单右旋
				newTreeNode = RightRotate(node)
			} else {
				// 表示在左子树中插入右子节点导致失衡，需要先左后右双旋
				newTreeNode = LeftRightRotation(node)
			}
		}
	}
	if newTreeNode == nil {
		// 根节点没变，直接更新子树高度，并返回当前节点指针
		node.Updateheight()
		return node
	} else {
		newTreeNode.Updateheight()
		return newTreeNode
	}
}

// Updateheight 更新节点树高度
func (node *AVLNode) Updateheight() {
	if node == nil {
		return
	}
	// 分别计算左子树和右子树的高度
	leftheight, rightheight := 0, 0
	if node.lchild != nil {
		leftheight = node.lchild.height
	}
	if node.rchild != nil {
		rightheight = node.rchild.height
	}
	// 以更高的子树高度作为节点树高度
	maxheight := leftheight
	if rightheight > maxheight {
		maxheight = rightheight
	}
	// 最终高度要加上节点本身所在的那一层
	node.height = maxheight + 1
}

// BalanceFactor 计算节点平衡因子（即左右子树的高度差）
func (node *AVLNode) BalanceFactor() int {
	leftheight, rightheight := 0, 0
	if node.lchild != nil {
		leftheight = node.lchild.height
	}
	if node.rchild != nil {
		rightheight = node.rchild.height
	}
	return leftheight - rightheight
}

// RightRotate 右旋操作
func RightRotate(node *AVLNode) *AVLNode {
	pivot := node.lchild   // pivot 表示新插入的节点
	pivotR := pivot.rchild // 暂存 pivot 右子树入口节点
	pivot.rchild = node    // 右旋后最小不平衡子树根节点 node 变成 pivot 的右子节点
	node.lchild = pivotR   // 而 pivot 原本的右子节点需要挂载到 node 节点的左子树上
	// 只有 node 和 pivot 的高度改变了
	node.Updateheight()
	pivot.Updateheight()
	// 返回右旋后的子树根节点指针，即 pivot
	return pivot
}

// LeftRotate 左旋操作
func LeftRotate(node *AVLNode) *AVLNode {
	pivot := node.rchild   // pivot 表示新插入的节点
	pivotL := pivot.lchild // 暂存 pivot 左子树入口节点
	pivot.lchild = node    // 左旋后最小不平衡子树根节点 node 变成 pivot 的左子节点
	node.rchild = pivotL   // 而 pivot 原本的左子节点需要挂载到 node 节点的右子树上

	// 只有 node 和 pivot 的高度改变了
	node.Updateheight()
	pivot.Updateheight()

	// 返回旋后的子树根节点指针，即 pivot
	return pivot
}

// LeftRightRotation 双旋操作（先左后右）
func LeftRightRotation(node *AVLNode) *AVLNode {
	node.rchild = LeftRotate(node.rchild)
	return RightRotate(node)
}

// RightLeftRotation 先右旋后左旋
func RightLeftRotation(node *AVLNode) *AVLNode {
	node.rchild = RightRotate(node.rchild)
	return LeftRotate(node)
}

// Traverse 中序遍历平衡二叉树
func (tree *AVLTree) Traverse() {
	// 从根节点开始遍历
	tree.root.Traverse()
}

func (node *AVLNode) Traverse() {
	// 节点为空则退出当前递归
	if node == nil {
		return
	}
	// 否则先从左子树最左侧节点开始遍历
	node.rchild.Traverse()
	// 打印位于中间的根节点
	fmt.Printf("%d(%d) ", node.data, node.BalanceFactor())
	// 最后按照和左子树一样的逻辑遍历右子树
	node.rchild.Traverse()
}

// IsAVLTree 判断是不是平衡二叉树
func (tree *AVLTree) IsAVLTree() bool {
	if tree == nil || tree.root == nil {
		return true
	}

	// 判断每个节点是否符合平衡二叉树的定义
	if tree.root.IsBalanced() {
		return true
	}

	return false
}

// IsBalanced 判断节点是否符合平衡二叉树的定义
func (node *AVLNode) IsBalanced() bool {
	// 左右子树都为空是叶子节点
	if node.lchild == nil && node.rchild == nil {
		// 叶子节点高度都是 1
		if node.height == 1 {
			return true
		} else {
			fmt.Println("叶子节点高度值: ", node.height)
			return false
		}
	} else if node.lchild != nil && node.rchild != nil {
		// 左右子树不为空
		// 左子树所有节点值必须比父节点小，右子树所有节点值必须比父节点大（AVL 树首先是二叉排序树）
		if node.lchild.data > node.data || node.rchild.data < node.data {
			// 不符合 AVL 树定义
			fmt.Printf("父节点值是 %v, 左子节点值是 %v, 右子节点值是 %v\n", node.data, node.lchild.data, node.rchild.data)
			return false
		}
		// 计算平衡因子 BF 绝对值
		bf := node.lchild.height - node.rchild.height
		// 平衡因子不能大于 1
		if math.Abs(float64(bf)) > 1 {
			fmt.Println("平衡因子 BF 值: ", bf)
			return false
		}
		// 如果左子树比右子树高，那么父节点的高度等于左子树 +1
		if node.lchild.height > node.rchild.height {
			if node.height != node.lchild.height+1 {
				fmt.Printf("%#v 高度: %v, 左子树高度: %v, 右子树高度: %v\n", node, node.height, node.lchild.height, node.rchild.height)
				return false
			}
		} else {
			// 如果右子树比左子树高，那么父节点的高度等于右子树 +1
			if node.height != node.rchild.height+1 {
				fmt.Printf("%#v 高度: %v, 左子树高度: %v, 右子树高度: %v\n", node, node.height, node.lchild.height, node.rchild.height)
				return false
			}
		}
		// 递归判断左子树
		if !node.lchild.IsBalanced() {
			return false
		}
		// 递归判断右子树
		if !node.rchild.IsBalanced() {
			return false
		}
	} else {
		// 只存在一棵子树
		if node.rchild != nil {
			// 子树高度只能是 1
			if node.rchild.height == 1 && node.rchild.lchild == nil && node.rchild.rchild == nil {
				if node.rchild.data < node.data {
					// 右子节点值必须比父节点值大
					fmt.Printf("节点值: %v,(%#v,%#v)", node.data, node.rchild, node.lchild)
					return false
				}
			} else {
				fmt.Printf("节点值: %v,(%#v,%#v)", node.data, node.rchild, node.lchild)
				return false
			}
		} else {
			if node.lchild.height == 1 && node.lchild.lchild == nil && node.lchild.rchild == nil {
				if node.lchild.data > node.data {
					// 左子节点值必须比父节点值小
					fmt.Printf("节点值: %v,(%#v,%#v) child", node.data, node.rchild, node.lchild)
					return false
				}
			} else {
				fmt.Printf("节点值: %v,(%#v,%#v) child", node.data, node.rchild, node.lchild)
				return false
			}
		}
	}
	return true
}

// Delete 删除指定节点
func (tree *AVLTree) Delete(data int) {
	// 空树直接返回
	if tree.root == nil {
		return
	}
	// 删除指定节点，和插入节点一样，根节点也会随着 AVL 树的旋转动态变化
	tree.root = tree.root.Delete(data)
}

func (node *AVLNode) Delete(data int) *AVLNode {
	// 空节点直接返回 nil
	if node == nil {
		return nil
	}
	if data < node.data {
		// 如果删除节点值小于当前节点值，则进入当前节点的左子树删除元素
		node.lchild = node.lchild.Delete(data)
		// 删除后要更新左子树高度
		node.lchild.Updateheight()
	} else if data > node.data {
		// 如果删除节点值大于当前节点值，则进入当前节点的右子树删除元素
		node.rchild = node.rchild.Delete(data)
		// 删除后要更新右子树高度
		node.rchild.Updateheight()
	} else {
		// 找到待删除节点后
		// 第一种情况，删除的节点没有左右子树，直接删除即可
		if node.lchild == nil && node.rchild == nil {
			// 返回 nil 表示直接删除该节点
			return nil
		}
		// 第二种情况，待删除节点包含左右子树，选择高度更高的子树下的节点来替换待删除的节点
		// 如果左子树更高，选择左子树中值最大的节点，也就是左子树最右边的叶子节点
		// 如果右子树更高，选择右子树中值最小的节点，也就是右子树最左边的叶子节点
		// 最后，删除这个叶子节点即可
		if node.lchild != nil && node.rchild != nil {
			// 左子树更高，拿左子树中值最大的节点替换
			if node.lchild.height > node.rchild.height {
				maxNode := node.lchild
				for maxNode.rchild != nil {
					maxNode = maxNode.rchild
				}
				// 将值最大的节点值赋值给待删除节点
				node.data = maxNode.data
				// 然后把值最大的节点删除
				node.lchild = node.lchild.Delete(maxNode.data)
				// 删除后要更新左子树高度
				node.lchild.Updateheight()
			} else {
				// 右子树更高，拿右子树中值最小的节点替换
				minNode := node.rchild
				for minNode.lchild != nil {
					minNode = minNode.lchild
				}
				// 将值最小的节点值赋值给待删除节点
				node.data = minNode.data
				// 然后把值最小的节点删除
				node.rchild = node.rchild.Delete(minNode.data)
				// 删除后要更新右子树高度
				node.rchild.Updateheight()
			}
		} else {
			// 只有左子树或只有右子树
			// 只有一棵子树，该子树也只是一个节点，则将该节点值赋值给待删除的节点，然后置子树为空
			if node.lchild != nil {
				// 第三种情况，删除的节点只有左子树
				// 根据 AVL 树的特征，可以知道左子树其实就只有一个节点，否则高度差就等于 2 了
				node.data = node.lchild.data
				node.height = 1
				node.lchild = nil
			} else if node.rchild != nil {
				// 第四种情况，删除的节点只有右子树
				// 根据 AVL 树的特征，可以知道右子树其实就只有一个节点，否则高度差就等于 2 了
				node.data = node.rchild.data
				node.height = 1
				node.rchild = nil
			}
		}
		// 找到指定值对应的待删除节点并进行替换删除后，直接返回该节点
		return node
	}
	// 左右子树递归删除节点后需要平衡
	var newNode *AVLNode
	// 相当删除了右子树的节点，左边比右边高了，不平衡
	if node.BalanceFactor() == 2 {
		if node.lchild.BalanceFactor() >= 0 {
			newNode = RightRotate(node)
		} else {
			newNode = LeftRightRotation(node)
		}
		//  相当删除了左子树的节点，右边比左边高了，不平衡
	} else if node.BalanceFactor() == -2 {
		if node.rchild.BalanceFactor() <= 0 {
			newNode = LeftRotate(node)
		} else {
			newNode = RightLeftRotation(node)
		}
	}
	if newNode == nil {
		node.Updateheight()
		return node
	} else {
		newNode.Updateheight()
		return newNode
	}
}
