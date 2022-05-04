package dp

import (
	"dsaa/util"
	"fmt"
	"sort"
)

//给n个区间,问:最多可以选择多少个不重叠的区间.
func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	fmt.Printf("intervals: %v\n", intervals)
	n := len(intervals)
	// 状态: dp[i] 表示以区间i为最后一个区间,可以选出的区间数量的最大值

	// 理论上第i个区间可以与前面的任何一个或多个区间构成不重叠的区间.
	// 所有要分析第i个区间与它前面的所有区间组合的情况.并记录此时的最大不重叠的区间数.
	// 当前面某一个区间j的右端点<=第i个区间的左端点时,则认为区间j与第i个区间不重叠.

	dp := make([]int, n)

	// 初始化 : 以任意区间k结尾,它能选出的区间数量至少为1.即区间k自身.
	for i := range dp {
		dp[i] = 1
	}
	//计算以区间i为最后一个区间,可以选出的最大不重叠的区间数.
	for i := 1; i < n; i++ {
		//判断区间i与它前面的任何一个或多个区间是否重叠
		for j := 0; j < i; j++ {
			//如果区间j的右端点<=第i个区间的左端点时
			//则区间j与区间i不重叠.
			//我们在所有满足要求的 区间j中，选择 DP[j] ​ 最大的那一个
			//即 max{dp[i] ,dp[j]+1}.(+1:加上区间i,最大不重叠区间数+1.)

			//因为j<i
			//此时问题从求dp[i],变为了求dp[j],即将原问题dp[i]分为了规模更小的子问题dp[j].
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = util.Max(dp[i], dp[j]+1)
			}
		}
	}
	//我们在所有满足要求的 dp[i] 中，选择 dp[i]最大的那一个
	return n - util.MaxIntSlice(dp...)
}
