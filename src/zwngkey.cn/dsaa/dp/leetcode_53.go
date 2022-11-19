/*
 * @Author: zwngkey
 * @Date: 2022-04-22 07:36:29
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-19 04:01:57
 * @Description:

	给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	子数组 是数组中的一个连续部分。
*/
package dp

import (
	"zwngkey.cn/dsaa/util"
)

func MaxSubArray(nums []int) int {
	n := len(nums)
	max := nums[n-1]

	//f[i]:表示以i开始的所有子数组的和中的最大值.
	f := make([]int, n)
	f[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		f[i] = util.Max(f[i+1]+nums[i], nums[i])
		max = util.Max(max, f[i])
	}
	return max
}

func MaxSubArray1(nums []int) int {
	n := len(nums)
	maxSum := nums[0] //初始最大和
	//f[i]:表示以nums[i]结尾的连续子数组的最大和.
	f := make([]int, n)
	f[0] = nums[0] //初始条件
	for i := 1; i < n; i++ {
		f[i] = util.Max(f[i-1]+nums[i], nums[i]) //状态转移方程
		maxSum = util.Max(maxSum, f[i])          //更新最大和
	}
	return maxSum
}

// 贪心
func MaxSubArray2(nums []int) int {
	maxSum := nums[0]
	cusSum := 0
	for _, v := range nums {
		//if cusSum > 0 {
		if cusSum+v > v { //这步操作就是 cusSum 在v 与 cusSum+v 中取较大值.
			cusSum += v
		} else {
			cusSum = v
		}
		// cusSum = util.Max(v, cusSum+v)  上面可以简写成这样.
		maxSum = util.Max(cusSum, maxSum)
	}
	return maxSum
}
