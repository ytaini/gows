/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-03 19:26:06
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-03 23:53:55
 * @Description:
	插入排序:
	  思想: 把n个待排序的元素看成一个有序表和无序表,开始时有序表中只包含一个元素,无序表中包含n-1个元素,
	  	排序过程中每次从无序表中取出第一个元素.把它依次与有序表元素进行比较,将他插入到有序表中的适当位置.使之成为新的有序表.
*/
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
	1.将s中第一个元素 当做 一个有序序列.
	2.依次将后面的元素,加入到这个有序序列中.
*/

// 升序.
func InsertSort(s []int) {
	for i := 1; i < len(s); i++ {
		val := s[i] // 记录待排序元素的值..
		k := i - 1  // 指向有序序列最后一个元素的下标.

		// 搜索val存放的正确位置.
		// for k >= 0 && val > s[k] { //降序
		if val < s[k] { //只有待排序的元素小于有序序列最后一个元素时,才需要移动.
			for k >= 0 && val < s[k] {
				s[k+1] = s[k] //将排好序的元素后移.
				k--           //指向前一个元素.
			}
			// 到这里说明第k位的值 < val
			// 所有需要插在k+1位.
			s[k+1] = val
		}
	}
}

func Test3(t *testing.T) {
	// seq := []int{4, 3, -1, 5, 6, 8, 2, 10, 3}
	// seq := []int{1, 2, 3, 4, 3, 6}
	// InsertSort(seq)
	// fmt.Println(seq)

	seq := make([]int, 400000)
	for i := 0; i < 400000; i++ {
		seq[i] = rand.Intn(10000000)
	}
	// 50000 0.33568s
	// 100000 1.229207042s
	// 200000 4.841481708s
	// 300000 10.789640458000001s
	// 400000 19.146305792s
	// 1000000 123.168497208s
	t1 := time.Now()
	InsertSort(seq)
	t2 := time.Since(t1)
	fmt.Println(t2.Seconds())
}
