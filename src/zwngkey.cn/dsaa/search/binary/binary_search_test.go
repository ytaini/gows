/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 02:14:39
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-18 02:46:32
 * @Description:
	二分查找算法(非递归)
*/
package binarysearch

import (
	"fmt"
	"sort"
	"testing"
)

func BinarySearchNoRecur(seq []int, target int) (index int) {
	// 定义初始左右指针分别指向第一个元素与最后一个元素
	left := 0
	right := len(seq) - 1
	for left <= right {
		mid := (left + right) / 2 //mid 始终指向中间元素.
		if seq[mid] == target {
			return mid
		}
		if seq[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearchRecur(seq []int, target int) (index int) {
	var f func(left, right int) int
	f = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		if seq[mid] == target {
			return mid
		}
		if seq[mid] < target {
			return f(mid+1, right)
		}
		return f(left, mid-1)
	}
	return f(0, len(seq)-1)
}

func Test(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7}
	sort.Ints(seq) //先对seq排序.
	index := BinarySearchNoRecur(seq, 7)
	fmt.Println(index)
	index = BinarySearchRecur(seq, 8)
	fmt.Println(index)
}
