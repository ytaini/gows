/*
- @Author: wzmiiiiii

- @Date: 2022-11-03 17:45:00

  - @LastEditors: wzmiiiiii

  - @LastEditTime: 2022-12-07 18:23:19

  - @Description:
    冒泡排序:
    思想:对序列,从前向后(下标从小->大),依次比较相邻元素的大小.若逆序则交换位置.使靠后的元素为越来越大(或越来越小)

    优化: 当某一次循环中,未发生任何交换.说明序列已经排好序了.
    通过一个标志位flag.
*/
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 4,3,-1,5,6,8,2,10,3
func bubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		flag := false
		for j := 0; j < len(s)-i-1; j++ {
			// 升序
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				flag = true
			}
			//降序
			// if s[j] < s[j+1] {
			// 	s[j], s[j+1] = s[j+1], s[j]
			// 	flag = true
			// }
		}
		if !flag {
			break
		}
	}
}

func Test(t *testing.T) {
	// seq := []int{4, 3, -1, 5, 6, 8, 2, 10, 3}
	// seq := []int{1, 2, 3, 4, 3, 6}
	seq := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		seq[i] = rand.Intn(10000000)
	}
	// 50000  3.012327125s
	// 100000 12.634303625s
	// 200000 >30s
	t1 := time.Now()
	bubbleSort(seq)
	t2 := time.Since(t1)
	fmt.Println(t2.Seconds())
}
