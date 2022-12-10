package lc01

// TwoSum2 哈希表 妙!!!
// 创建一个哈希表，对于nums中每一个值nums[i]，首先查询哈希表中是否存在 target - nums[i]，
// 然后将 nums[i] 插入到哈希表中，即可保证不会让 x 和自己匹配。
func TwoSum2(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		if t, ok := mp[target-v]; ok {
			return []int{t, i}
		}
		mp[v] = i
	}
	return nil
}

// TwoSum3 哈希表
func TwoSum3(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		mp[v] = i
	}
	for i, v := range nums {
		k, ok := mp[target-v]
		if ok && i != k {
			return []int{i, k}
		}
	}
	return nil
}

// TwoSum 自己写的哈希表
func TwoSum(nums []int, target int) (res []int) {
	mp := make(map[int][]int)
	for i, v := range nums {
		if t, ok := mp[v]; ok {
			mp[v] = append(t, i)
		} else {
			mp[v] = []int{i}
		}
	}
	for k, v := range mp {
		tmp := target - k
		if t, ok := mp[tmp]; ok && k == tmp {
			if len(t) == 2 {
				res = t
				break
			}
			continue
		} else if ok {
			res = append(res, t[0])
			res = append(res, v[0])
			break
		}
	}
	return res
}

// TwoSum1 暴力枚举
func TwoSum1(nums []int, target int) (res []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
				break
			}
		}
	}
	return
}
