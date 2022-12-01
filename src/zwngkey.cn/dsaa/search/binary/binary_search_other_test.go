/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 03:13:03
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-18 04:21:08
 * @Description:
	二分查找的变形版本:
		1.从给定排序序列中查找第一个等于给定值的元素
		2.从给定序列中查找最后一个匹配元素
		3.在给定序列中查找第一个大于等于给定值的元素
		4.在给定序列中查找最后一个小于等于给定值的元素
*/
package binarysearch

import (
	"fmt"
	"testing"
)

// 1.从给定排序序列中查找第一个等于给定值的元素
func FindFirstEleNonRecur(seq []int, target int) int {
	left := 0
	right := len(seq) - 1
	for left <= right {
		mid := (left + right) / 2
		if seq[mid] == target {
			if mid == 0 || seq[mid-1] != target {
				return mid
			}
			right = mid - 1
		} else if seq[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func FindFirstEleRecur(seq []int, target int) int {
	var f func(left, right int) int
	f = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		if seq[mid] == target {
			if mid == 0 || seq[mid-1] != target {
				return mid
			}
			return f(left, mid-1)
		} else if seq[mid] < target {
			return f(mid+1, right)
		} else {
			return f(left, mid-1)
		}
	}
	return f(0, len(seq)-1)
}

func Test1(t *testing.T) {
	seq := []int{1, 2, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8}
	fmt.Println(FindFirstEleRecur(seq, 11))
	fmt.Println(FindFirstEleNonRecur(seq, 11))
}

// 2. 从给定序列中查找最后一个匹配元素
func FindLastEleNonRecur(seq []int, target int) int {
	left := 0
	right := len(seq) - 1
	for left <= right {
		mid := (left + right) / 2
		if seq[mid] == target {
			if mid == len(seq)-1 || seq[mid+1] != target {
				return mid
			}
			left = mid + 1
		} else if seq[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func FindLastEleRecur(seq []int, target int) int {
	var f func(left, right int) int
	f = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		if seq[mid] == target {
			if mid == len(seq)-1 || seq[mid+1] != target {
				return mid
			}
			return f(mid+1, right)
		} else if seq[mid] < target {
			return f(mid+1, right)
		} else {
			return f(left, mid-1)
		}
	}
	return f(0, len(seq)-1)
}

func Test2(t *testing.T) {
	seq := []int{1, 2, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8}
	fmt.Println(FindLastEleNonRecur(seq, 4))
	fmt.Println(FindLastEleRecur(seq, 4))
}

// 3.在给定序列中查找第一个大于等于给定值的元素
func FindEleNonRecur(seq []int, target int) int {
	left := 0
	right := len(seq) - 1
	for left <= right {
		mid := (left + right) / 2
		if seq[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || seq[mid-1] < target {
				return mid
			}
			right = mid - 1
		}
	}
	return -1
}

func FindEleRecur(seq []int, target int) int {
	var f func(left, right int) int
	f = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		if seq[mid] < target {
			return f(mid+1, right)
		} else {
			if mid == 0 || seq[mid-1] < target {
				return mid
			}
			return f(left, mid-1)
		}
	}
	return f(0, len(seq)-1)
}

func Test3(t *testing.T) {
	seq := []int{1, 3, 3, 4, 5, 5, 7, 7, 8, 9, 9}
	fmt.Println(FindEleRecur(seq, 6))
	fmt.Println(FindEleNonRecur(seq, 6))
}

// 4. 在给定序列中查找最后一个小于等于给定值的元素
func FindEleNonRecur1(seq []int, target int) int {
	left := 0
	right := len(seq) - 1
	for left <= right {
		mid := (left + right) / 2
		if seq[mid] > target {
			right = mid - 1
		} else {
			if mid == len(seq)-1 || seq[mid+1] > target {
				return mid
			}
			left = mid + 1
		}
	}
	return -1
}

func FindEleRecur1(seq []int, target int) int {
	var f func(left, right int) int
	f = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		if seq[mid] > target {
			return f(left, mid-1)
		} else {
			if mid == len(seq)-1 || seq[mid+1] > target {
				return mid
			}
			return f(mid+1, right)
		}
	}
	return f(0, len(seq)-1)
}

func Test4(t *testing.T) {
	seq := []int{1, 3, 3, 4, 5, 5, 7, 7, 8, 9, 9}
	fmt.Println(FindEleRecur1(seq, 5))
	fmt.Println(FindEleNonRecur1(seq, 5))
}
