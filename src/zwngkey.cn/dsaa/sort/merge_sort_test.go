/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-04 17:24:17
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-05 00:41:19
 * @Description:
	归并排序:
		是利用归并的思想实现的排序方法，该算法采用经典的分治（divide-and-conquer）策略
			（分治法将问题分(divide)成一些小的问题然后递归求解，而治(conquer)的阶段则将分的阶段得到的各答案"修补"在一起，即分而治之)。

	归并排序，基于比较，稳定算法，时间O(nlogn)，空间O(logn) | O(n)

	归并排序可以采用递归去实现（也可采用迭代的方式去实现）。分阶段可以理解为就是递归拆分子序列的过程。
		都需要开辟一个大小为n的数组中转

	归并排序的思想主要是利用栈的特性，递归进行分解操作，肯定是从大数组向小数组进行分解，递归进行合并操作，肯定是从小数组向大数组进行合并

	算法:
		将数组分为左右两部分，递归左右两块，最后合并，即归并
		如在一个合并中，将两块部分的元素，遍历取较小值填入结果集
		类似两个有序链表的合并，每次两两合并相邻的两个有序序列，直到整个序列有序

*/
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func merge1(left, right []int) []int {
	temp := make([]int, len(left)+len(right))
	lp, rp := 0, 0
	i := 0
	// 通过遍历完成比较填入temp中
	for ; lp < len(left) && rp < len(right); i++ {
		if left[lp] <= right[rp] {
			temp[i] = left[lp]
			lp++
		} else {
			temp[i] = right[rp]
			rp++
		}
	}
	// 如果left或者right还有剩余元素，添加到结果集的尾部
	// for j := lp; j < len(left); j++ {
	// 	temp[i] = left[j]
	// 	i++
	// }
	// for j := rp; j < len(right); j++ {
	// 	temp[i] = right[j]
	// 	i++
	// }

	// 方式二:
	// 如果left或者right还有剩余元素，添加到结果集的尾部
	copy(temp[i:], left[lp:])
	copy(temp[i:], right[rp:])
	return temp
}

// 递归，自下而上的方式合并
func MergeSort1(s []int) []int {
	var mergeSort func(s []int) []int
	mergeSort = func(s []int) []int {
		if len(s) <= 1 {
			return s
		}
		split := len(s) / 2
		// 分
		right := mergeSort(s[:split])
		left := mergeSort(s[split:])
		// 治(合并)
		return merge1(right, left)
	}
	temp := mergeSort(s)
	copy(s, temp) // 可以用返回值,也可以直接用s
	return temp
}

func merge2(s []int, left, right, mid int) {
	temp := make([]int, right-left+1)
	l, r, i := left, mid+1, 0
	for l <= mid && r <= right {
		if s[l] <= s[r] {
			temp[i] = s[l]
			l++
		} else {
			temp[i] = s[r]
			r++
		}
		i++
	}
	for l <= mid {
		temp[i] = s[l]
		l++
		i++
	}
	for r <= right {
		temp[i] = s[r]
		r++
		i++
	}
	copy(s[left:right+1], temp)
}

// 递归，自下而上的方式合并
func MergeSort2(s []int) []int {
	var mergeSort func(s []int, left, right int)
	mergeSort = func(s []int, left, right int) {
		if left < right {
			mid := (left + right) / 2
			mergeSort(s, left, mid)
			mergeSort(s, mid+1, right)
			merge2(s, left, right, mid)
		}
	}
	mergeSort(s, 0, len(s)-1)
	return s
}

// 非递归现实归并,自下而上的方式合并
// 非递归的重点在于如何确定并合理的分解待排序数组。
func MergeSort3(s []int) []int {
	//通过 gap 将s划分为len(s)/2gap个区间.
	//当 gap 为 1 的时,s被划分为4个区间，对每个区间进行归并.
	//当 gap 为 2 的时,s被划分为2个区间，对每个区间进行归并.
	//当 gap 为 4 的时,s被划分为1个区间，对每个区间进行归并.
	for gap := 1; gap < len(s); gap *= 2 {
		for j := 0; j+gap <= len(s); j += gap * 2 {
			l := j
			mid := j + gap - 1
			r := j + 2*gap - 1
			if r > len(s)-1 { //整个待排序数组为奇数的情况
				r = len(s) - 1
			}
			merge2(s, l, r, mid)
		}
	}
	return s
}

// 非递归现实归并,自下而上的方式合并
func MergeSort4(s []int) []int {
	res := make([]int, 0)
	for i := 1; i < len(s); i *= 2 {
		for j := 0; j+i < len(s); j = j + 2*i {
			if j+2*i > len(s) {
				res = merge1(s[j:j+i], s[j+i:])
			} else {
				res = merge1(s[j:j+i], s[j+i:j+2*i])
			}
			index := j
			for _, v := range res {
				s[index] = v
				index++
			}
		}
	}
	return res
}

func Test11(t *testing.T) {
	s := []int{11, 8, 3, 9, 7, 1, 2}
	// MergeSort1(s)
	// MergeSort2(s)
	// MergeSort3(s)
	s1 := MergeSort4(s)
	fmt.Println(s)
	fmt.Println(s1)
}

// 比较归并3与归并4
func Test15s(t *testing.T) {
	seq1 := make([]int, 10000000)
	seq := make([]int, 10000000)
	for i := 0; i < 10000000; i++ {
		v := rand.Intn(100000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	MergeSort3(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	MergeSort4(seq1)
	t2 = time.Since(t1)
	fmt.Println(t2)

	for i := 0; i < 10000000; i++ {
		if seq[i] != seq1[i] {
			fmt.Println("false")
			return
		}
	}
	fmt.Println(true)
}

// 比较归并2与归并3
func Test14(t *testing.T) {
	seq1 := make([]int, 5000000)
	seq := make([]int, 5000000)
	for i := 0; i < 5000000; i++ {
		v := rand.Intn(100000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	MergeSort1(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	MergeSort3(seq1)
	t2 = time.Since(t1)
	fmt.Println(t2)

	for i := 0; i < 5000000; i++ {
		if seq[i] != seq1[i] {
			fmt.Println("false")
			return
		}
	}
	fmt.Println(true)
}

// 比较归并1与归并2
func Test13(t *testing.T) {

	seq1 := make([]int, 20000000)
	seq := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		v := rand.Intn(100000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	//1000000	106.561959m
	//5000000	527.731542ms
	//10000000 	1.114845458s
	//20000000	2.327849459s
	MergeSort1(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	//1000000	100.36525ms
	//5000000	532.446542ms
	//10000000 	1.124789417s
	//20000000	2.3209445s
	MergeSort2(seq1)
	t2 = time.Since(t1)
	fmt.Println(t2)

	for i := 0; i < 20000000; i++ {
		if seq[i] != seq1[i] {
			fmt.Println("false")
			return
		}
	}
	fmt.Println(true)
}

// 比较归并1与快速
func Test12(t *testing.T) {

	seq1 := make([]int, 1000000)
	seq := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		v := rand.Intn(1000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	//1000000	106.561959m
	//5000000	541.86625ms
	//10000000 	1.114845458s
	//20000000	2.327849459s
	seq = MergeSort1(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	//1000000	65.122542ms
	//5000000	362.930584ms
	//10000000 	710.020041ms
	//20000000	1.52779375s
	QuickSort2(seq1)
	t2 = time.Since(t1)
	fmt.Println(t2)

	for i := 0; i < 1000000; i++ {
		if seq[i] != seq1[i] {
			fmt.Println("false")
			return
		}
	}
	fmt.Println(true)
}
