package dp

/*
给你一个整数数组 nums ，你可以对它进行一些操作。
每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。

*/
func DeleteAndEarn(nums []int) int {
	//a:将nums中的元素值与a的下标进行映射
	//元素值出现,a对应的下标中的值+1
	a := make([]int, 10001)
	for _, v := range nums {
		a[v] += 1
	}
	//将a的下标与其元素值相乘
	//得到每个下标对应的点数
	//相邻下标不能都选.
	b := make([]int, 10001)
	for i, v := range a {
		b[i] = i * v
	}
	return Rob(b)

}
