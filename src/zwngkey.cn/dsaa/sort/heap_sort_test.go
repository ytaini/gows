/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-14 17:16:39
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-15 01:40:23
 * @Description:
	二叉堆: 二叉堆本质上是一种完全二叉树，它分为两个类型
		1. 大顶堆：大顶堆任何一个父节点，都大于或等于它左、右孩子节点的值
		2. 小顶堆：小顶堆的任何一个父节点，都小于或等于它左、右孩子节点的值
	  二叉堆的根节点叫作堆顶。最大堆和最小堆的特点决定了：最大堆的堆顶是整个堆中的最大元素；最小堆的堆顶事整个堆中的最小元素。

	在代码实现上，由于堆是一个完全二叉树，所以在存储上可以使用数组实现，假设父节点的下标为 parent，那么它的左孩子下标就是 2 * parent+1，右孩子下标就是 2 * parent+2

	堆排序:
		堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。基于比较交换的不稳定算法，平均时间复杂度为 Ο(nlogn)。，空间O(1)
		堆是一个本质上是完全二叉树的结构，并同时满足堆的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。
			大顶堆：每个节点的值都大于或等于其子节点的值，在堆排序算法中用于升序排列；
			小顶堆：每个节点的值都小于或等于其子节点的值，在堆排序算法中用于降序排列；

	主要思路：
		1.建堆，从最后一个非叶子节点开始依次堆化,假设有 n 个结点,最后一个非叶子节点的下标为(n/2 - 1).注意逆序，从下往上堆化
		建堆流程：父节点与子节点比较，子节点大则交换父子节点，父节点索引更新为子节点，循环操作
		2.尾部遍历操作，弹出元素，再次堆化
		弹出元素排序流程：从最后节点开始，交换头尾元素，由于弹出，arrlen--，再次对剩余数组元素建堆，循环操作

*/
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func HeapSort(arr []int) []int {
	s := len(arr)
	// 构建大顶堆
	buildMaxHeap(arr)
	// 依次交换元素，然后再堆化，相当于依次把最大值放入尾部
	for k := s - 1; k >= 0; k-- {
		arr[0], arr[k] = arr[k], arr[0]
		s--
		Heapify1(arr, 0, s)
	}
	return arr
}

// 构建大顶堆
func buildMaxHeap(arr []int) {
	s := len(arr)
	i := (s - 1)              //最后一个叶子结点的下标
	j := (i - 1) / 2          //最后一个叶子结点的父节点的下标.
	for k := j; k >= 0; k-- { // 从最后一个非叶子节点开始堆化
		Heapify1(arr, k, s)
	}
}

// i: 非叶子结点在数组中的索引
// 递归堆化函数(大顶堆)
func Heapify(tree []int, i, arrlen int) { //慢
	if i >= arrlen {
		return
	}
	l := i*2 + 1 //左孩子
	r := l + 1   //右孩子
	max := i     //记录最大值
	if l < arrlen && tree[l] > tree[max] {
		max = l
	}
	if r < arrlen && tree[r] > tree[max] {
		max = r
	}
	if max != i {
		tree[i], tree[max] = tree[max], tree[i]
		Heapify(tree, max, arrlen)
	}
}

// i: 非叶子结点在数组中的索引
// 非递归堆化函数(小顶堆)
func Heapify1(tree []int, i, arrlen int) { //快
	// 大顶堆堆化，堆顶值小一直下沉
	for {
		child := i*2 + 1 // 左孩子节点索引
		if child >= arrlen {
			break
		}
		if child+1 < arrlen && tree[child] > tree[child+1] { //比较左右孩子，取大值，否则child不用++
			child++
		}
		if tree[i] < tree[child] { // 如果父节点已经大于左右孩子大值，已堆化
			break
		}
		// 孩子节点大值上冒
		tree[i], tree[child] = tree[child], tree[i]
		i = child // 更新父节点到子节点，继续往下比较，不断下沉
	}
}

func Test21(t *testing.T) {
	arr := []int{8, 10, 2, 3, 1, 9, 7, 5, 4, 6}
	HeapSort(arr)
	fmt.Println(arr)
}

// 比较快排2与 堆排序
func Test22(t *testing.T) {

	seq1 := make([]int, 20000000)
	seq := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		v := rand.Intn(100000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	//1000000	70.109833ms
	//5000000	344.276959ms
	//10000000 	702.800041ms
	//20000000	1.59018475ss
	QuickSort2(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	//1000000	91.283166ms
	//5000000	787.524167ms
	//10000000 	2.094627416s
	//20000000	5.01153475s
	HeapSort(seq1)
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