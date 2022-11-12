/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-13 00:28:41
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-13 03:50:21
 * @Description:
 */
package binarytree

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	root := &BiTNode{data: 0}
	node1 := &BiTNode{data: 1}
	node2 := &BiTNode{data: 2}
	node3 := &BiTNode{data: 3}
	node4 := &BiTNode{data: 4}
	node5 := &BiTNode{data: 5}
	node6 := &BiTNode{data: 6}

	root.lchild = node1
	root.rchild = node2
	node1.lchild = node3
	node1.rchild = node4
	node3.lchild = node5
	node3.rchild = node6

	tree := &BiTree{root}
	// tree.PreOrderTreverseRecus()
	// tree.PreOrderTreverseNonRecur()
	// tree.InOrderTreverseRecur()
	// tree.InOrderTreverseNonRecur()
	// tree.PostOrderTreverseRecur()
	// tree.PostOrderTreverseNonRecur1()
	// tree.PostOrderTreverseNonRecur2()
	// tree.LevelOrderTreverse()

	// is := tree.PreOrderSearchRecur(4)
	// fmt.Println(is)
	// is := tree.InOrderSearchRecur(4)
	// fmt.Println(is)
	// is := tree.PostOrderSearchRecur(4)
	// fmt.Println(is)
	// is := tree.PreOrderSearchNonRecur(4)
	// fmt.Println(is)
	// is := tree.InOrderSearchNonRecur(4)
	// fmt.Println(is)
	// is := tree.PostOrderSearchNonRecur(4)
	// fmt.Println(is)

	// is, node := tree.PreOrderRecurDelete1(0)
	// fmt.Println(is, node)
	// tree.PreOrderTreverseNonRecur()
	is, node := tree.PreOrderRecurDelete2(0)
	fmt.Println(is, node)
	tree.PostOrderTreverseNonRecur2()
}
