/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-02 16:40:34
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-03 17:39:01
 * @Description:
	下面的代码，将使用始于零的数组来表示字符串. 比如，若字符串S = "ABC"，则S[2]表示字符'C'

	KMP算法: 用于在一个文本串S内查找一个模式串P 的出现位置

*/
package kmp

import (
	"fmt"
	"testing"
)

/*
	以模式串P ="ABCDABD"为例
	- 模式串P 的 next 数组定义为：
		一般定义next[0] = -1
		next[k],next 数组各值的含义：代表当前位置的字符P[k]之前的子串P[0:k]的最大相同前后缀长度。
		next 数组相当于告诉我们： 如模式串P中在j 处的字符跟文本串S在i 处的字符匹配失配时，
			j需要移动到 P[最大相同前后缀的长度]的位置,也就是移动到P[next[j]]的位置.即next数组其实记录了此时j需要回退到P的哪一个下标.
			这样就可以用next [j] 处的字符继续跟文本串i 处的字符匹配.
			相当于模式串向右移动 j - next[j] 位。


	  - 例如如果next [i] = k，代表P[i] 之前的子串P[0:i]的最大相同前后缀长度为k 。
	  	如:"ABCDABD"的next数组中next[5] = 1,代表P[5]字符'B'之前的子串P[0:5]"ABCDA"的最大相同前后缀长度为1.


	①寻找最长公共前后缀长度,求得部分匹配表.
		它的各个子串的最长公共前后缀长度如下:(真子串:指不等于原串的子串)
			"A" 		0
			"AB" 		0
			"ABC"		0
			"ABCD" 		0
			"ABCDA" 	1
			"ABCDAB"	2
			"ABCDABD"	0

	②根据部分匹配表可以得到next数组
		next 数组考虑的是除当前字符外的最长相同前缀后缀，所以通过第①步骤求得各个前缀后缀的公共元素的最大长度后，只要稍作变形即可：
			将第①步骤中求得的值整体右移一位，然后初值赋为-1，
		  比如:对于aba来说，第3个字符a之前的字符串ab中有长度为0的相同前缀后缀，所以第3个字符a对应的next值为0；
		  	而对于abab来说，第4个字符b之前的字符串aba中有长度为1的相同前缀后缀a，所以第4个字符b对应的next值为1（相同前缀后缀的长度为k，k = 1）。

	③根据next数组进行匹配

*/

// "ABABAABAA"
// 部分匹配表,可以将部分匹配表直接作为next数组.
// 可以将部分匹配表整体右移一位.next[0] = -1.作为next数组.
func PrefixTable(p string) []int {
	prefix := make([]int, len(p))
	j := 0 //指向前缀末尾位置,也代表着p[0:i+1]的最大相等前后缀的长度.
	i := 1 //指向后缀末尾位置,初始化为1,这样才能开始与前缀进行比较.
	prefix[0] = 0
	for i < len(p) {
		// 前后缀不相同
		for j > 0 && p[j] != p[i] {
			j = prefix[j-1] //j回退,找到j前一位对应的prefix值,就是回退的下标.
		}
		if p[j] == p[i] {
			j++ //子串p[0:i+1]的最大相等前后缀的长度加+1
		}
		prefix[i] = j //p[0:i+1]的最大相等前后缀长度为j.
		i++
	}
	return prefix
}

// "DABCDABDE"
// 直接得到next数组.
func KmpNext(p string) []int {
	next := make([]int, len(p))
	next[0] = -1
	j := -1
	i := 0
	for i < len(p)-1 {
		//p[k]表示前缀，p[j]表示后缀
		if j == -1 || p[i] == p[j] {
			j++
			i++
			next[i] = j
		} else {
			j = next[j] //?
		}
	}
	return next
}

// "ABABAABAA"
func KmpNextval(p string) []int {
	next := make([]int, len(p))
	next[0] = -1
	j := -1
	i := 0
	for i < len(p)-1 {
		//p[k]表示前缀，p[j]表示后缀
		if j == -1 || p[i] == p[j] {
			j++
			i++
			if p[i] != p[j] {
				next[i] = j
			} else {
				next[i] = next[j]
			}
		} else {
			j = next[j] //?
		}
	}
	return next
}

// 使用部分匹配表作为next数组.
func KmpSearchByPrefix(str, substr string) (index int) {
	prefix := PrefixTable(substr)
	i, j := 0, 0
	for i < len(str) && j < len(substr) {
		if j != 0 && str[i] != substr[j] {
			j = prefix[j-1]
			continue
		}
		if str[i] == substr[j] {
			j++
		}
		i++
	}
	if j == len(substr) {
		return i - j
	}
	return -1
}

// 直接使用next数组.
func KmpSearchByNext(str, substr string) (index int) {
	next := KmpNext(substr)
	i, j := 0, 0
	for i < len(str) && j < len(substr) {
		if str[i] != substr[j] {
			j = next[j]
			if j == -1 {
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}
	if j == len(substr) {
		return i - j
	}
	return -1
}

// 直接使用next数组.
// 上面👆🏻代码的改进.
func KmpSearchByNextval(str, substr string) (index int) {
	// next := KmpNext(substr)
	next := KmpNextval(substr)
	i, j := 0, 0
	for i < len(str) && j < len(substr) {
		if j == -1 || str[i] == substr[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(substr) {
		return i - j
	}
	return -1
}

func Test(t *testing.T) {
	prefix := PrefixTable("ABCDABD")
	next := KmpNext("ABCDABD")
	nextval := KmpNextval("ABCDABD")
	fmt.Println(prefix)
	fmt.Println(next)
	fmt.Println(nextval)

	prefix1 := PrefixTable("ABABAAABABAA")
	next1 := KmpNext("ABABAAABABAA")
	nextval1 := KmpNextval("ABABAAABABAA")
	fmt.Println(prefix1)
	fmt.Println(next1)
	fmt.Println(nextval1)

	s := "ABDCAABABAAABABAAFDA"
	p := "ABABAAABABAA"
	fmt.Println(KmpSearchByNext(s, p))
	fmt.Println(KmpSearchByNextval(s, p))
	fmt.Println(KmpSearchByPrefix(s, p))

	s = "BBC ABCDAB ABCDABCDABDE"
	p = "ABCDABD"
	fmt.Println(KmpSearchByNext(s, p))
	fmt.Println(KmpSearchByNextval(s, p))
	fmt.Println(KmpSearchByPrefix(s, p))
}
