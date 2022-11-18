/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 04:44:03
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-18 04:46:32
 * @Description:
	BF 是 Brute Force 的缩写，中文译作暴力匹配算法，也叫朴素匹配算法。

	术语: 主串和模式串。
		简单来说，我们要在字符串 A 中查找子串 B，那么 A 就是主串，B 就是模式串。
*/
package bf

// 使用三指针暴力搜索
func BFSearch(str, substr string) (index int) {
	i, j := 0, 0
	//优化: 当i>len(str)-len(substr)时,子串str[i,len(str)]长度小于len(substr),不可能存在len(substr)长的子串了.
	v := 0 //记录当前i每次的开始位置.
	for i < len(str) && j < len(substr) {
		if v == len(str)-len(substr)+1 {
			return -1
		}
		if str[i] == substr[j] {
			j++
			i++
			continue
		}
		// 比较指针回溯
		i = i - j + 1 //i回到i+1的位置.
		j = 0         //j置0
		v = i         //记录i
	}
	// 只有当j等于子串长度时,才是找到了
	if j == len(substr) {
		return i - j
	}
	return -1
}
