/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 14:05:56
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-19 04:26:57
 * @Description:
	53. 最大子数组和
		给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

		子数组 是数组中的一个连续部分。
*/
package divideandconquer

import "zwngkey.cn/dsaa/util"

// 我们定义一个操作 get(a, l, r) 表示查询 a 序列 [l,r] 区间内的最大子段和，
// 	那么最终我们要求的答案就是 get(nums, 0, len(nums) - 1)
// 如何分治实现这个操作呢？
// 对于一个区间 [l,r]，我们取 m =(l+r)/2 对区间 [l,m] 和 [m+1,r] 分治求解。
// 当递归逐层深入直到区间长度缩小为 1 的时候，递归「开始回升」.
// 这个时候我们考虑如何通过 [l,m] 区间的信息和 [m+1,r] 区间的信息合并成区间 [l,r] 的信息。最关键的两个问题是：
//	我们要维护区间的哪些信息呢？
//	我们如何合并这些信息呢？
// 对于一个区间 [l,r]，我们可以维护四个量:
// 	- lSum 表示 [l,r] 内以 l 为左端点的最大子段和
// 	- rSum 表示 [l,r] 内以 r 为右端点的最大子段和
// 	- mSum 表示 [l,r] 内的最大子段和
//	- iSum 表示 [l,r] 的区间和
// 以下简称 [l,m] 为 [l,r]的「左子区间」，[m+1,r] 为 [l,r]的「右子区间」。
// 我们考虑如何维护这些量呢（如何通过左右子区间的信息合并得到 [l,r] 的信息）？
// 对于长度为 1 的区间 [i, i]，四个量的值都和 nums[i] 相等。
// 对于长度大于 1 的区间：
// 	- 首先最好维护的是 iSum，区间 [l,r]的 iSum 就等于「左子区间」的iSum 加上「右子区间」的 iSum。
// 	- 对于 [l,r] 的 lSum，存在两种可能，它要么等于「左子区间」的 lSum，要么等于「左子区间」的 iSum 加上「右子区间」的 lSum，二者取大。
// 	- 对于 [l,r] 的 rSum，同理，它要么等于「右子区间」的 rSum，要么等于「右子区间」的 iSum 加上「左子区间」的 rSum，二者取大。
// 	- 当计算好上面的三个量之后，就很好计算 [l,r] 的 mSum 了。
// 		我们可以考虑 [l,r] 的 mSum 对应的区间是否跨越 m
// 		  	- 它可能不跨越 m，也就是说 [l,r] 的 mSum 可能是「左子区间」的 mSum 和 「右子区间」的 mSum 中的一个；
// 			- 它也可能跨越 m，可能是「左子区间」的 rSum 和 「右子区间」的 lSum 求和。三者取大。
// 这样问题就得到了解决。

func MaxSubArray(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) / 2
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := util.Max(l.lSum, l.iSum+r.lSum)
	rSum := util.Max(r.rSum, r.iSum+l.rSum)
	mSum := util.Max(util.Max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

type Status struct {
	lSum, rSum, mSum, iSum int
}
