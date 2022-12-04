/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-04 01:24:43
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-04 22:05:12
 */
package btree

import (
	"fmt"
	"testing"
)

func TestItems(t *testing.T) {
	is := items{Int(1), Int(2), Int(3), Int(5), Int(6), Int(7), Int(8)}
	fmt.Println(is.find(Int(4)))
}

func Test(t *testing.T) {
	bt := NewBtree(5)
	bt.Insert(Int(4))
	bt.Insert(Int(5))
	bt.Insert(Int(6))
	bt.Insert(Int(7))
	bt.root.items = append(bt.root.items, Int(8))
	bt.PrintBTree()
	fmt.Println(bt.IsBTree())
}
