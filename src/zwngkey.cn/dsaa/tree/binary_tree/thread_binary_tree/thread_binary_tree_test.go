/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-14 03:51:16
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-14 06:13:18
 * @Description:
 */
package threadbinarytree

import (
	"testing"
)

func Test(t *testing.T) {
	s := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	tree := CreateTBiTree(s)
	// tree.InThreading()
	// tree.InOrderTraverse()
	// tree.PreThreading()
	// tree.PreOrderTraverse()
	tree.PostThreading()
}
