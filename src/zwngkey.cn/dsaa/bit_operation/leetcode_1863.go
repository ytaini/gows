/*
 * @Author: zwngkey
 * @Date: 2022-04-22 07:38:52
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 05:50:35
 * @Description:
 */
package bitoper

func SubsetXORSum(nums []int) int {
	var sum int
	var temp int
	length := len(nums)
	for i := 1; i < (1 << length); i++ {
		temp = 0
		for j := 0; j < length; j++ {
			//if  (i >> j) & 1 == 1
			if (i & (1 << j)) > 0 {
				temp ^= nums[j]
			}
		}
		sum += temp
	}
	return sum

}
func SubsetXORSum1(nums []int) int {

	result := 0

	var dfs func(int, int)

	dfs = func(sum, i int) {
		if i >= len(nums) {
			result += sum
			return
		}
		dfs(sum, i+1)
		dfs(sum^nums[i], i+1)
	}

	dfs(0, 0)

	return result
}
