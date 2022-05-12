/*
 * @Author: zwngkey
 * @Date: 2022-04-22 06:17:17
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:07:47
 * @Description:
 */
package dp

import (
	"zwngkey.cn/dsaa/util"
)

func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	//f[i]表示爬到第i个台阶所需的最小花费
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 0
	for i := 2; i <= n; i++ {
		f[i] = util.Min(f[i-1]+cost[i-1], f[i-2]+cost[i-2])
	}
	return f[n]
}
