/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-02 15:07:38
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-02 18:02:44
 * @Description:
	在目标(字符串)中寻找一个给定的模式(也是字符串)，返回目标和模式匹配的第一个子串的首字符位置.没有返回-1.
*/
package gostring

import (
	"fmt"
	"testing"
)

/*
	下面的算法在go中匹配中文不会返回正确的index,
	go中string的底层数组为byte数组.
*/

func findSubStringIndex(str, substr string) (index int) {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:len(substr)+i] == substr {
			return i
		}
	}
	return -1
}

// 使用三指针暴力搜索
func findSubStringIndex2(str, substr string) (index int) {
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

func Test41(t *testing.T) {
	str := "aaaaaaaabab"
	substr := "aaaba"
	// substr := "aaaaaaaabab"
	// substr := ""
	// substr := "aaaaaaaaaaaba"
	fmt.Println(findSubStringIndex(str, substr))
	fmt.Println(findSubStringIndex2(str, substr))
}

func Test42(t *testing.T) {
	str := "你好啊"
	for i, v := range str {
		fmt.Println(string(v), string(str[i]), i)
		fmt.Printf("%c,%c,%d\n", v, str[i], i)
		fmt.Printf("%v,%v,%d\n", v, str[i], i)
		// 输出:
		// 你 ä 0
		// 你,ä,0
		// 20320,228,0
		// 好 å 3
		// 好,å,3
		// 22909,229,3
		// 啊 å 6
		// 啊,å,6
		// 21834,229,6
	}
	fmt.Println([]byte(str))
	// [228 189 160 229 165 189 229 149 138]
	fmt.Println([]rune(str))
	// [20320 22909 21834]
}
