/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-02 00:38:23
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-02 02:19:19
 */
package tree23

import "fmt"

//  The 2-3 tree is formed by nodes that stores the elements of the structure.
//  Each node contains at most two elements and one at least.
//  In case there is only one element in a node, this is always in the left,so in this case the right element is null.
//  The 2-3 Tree structure defines two type of Nodes/children:
//   - 2 Node : This node only has two children, always left and mid. The right element is empty (null) and the
//              right node/child is also null.
//   - 3 Node : This node has the two elements, so it also has 3 children: left, mid and right. It is full.

type node[T comparable] struct {
	left, mid, right          *node[T]
	leftElement, rightElement T
}

func (n *node[T]) String() string {
	return fmt.Sprintf("node [leftElement=%v,rightElement=%v]", n.leftElement, n.rightElement)
}

// create an empty node/child
func newEmptyNode[T comparable]() *node[T] {
	return &node[T]{}
}

// Constructor of a 3 Node without the children defined yet (null references).
// Precondition(先决条件): The left element must be less than the right element.
func new3NodeNoChildren[T comparable](leftElement, rightElement T) *node[T] {
	return &node[T]{
		leftElement:  leftElement,
		rightElement: rightElement,
	}
}

// Constructor of a 3 Node with the left and mid nodes/children defined.
// Precondition(先决条件): The left element must be less than the right element.
func new3NodeNoRightChild[T comparable](leftElement, rightElement T, left, mid *node[T]) *node[T] {
	return &node[T]{
		left:         left,
		mid:          mid,
		leftElement:  leftElement,
		rightElement: rightElement,
	}
}

// @return true if we are on the deepest level of the tree (a leaf) or false if not
func (n *node[T]) isLeaf() bool {
	return n.left == nil && n.mid == nil && n.right == nil
}

func (n *node[T]) is2Node() bool {
	var t T
	return n.rightElement == t
}

func (n *node[T]) is3Node() bool {
	return !n.is2Node()
}

// isBalanced 检查这棵树是否平衡良好
// @return true if the tree is well-balanced, false if not
func (n *node[T]) isBalanced() bool {
	var t T
	balanced := false
	if n.isLeaf() {
		balanced = true
	} else if n.left.leftElement != t && n.mid.leftElement != t { // There are two cases: 2 Node or 3 Node
		if n.is3Node() { // 3 Node
			if n.right.leftElement != t {
				balanced = true
			}
		} else { // 2 Node
			balanced = true
		}
	}
	return balanced
}

func (n *node[T]) replaceMax() (max T) {
	var t T
	if !n.isLeaf() { // Recursive case, we are not on the deepest level
		if n.rightElement != t {
			max = n.right.replaceMax() // If there is an element on the right, we continue on the right
		} else {
			max = n.mid.replaceMax() // else, we continue on the mid
		}
	} else { // Trivial(简单) case, we are on the deepest level of the tree
		if n.rightElement != t {
			max = n.rightElement
			n.rightElement = t
		} else {
			max = n.leftElement
			n.leftElement = t
		}
	}
	if !n.isBalanced() {
		n.rebalance()
	}
	return max
}

func (n *node[T]) replaceMin() (min T) {
	var t T
	if !n.isLeaf() {
		min = n.left.replaceMin() // 递归情况，只要没有到达最深层，我们总是往左往下走
	} else {
		min = n.leftElement
		n.leftElement = t
		if n.rightElement != t {
			n.leftElement = n.rightElement
			n.rightElement = t
		}
	}
	if !n.isBalanced() { // 这种情况发生在右边没有元素的时候，在第一次上升时会重新平衡
		n.rebalance()
	}
	return min
}

// rebalance the deepest level of the tree from the second deepest.
// 从第二深的树层重新平衡树的最深层。
// The algorithm tries to put one element in each child, but there is a critical case where we must balance the
// - tree from a higher level removing the current level.
func (n *node[T]) rebalance() {
	var t T
	for !n.isBalanced() {
		if n.left.leftElement == t { // The unbalance is in the left child
			n.left.leftElement = n.leftElement
			n.leftElement = n.mid.leftElement

			if n.mid.rightElement != t {
				n.mid.leftElement = n.mid.rightElement
				n.mid.rightElement = t
			} else {
				n.mid.leftElement = t
			}
		} else if n.mid.leftElement == t { // The unbalance is in the right child
			if n.rightElement == t {
				if n.left.leftElement != t && n.left.rightElement == t && n.mid.leftElement == t {
					n.rightElement = n.leftElement
					n.leftElement = n.left.leftElement
					n.left = nil
					n.mid = nil
					n.right = nil
				} else {
					n.mid.leftElement = n.leftElement
					if n.left.rightElement == t {
						n.leftElement = n.left.leftElement
						n.left.leftElement = t
					} else {
						n.leftElement = n.left.rightElement
						n.left.rightElement = t
					}
					if n.left.leftElement == t && n.mid.leftElement == t {
						n.left = nil
						n.mid = nil
						n.right = nil
					}
				}
			} else {
				n.mid.leftElement = n.rightElement
				n.rightElement = n.right.leftElement
				if n.right.rightElement != t {
					n.right.leftElement = n.right.rightElement
					n.right.rightElement = t
				} else {
					n.right.leftElement = t
				}
			}
		} else if n.rightElement != t && n.right.leftElement == t {
			// In this case we can have two situations:
			// (1) The mid child is full, so we have to do a shift of the elements to the right
			// (2) The mid child only has the left element, then we have to put the right element
			//  	   of the current node as the right element of the mid child
			if n.mid.rightElement != t { // (1)
				n.right.leftElement = n.rightElement
				n.rightElement = n.mid.rightElement
				n.mid.rightElement = t
			} else {
				n.mid.rightElement = n.rightElement
				n.rightElement = t
			}
		}
	}
}
