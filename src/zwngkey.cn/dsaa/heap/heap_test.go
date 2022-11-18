/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 07:22:36
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-18 07:46:04
 * @Description:

 */
package heap

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	seq := []int{3, 1, 2, 5, 7, 9, 22, 8, 12, 0}
	hp := NewHeapByArray(seq)
	hp.Push(-1)

	fmt.Println(hp.Pop()) // -1
	fmt.Println(hp.Pop()) // 0
	fmt.Println(hp.Pop()) // 1

	for _, v := range *hp {
		fmt.Println(v)
	}
	hp.Remove(2)
	for _, v := range *hp {
		fmt.Println(v)
	}

	(*hp)[0] = 10
	hp.Fix(0)
	for _, v := range *hp {
		fmt.Println(v)
	}
}
