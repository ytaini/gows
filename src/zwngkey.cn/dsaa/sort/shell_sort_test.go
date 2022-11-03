/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-03 20:45:36
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-04 02:06:37
 * @Description:
	插入排序可能存在的问题:
		当 需要插入的数是较小的数时,后移的次数明显增多,对效率有影响.

	希尔排序: 也是一种插入排序.也称缩小增量排序.

		思想:
			把记录按下标的一定增量分组.对每组使用插入排序.
				随着增量逐渐减少,每组包含的关键词越来越多,当增量减至1时,整个序列恰好被分为一组,算法便终止.
*/
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
	希尔排序在数组中采用跳跃式分组的策略，通过某个增量将数组元素划分为若干组，然后分组进行插入排序，随后逐步缩小增量，继续按组进行插入排序操作，直至增量为1。
	希尔排序通过这种策略使得整个数组在初始阶段达到从宏观上看基本有序，小的基本在前，大的基本在后。然后缩小增量，到增量为1时，其实多数情况下只需微调即可，不会涉及过多的数据移动。

	增量gap=length/2，缩小增量继续以gap = gap/2的方式，这种增量选择我们可以用一个序列来表示，{n/2,(n/2)/2...1}，称为增量序列。
	希尔排序的增量序列的选择与证明是个数学难题，我们选择的这个增量序列是比较常用的，也是希尔建议的增量，称为希尔增量，但其实这个增量序列不是最优的。
*/

func shellSort(s []int) {
	for gap := len(s) / 2; gap >= 1; gap /= 2 {
		for i := gap; i < len(s); i++ { // 这里会交叉给每个分组排序
			for j := i - gap; j >= 0 && s[j] > s[j+gap]; j -= gap { //对分组进行冒泡排序
				// 希尔排序时,对有序序列在插入时,采用交换法.
				s[j], s[j+gap] = s[j+gap], s[j]
			}
		}
	}
}

// 希尔排序时,对有序序列在插入时,采用移动法.
func shellSort1(s []int) {
	for gap := len(s) / 2; gap >= 1; gap /= 2 {
		for i := gap; i < len(s); i++ { // 这里会交叉给每个分组排序
			// 对分组进行插入排序
			j := i
			val := s[j]
			if val < s[j-gap] {
				for j-gap >= 0 && val < s[j-gap] {
					s[j] = s[j-gap]
					j -= gap
				}
				s[j] = val
			}
		}
	}
}

// 比较移动法与交换法.
func Test6(t *testing.T) {
	seq1 := make([]int, 10000000)
	seq := make([]int, 10000000)
	for i := 0; i < 10000000; i++ {
		v := rand.Intn(100000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	//	1000000	 178.631417ms
	//	10000000 2.200363625s
	shellSort(seq)
	t2 := time.Since(t1)
	fmt.Println(t2)

	t1 = time.Now()
	//	1000000 127.412959ms
	//	10000000 1.517237458s
	shellSort1(seq1)
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

// 测试希尔排序的交换法.
func Test5(t *testing.T) {
	// seq := []int{8, 9, 1, 7, 2, 3, 5, 4, 5, 0}
	// seq := []int{1, 2, 3, 4, 3, 6}
	// shellSort(seq)
	// fmt.Println(seq)
	seq := make([]int, 10000000)
	for i := 0; i < 10000000; i++ {
		seq[i] = rand.Intn(1000000000)
	}
	// 50000	0.013364792
	// 100000  	0.020024458
	// 500000 	0.091819416
	// 1000000 	0.187775
	// 5000000	1.056274584
	// 10000000 2.352806959
	t1 := time.Now()
	shellSort(seq)
	t2 := time.Since(t1)
	fmt.Println(t2.Seconds())
}

// 对比insertSort 与shell排序.
func Test4(t *testing.T) {
	seq1 := make([]int, 1000000)
	seq := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		v := rand.Intn(100000000)
		seq[i] = v
		seq1[i] = v
	}

	t1 := time.Now()
	shellSort(seq)
	t2 := time.Since(t1)
	fmt.Println(t2.Seconds())

	t1 = time.Now()
	InsertSort(seq1)
	t2 = time.Since(t1)
	fmt.Println(t2.Seconds())

	for i := 0; i < 1000000; i++ {
		if seq[i] != seq1[i] {
			fmt.Println("false")
			return
		}
	}
	fmt.Println(true)
}
