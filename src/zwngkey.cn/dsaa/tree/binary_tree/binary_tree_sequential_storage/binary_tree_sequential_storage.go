/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-12 01:45:50
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-12 02:43:20
 * @Description:
	见binary_tree_sequential_storage.md
*/
package binarytree

type BiTNode struct {
	data any
}

type BiTree struct {
	bt []BiTNode
}

func InitBiTree(n int) (b *BiTree) {
	return &BiTree{make([]BiTNode, n)}
}

var ClearBiTree = InitBiTree
var DestroyBiTree = InitBiTree
