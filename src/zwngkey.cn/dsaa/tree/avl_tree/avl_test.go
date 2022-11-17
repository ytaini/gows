/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-17 07:15:27
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-17 09:06:49
 * @Description:
 */
package avltree

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	seq := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	avlTree := NewAVLTree(seq[0])
	for i := 1; i < len(seq); i++ {
		avlTree.Insert(seq[i])
	}
	// 判断是否是平衡二叉树
	fmt.Print("avlTree 是平衡二叉树: ")
	fmt.Println(avlTree.IsAVLTree())

	// 中序遍历生成的二叉树看是否是二叉排序树
	fmt.Print("中序遍历结果: ")
	avlTree.Traverse()
	fmt.Println()

	// 查找节点
	fmt.Print("查找值为 5 的节点: ")
	fmt.Printf("%v\n", avlTree.Find(5))

	// 删除节点
	avlTree.Delete(5)

	// 删除后是否还是平衡二叉树
	fmt.Print("avlTree 仍然是平衡二叉树: ")
	fmt.Println(avlTree.IsAVLTree())
	fmt.Print("删除节点后的中序遍历结果: ")
	avlTree.Traverse()
	fmt.Println()
}
