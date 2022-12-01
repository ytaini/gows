/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 12:23:01
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-18 13:22:40
 * @Description:
	226. 翻转二叉树
		给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
*/
package leetcode

import (
	"fmt"
	"testing"

	"zwngkey.cn/dsaa/queue"
)

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := root
	for q.Left != nil || q.Right != nil {
		q.Left = InvertTree(q.Left)
		q.Right = InvertTree(q.Right)
		q.Left, q.Right = q.Right, q.Left
		return q
	}
	return q
}

// 递归实现也就是深度优先遍历的方式，那么对应的就是广度优先遍历。
// 使用队列
func InvertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := queue.New[*TreeNode]()
	q.Push(root)
	for !q.IsEmpty() {
		p, _ := q.Pop()
		if p.Left != nil {
			q.Push(p.Left)
		}
		if p.Right != nil {
			q.Push(p.Right)
		}
		if p.Left != nil || p.Right != nil {
			p.Left, p.Right = p.Right, p.Left
		}
	}
	return root
}

// 切片模拟队列
func InvertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		q1 := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			p := q[j]
			if p.Left != nil {
				q1 = append(q1, p.Left)
			}
			if p.Right != nil {
				q1 = append(q1, p.Right)
			}
			if p.Left != nil || p.Right != nil {
				p.Left, p.Right = p.Right, p.Left
			}
		}
		q = q1
	}
	return root
}

func Test3(t *testing.T) {
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	root.Left = node1

	// root = InvertTree(root)
	root = InvertTree1(root)
	fmt.Println(root.Left)
	fmt.Println(root.Right)
}

func Test2(t *testing.T) {
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4
	// root = InvertTree(root)
	root = InvertTree1(root)
	fmt.Println(root.Left)
	fmt.Println(root.Left.Left)
	fmt.Println(root.Left.Right)
	fmt.Println(root.Right)
}
