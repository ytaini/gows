package dp

import (
	"dsaa/util"
)

/*
给定一个整数n(1<=n<=100),再给定一个n个整数的数组nums,每个整数可以选择取或不取,
如果取下标为i的数,那么下标为i-1,i+1的数就不能取.
按照上述规则选取一些整数,使得选出来的整数得到的总和最大,返回这个最大值.
*/
func Rob(nums []int) int {
	n := len(nums)
	//f[i]:表示前i+1个整数,通过某种方案能够获取的最大值.
	f := make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < n; i++ {
		if i == 1 {
			f[i] = util.Max(nums[0], nums[1])
		} else {
			f[i] = util.Max(f[i-1], f[i-2]+nums[i])
		}
	}
	return f[n-1]
}

// 上述方法使用了数组存储结果。
// 考虑到每间房屋的最高总金额只和该房屋的前两间房屋的最高总金额相关，
// 因此可以使用滚动数组，在每个时刻只需要存储前两间房屋的最高总金额。
func Rob1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	first := nums[0]
	second := util.Max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		first, second = second, util.Max(first+nums[i], second)
	}
	return second
}
