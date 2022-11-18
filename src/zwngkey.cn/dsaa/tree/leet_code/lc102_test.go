/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 11:01:08
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-18 12:03:26
 * @Description:
	102. 二叉树的层序遍历

	给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

	示例1:
			3
		   / \
		  9   20
		     /  \
			15	 7

		输入：root = [3,9,20,null,null,15,7]
		输出：[[3],[9,20],[15,7]]

	示例 2：
		输入：root = [1]
		输出：[[1]]

	示例 3：
		输入：root = []
		输出：[]
*/
package leetcode

import (
	"fmt"
	"testing"

	"zwngkey.cn/dsaa/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 因为go中没有现成的队列可以使用,所以需要通过切片模拟队列
func LevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}

// 使用队列
func LevelOrder1(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}
	q := queue.New[*TreeNode]()
	q.Push(root)
	for !q.IsEmpty() {
		res = append(res, []int{})
		level := len(res) - 1
		for count := q.Len(); count > 0; count-- { //count:当前层元素个数
			node, _ := q.Pop()
			res[level] = append(res[level], node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}
	return res
}

func Test1(t *testing.T) {
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4
	// fmt.Println(LevelOrder1(root))
	fmt.Println(LevelOrder(root))
}
