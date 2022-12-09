/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 04:49:06
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 06:06:40
 */
package main

import (
	"fmt"

	. "zwngkey.cn/designpattern/behavioral/strategy/case2"
)

type Int int

// 具体策略类
type IntComparator struct{}

func (IntComparator) Compare(i1, i2 Int) int {
	return int(i1 - i2)
}

func main() {
	cats := []Cat{
		{Height: 3, Weight: 3},
		{Height: 1, Weight: 1},
		{Height: 5, Weight: 5},
		{Height: 2, Weight: 2},
		{Height: 4, Weight: 4},
		{Height: 4, Weight: 4},
	}
	// 基于不同的策略进行排序.
	// sorter := NewSorter[Cat](CatWeightComparator{}) //基于Cat类型的Weight属性非递减排序
	sorter := NewSelectSorter[Cat](CatHeightComparator{}) //基于Cat类型的Height非递增排序
	sorter.Sort(cats)
	fmt.Println(cats)

	ints := []Int{3, 1, 5, 2, 6, 2, 8, 1}
	sorter1 := NewSelectSorter[Int](IntComparator{})
	sorter1.Sort(ints)
	fmt.Println(ints)

}
