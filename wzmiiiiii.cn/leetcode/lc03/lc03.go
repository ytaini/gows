package lc03

// LengthOfLongestSubstring3 滑动窗口3
func LengthOfLongestSubstring3(s string) (max int) {
	mp := map[byte]int{}
	rk := -1 // 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	for l := 0; l < len(s); l++ {
		if l != 0 {
			delete(mp, s[l-1]) // 左指针向右移动一格，移除一个字符
		}
		for rk+1 < len(s) && mp[s[rk+1]] == 0 {
			// 不断地移动右指针
			mp[s[rk+1]]++
			rk++
		}
		// 第 l 到 rk 个字符是一个极长的无重复字符子串
		max = Max(max, rk-l+1)
	}
	return max
}

// LengthOfLongestSubstring2 滑动窗口2
func LengthOfLongestSubstring2(s string) (max int) {
	mp := map[byte]int{}
	l := 0                        // l:左指针
	for r := 0; r < len(s); r++ { // r:右指针
		if index, ok := mp[s[r]]; ok {
			l = Max(l, index+1)
		}
		mp[s[r]] = r
		max = Max(max, r-l+1)
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// LengthOfLongestSubstring1 滑动窗口1
func LengthOfLongestSubstring1(s string) (max int) {
	mp := map[byte]int{}
	for r := 0; r < len(s); r++ {
		if index, ok := mp[s[r]]; ok {
			for k, v := range mp {
				if v <= index {
					delete(mp, k)
				}
			}
		}
		mp[s[r]] = r
		max = Max(max, len(mp))
	}
	return max
}

// LengthOfLongestSubstring 暴力
func LengthOfLongestSubstring(s string) (max int) {
	for i := 0; i < len(s); i++ {
		if len(s)-i < max {
			break
		}
		mp := map[byte]struct{}{}
		mp[s[i]] = struct{}{}
		for j := i + 1; j < len(s); j++ {
			if _, ok := mp[s[j]]; !ok {
				mp[s[j]] = struct{}{}
			} else {
				break
			}
		}
		if len(mp) > max {
			max = len(mp)
		}
	}
	return max
}
