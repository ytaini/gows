package nt

import (
	"dsaa/util"
	"sort"
)

func FindGCD(nums []int) int {
	sort.Ints(nums)
	return util.GCDv2(nums[0], nums[len(nums)-1])
}
