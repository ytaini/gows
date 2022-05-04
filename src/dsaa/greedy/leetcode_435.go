package greedy

import (
	"fmt"
	"sort"
)

// 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重叠 。

// 题目的要求等价于「选出最多数量的区间，使得它们互不重叠」

// 我们不妨想一想应该选择哪一个区间作为首个区间。
// 假设在某一种最优的选择方法中，[Lk,Rk]是首个（即最左侧的）区间，那么它的左侧没有其它区间，右侧有若干个不重叠的区间。
// 设想一下，如果此时存在一个区间 [Lj,Rj]使得 Rj < Rk，即区间 j 的右端点在区间 k 的左侧，
// 那么我们将区间 k 替换为区间 j，其与剩余右侧被选择的区间仍然是不重叠的。
// 而当我们将区间 k 替换为区间 j 后，就得到了另一种最优的选择方法。

// 我们可以不断地寻找右端点在首个区间右端点左侧的新区间，将首个区间替换成该区间。
// 那么当我们无法替换时，首个区间就是所有可以选择的区间中右端点最小的那个区间。
// 因此我们将所有区间按照右端点从小到大进行排序，那么排完序之后的首个区间，就是我们选择的首个区间。

// 如果有多个区间的右端点都同样最小怎么办？
// 由于我们选择的是首个区间，因此在左侧不会有其它的区间，那么左端点在何处是不重要的，我们只要任意选择一个右端点最小的区间即可。

// 当确定了首个区间之后，所有与首个区间不重合的区间就组成了一个规模更小的子问题。
// 由于我们已经在初始时将所有区间按照右端点排好序了，因此对于这个子问题，我们无需再次进行排序，
// 只要找出其中与首个区间不重合并且右端点最小的区间即可。
// 用相同的方法，我们可以依次确定后续的所有区间。

// 在实际的代码编写中，我们对按照右端点排好序的区间进行遍历，
// 并且实时维护上一个选择区间的右端点 right。
// 如果当前遍历到的区间 [ Li , Ri ]与上一个区间不重合，即 Li ≥ right，
// 那么我们就可以贪心地选择这个区间，并将 right 更新为 Ri
func EraseOverlapIntervals(intervals [][]int) int {
	/*
		要选出最多数量的区间时,每个区间的结尾十分重要.
		选择区间i时,应该先选择右端点最小的区间.
		这样余留给其他区间的空间就越大,就能保留更多的区间.

		贪心选择策略:
		每次选择右端点最小的区间,且这个区间的左端点要比上一个选择的区间的右端点小.

	*/

	//按照区间右边界来排序.
	//升序
	// 按照右边界排序，就要从左向右遍历，因为右边界越小越好，只要右边界越小，留给下一个区间的空间就越大，所以从左向右遍历，优先选右边界小的。
	// 按照左边界排序，就要从右向左遍历，因为左边界数值越大越好（越靠右），这样就给前一个区间的空间就越大，所以可以从右向左遍历。

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	fmt.Printf("intervals: %v\n", intervals)
	n := len(intervals)

	right := intervals[0][1]
	//保存保留的区间个数.
	//因为最少有一个区间,所以保留的区间个数初始值为1
	cnt := 1
	for _, v := range intervals[1:] {
		//当第i个区间的左端点大于已保留的最后一个区间的右端点时,保留的区间个数+1,已保留的最后一个区间改变.
		if v[0] >= right {
			cnt++
			right = v[1]
		}
	}
	//区间总数-最多可保留的区间数=最少要删除的区间数
	return n - cnt
}

func EraseOverlapIntervals1(intervals [][]int) int {

	//按照区间开头来排序.
	//升序
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Printf("intervals: %v\n", intervals)
	n := len(intervals)
	right := n - 1
	cnt := 1
	for i := n - 2; i >= 0; i-- {
		if intervals[i][1] <= intervals[right][0] {
			right = i
			cnt++
		}
	}
	return n - cnt
}

func EraseOverlapIntervals2(intervals [][]int) int {

	//按照区间开头来排序.
	//降序
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] > intervals[j][0]
	})
	fmt.Printf("intervals: %v\n", intervals)
	n := len(intervals)
	right := intervals[0][0]
	cnt := 1
	for _, v := range intervals[1:] {
		if v[1] <= right {
			right = v[0]
			cnt++
		}
	}
	return n - cnt
}
