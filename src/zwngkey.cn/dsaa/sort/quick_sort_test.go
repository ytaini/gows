/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-04 00:46:05
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-04 03:13:56
 * @Description:
	快速排序（Quicksort）是对冒泡排序的一种改进，借用了分治的思想

	思想:
		通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，
			然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列

	1.在待排序的元素任取一个元素作为基准（通常选第一个元素，但最好的方法是从待排序元素中随机选取一个为基准），称为基准元素（pivot）
	2.将待排序的元素进行分区，比基准元素大的元素放在它的右边，比基准元素小的放在它的左边
	3.对左右两个分区重复以上步骤，直到所有的元素都是有序的

*/
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 快速排序，基于比较，不稳定算法，时间平均O(nlogn)，最坏O(n^2)，空间O(logn)
// 分治思想，选基准，依次将剩余元素的小于基准放其左侧，大的放右侧
// 然后取基准的前半部分和后半部分进行同样处理，直至各子序列剩余一个元素结束，排序完成
// 注意：
// 小规模数据(n<100)，由于快排用到递归，性能不如插排
// 进行排序时，可定义阈值，小规模数据用插排，往后用快排
// golang的sort包用到了快排
// 升序
func quickSort(seq []int) {
	var quick func(seq []int, left, right int)

	quick = func(seq []int, left, right int) {
		if left > right {
			return
		}
		// 左右指针,基准
		l, r, pivot := left, right, seq[right]

		// 将小于pivot放其左侧，大的放右侧
		for l < r {
			for l < r && seq[l] <= pivot { //在pivot 左边找到了一个>pivot的下标.
				l++
			}
			for r > l && seq[r] >= pivot { //在pivot 右边找到了一个<pivot的下标.
				r--
			}
			seq[r], seq[l] = seq[l], seq[r] //交换它们的值
		}
		seq[r], seq[right] = seq[right], seq[r] //将基准放到中间
		quick(seq, left, r-1)                   //继续按同样的方式处理左半边
		quick(seq, r+1, right)                  //继续按同样的方式处理右半边
	}
	quick(seq, 0, len(seq)-1)
}

func quickSort1(seq []int) {
	var partition = func(seq []int, left, right int) int {
		l, r, pivot, index := left, right, seq[left], left
		for r >= l {
			for r >= l {
				if seq[r] < pivot {
					seq[l] = seq[r]
					index = r
					l++
					break
				}
				r--
			}
			for r >= l {
				if seq[l] > pivot {
					seq[r] = seq[l]
					index = l
					r--
					break
				}
				l++
			}
		}
		seq[index] = pivot
		return index
	}
	var quickSort func(seq []int, left, right int)

	quickSort = func(seq []int, left, right int) {
		if left > right {
			return
		}
		pivotIndex := partition(seq, left, right)
		quickSort(seq, left, pivotIndex-1)
		quickSort(seq, pivotIndex+1, right)
	}
	quickSort(seq, 0, len(seq)-1)
}

func Test8(t *testing.T) {
	seq := []int{2, 10, 8, 22, 34, 5, 12, 28, 21, 11}
	// seq := []int{4, 7, 6, 5, 3, 2, 8, 1}
	quickSort1(seq)
	fmt.Println(seq)
}

// 比较快排与希尔
func Test7(t *testing.T) {

	seq1 := make([]int, 20000000)
	seq := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		v := rand.Intn(1000000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	//1000000	72.449583ms
	//5000000	370.094083ms
	//10000000 	777.220208ms
	//20000000	1.59018475ss
	// quickSort(seq)
	quickSort1(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	//1000000	127.113917ms
	//5000000	747.176875ms
	//10000000 	1.692346166s
	//20000000	3.760462583s
	shellSort1(seq1)
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
