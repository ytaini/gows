/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-03 18:46:56
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-03 20:18:17
 * @Description:
	选择排序:
		思想: 第一次从arr[0:n]中选取最小值.与arr[0]交换,第二次从arr[1:n]中选取最小值与arr[1]交换....
			依次类推,共通过n-1次,得到有序序列.
*/
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func selectSort(s []int) {
	for i := 0; i < len(s)-1; i++ { //寻找n-1次最小值
		k := i                            //假设第i个为最小的.
		for j := i + 1; j < len(s); j++ { //依次遍历:[0,n],[1,n]...[n-1,n]
			// 升序
			if s[j] < s[k] {
				k = j
			}
			// 降序
			// if s[j] > s[k] {
			// 	k = j
			// }
		}
		if k != i { //第i个就是最小的时,就不用交换了.
			s[i], s[k] = s[k], s[i]
		}
	}
}
func Test1(t *testing.T) {
	// seq := []int{4, 3, -1, 5, 6, 8, 2, 10, 3}
	// seq := []int{1, 2, 3, 4, 3, 6}
	// selectSort(seq)
	// fmt.Println(seq)

	seq := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		seq[i] = rand.Intn(10000000)
	}
	// 50000  0.876472583s
	// 100000 3.785873667s
	// 200000 15.088780375s
	t1 := time.Now()
	selectSort(seq)
	t2 := time.Since(t1)
	fmt.Println(t2.Seconds())
}
