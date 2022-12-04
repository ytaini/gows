/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-03 23:36:22
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-04 22:29:14
 */
package btree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "zwngkey.cn/dsaa/tree/multi_way_search_tree/btree/btree"
)

func Test1(t *testing.T) {
	bt := NewBtree(5)
	bt.Insert(Int(4))
	bt.Insert(Int(5))
	bt.Insert(Int(6))
	bt.Insert(Int(7))
	bt.Insert(Int(8))
	bt.Insert(Int(9))
	bt.Insert(Int(10))
	bt.Insert(Int(11))
	bt.InsertMultiple([]int{20, 44, 89, 96, 25, 30, 33, 60, 75, 81, 85, 110, 120})
	bt.Insert(Int(101))
	bt.InsertMultiple([]int{150, 158, 130, 135, 138})
	bt.PrintBTree()
	bt.Delete(Int(158))
	bt.Delete(Int(150))
	bt.Delete(Int(6))
	bt.Delete(Int(9))
	bt.PrintBTree()
	fmt.Println(bt.IsBTree())
}

func Test2(t *testing.T) {
	bt := NewBtree(10)
	rand.Seed(time.Now().UnixNano())
	t1 := time.Now()
	for i := 0; i < 10000; i++ {
		bt.Insert(Int(rand.Int63n(100000)))
	}
	t2 := time.Since(t1)
	fmt.Println(t2)
	// for i := 0; i < 100000; i++ {
	// 	bt.Delete(Int(rand.Int63n(100000)))
	// }
	bt.PrintBTree()
}
