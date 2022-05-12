/*
 * @Author: zwngkey
 * @Date: 2022-04-24 05:05:33
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:08:12
 * @Description:
 */
package nt

import (
	"sort"
	"zwngkey.cn/dsaa/util"
)

func FindGCD(nums []int) int {
	sort.Ints(nums)
	return util.GCDv2(nums[0], nums[len(nums)-1])
}
