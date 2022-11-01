/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-31 21:12:22
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-31 21:21:08
 * @Description:
	版本2
*/
package bt

import (
	"fmt"
	"testing"
)

func combine1(n, k int) (result [][]int) {
	var bt func(int)
	var path []int
	bt = func(cur int) {
		// 剪枝：path 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 path
		if len(path)+(n-cur+1) < k {
			return
		}
		// 记录合法答案
		if len(path) == k {
			path := make([]int, k)
			copy(path, path)
			result = append(result, path)
			return
		}
		// 考虑选择当前位置
		path = append(path, cur)
		bt(cur + 1)
		path = path[:len(path)-1]
		// 考虑不选择当前位置
		bt(cur + 1)
	}
	bt(1)
	return result
}

func Test1(t *testing.T) {
	res := combine1(5, 4)
	fmt.Println(res)
}
