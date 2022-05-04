package dp

import (
	"dsaa/util"
)

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
*/
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
	max := nums[0]
	//f[i]:表示以i结尾的所有子数组的和中的最大值.
	f := make([]int, n)
	f[0] = nums[0]
	for i := 0; i < n; i++ {
		f[i] = util.Max(f[i-1]+nums[i], nums[i])
		max = util.Max(max, f[i])
	}
	return max
}
