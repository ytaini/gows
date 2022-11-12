/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-12 16:52:33
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-12 19:17:17
 * @Description:
 */
package main

import (
	"fmt"

	binarytree "zwngkey.cn/dsaa/tree/binary_tree/binary_tree_sequential_storage"
)

func main() {
	bt, err := binarytree.CreateBiTree(6)
	if err != nil {
		fmt.Println(err)
		return
	}
	// bt.PreOrder()
	// bt.InOrder()
	// bt.PostOrder()
	fmt.Printf("bt.BiTreeDepth(): %v\n", bt.BiTreeDepth())
}
