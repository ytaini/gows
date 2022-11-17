/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-17 01:05:19
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-17 04:13:12
 * @Description:
 */
package binarysorttree

import (
	"testing"
)

func Test(t *testing.T) {
	seq := []int{7, 3, 10, 1, 5, 9, 12, 2, 11}
	bst := CreateBST(seq)
	// fmt.Println(bst.Find2(5))
	// bst.Delete(1)
	// bst.Delete(5)
	// bst.Delete(9)
	// bst.Delete(2)
	// bst.Delete(11)
	// bst.Delete(3)
	bst.Delete(7)
	bst.InOrderTraverse()
	// bst.PreOrderTraverse()
}
