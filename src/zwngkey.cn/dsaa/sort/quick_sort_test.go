/*
- @Author: wzmiiiiii

- @Date: 2022-11-04 00:46:05

  - @LastEditors: wzmiiiiii

  - @LastEditTime: 2022-12-07 18:27:06

  - @Description:
    快速排序（Quicksort）是对冒泡排序的一种改进，借用了分治思想，选基准，依次将剩余元素的小于基准放其左侧，大的放右侧
    然后取基准的前半部分和后半部分进行同样处理，直至各子序列剩余一个元素结束，排序完成

    快速排序，基于比较，不稳定算法，时间平均O(nlogn)，最坏O(n^2)，空间O(logn)

    思想:
    通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，
    然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列

    1.在待排序的元素任取一个元素作为基准（通常选第一个元素，但最好的方法是从待排序元素中随机选取一个为基准），称为基准元素（pivot）
    2.将待排序的元素进行分区，比基准元素大的元素放在它的右边，比基准元素小的放在它的左边
    3.对左右两个分区重复以上步骤，直到所有的元素都是有序的

    注意：
    小规模数据(n<100)，由于快排用到递归，性能不如插排
    进行排序时，可定义阈值，小规模数据用插排，往后用快排
    golang的sort包用到了快排
*/
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 升序
func QuickSort1(seq []int) { // 第三快
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

func QuickSort2(seq []int) { // 第一快
	var QuickSort func(seq []int, left, right int)

	QuickSort = func(seq []int, left, right int) {
		//递归结束条件
		if left > right {
			return
		}
		// 分割并得到基准元素的位置
		pivotIndex := partition2(seq, left, right)
		QuickSort(seq, left, pivotIndex-1)
		QuickSort(seq, pivotIndex+1, right)
	}
	QuickSort(seq, 0, len(seq)-1)
}

func partition2(seq []int, left, right int) int {
	pivot := seq[left]  //取第一个元素作为基准元素
	l, r := left, right //左右指针.
	for r > l {
		// 从左往右搜索一个比pivot小的值.
		for r > l {
			if seq[r] < pivot {
				seq[l] = seq[r] // 将此值放在pivot的左边
				l++             // 左指针右移.
				break           // 找到了就退出此for.
			}
			r-- //没找到就一直往右找,直到r<l
		}
		// 从右往左搜索一个比pivot大的值.
		for r > l {
			if seq[l] > pivot {
				seq[r] = seq[l] // 将此值放在pivot的右边
				r--             //右指针左移.
				break
			}
			l++ //没找到就一直往左找,直到r<l
		}
	}
	seq[r] = pivot // 此时pivot的位置就为r
	return r
}

func QuickSort3(seq []int) { // 第二快
	var QuickSort func(seq []int, left, right int)

	QuickSort = func(seq []int, left, right int) {
		if left > right {
			return
		}
		pivotIndex := partition3(seq, left, right)
		QuickSort(seq, left, pivotIndex-1)
		QuickSort(seq, pivotIndex+1, right)
	}
	QuickSort(seq, 0, len(seq)-1)
}

func partition3(seq []int, left, right int) int {
	pivot := left
	index := pivot + 1 //记录基准位置
	for i := index; i <= right; i++ {
		if seq[i] < seq[pivot] {
			seq[i], seq[index] = seq[index], seq[i]
			index++
		}
	}
	seq[pivot], seq[index-1] = seq[index-1], seq[pivot]
	return index - 1
}

func Test8(t *testing.T) {
	// seq := []int{10, 2, 8, 22, 34, 5, 12, 28, 21, 11,4, 7, 6, 5, 3, 2, 8, 1,}
	seq := make([]int, 100)
	for i := 0; i < 100; i++ {
		seq[i] = rand.Intn(80)
	}

	QuickSort2(seq)
	fmt.Println(seq)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func quickSort1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := nums[0]
	nums = append(nums[1:], pivot)

	var less, greater []int

	for _, num := range nums[:len(nums)-1] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	return append(append(quickSort1(less), pivot), quickSort1(greater)...)
}

func Test101(t *testing.T) {
	arr := []int{5, 8, 2, 1, 6, 9, 4, 7, 3}
	fmt.Println(quickSort(arr)) // [1 2 3 4 5 6 7 8 9]
}

func Test20(t *testing.T) {

	seq1 := make([]int, 10000000)
	seq := make([]int, 10000000)
	for i := 0; i < 10000000; i++ {
		v := rand.Intn(100000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	quickSort(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	quickSort1(seq1)
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

// 比较快排2与快排3
func Test10(t *testing.T) {

	seq1 := make([]int, 20000000)
	seq := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		v := rand.Intn(100000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	//1000000	72.449583ms
	//5000000	363.460375ms
	//10000000 	751.975542ms
	//20000000	1.5665955s
	QuickSort3(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	//1000000	65.122542ms
	//5000000	362.930584ms
	//10000000 	736.2945ms
	//20000000	1.52779375s
	QuickSort2(seq1)
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

// 比较快排1与快排2
func Test9(t *testing.T) {

	seq1 := make([]int, 20000000)
	seq := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		v := rand.Intn(100000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	//1000000	72.449583ms
	//5000000	367.725375ms
	//10000000 	757.364291ms
	//20000000	1.59018475ss
	QuickSort1(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	//1000000	65.122542ms
	//5000000	362.930584ms
	//10000000 	736.2945ms
	//20000000	1.52674925s
	QuickSort2(seq1)
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

// 比较快排1与希尔
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
	QuickSort1(seq)
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
