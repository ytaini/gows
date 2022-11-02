/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-02 16:40:34
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-03 00:36:37
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
	- 模式串P 的 next 数组定义为：next[i] 表示 P[0:i+1] 这一个子串，next 数组各值的含义：代表当前位置的字符P[i]之前的字符串P[0:i]中，有多大长度的相同前缀后缀。
	  - 例如如果next [i] = k，代表P[i] 之前的字符串P[0:i]中有最大长度为k 的相同前缀后缀。
	  	如:"ABCAB"的next数组中next[4] = 1,代表P[4]字符'B'之前的字符串P[0:4]"ABCD"中有最大长度为1的相同前后缀.

	  - 此也意味着在某个字符失配时，该字符对应的next 值会告诉你下一步匹配中，模式串应该跳到哪个位置（跳到next [j] 的位置）。
	  	如果next [j] 等于0或-1，则跳到模式串的开头字符，若next [j] = k 且 k > 0，代表下次匹配跳到j 之前的某个字符，而不是跳到开头，且具体跳过了k 个字符。


	①寻找最长公共前后缀长度
		以P ="ABCDABD"为例,它的各个子串的最长公共前后缀长度如下:(真子串:指不等于原串的子串)
			"A" 		0
			"AB" 		0
			"ABC"		0
			"ABCD" 		0
			"ABCDA" 	1
			"ABCDAB"	2
			"ABCDABD"	0

	②求next数组
		next 数组考虑的是除当前字符外的最长相同前缀后缀，所以通过第①步骤求得各个前缀后缀的公共元素的最大长度后，只要稍作变形即可：
			将第①步骤中求得的值整体右移一位，然后初值赋为-1，
		  比如:对于aba来说，第3个字符a之前的字符串ab中有长度为0的相同前缀后缀，所以第3个字符a对应的next值为0；
		  	而对于abab来说，第4个字符b之前的字符串aba中有长度为1的相同前缀后缀a，所以第4个字符b对应的next值为1（相同前缀后缀的长度为k，k = 1）。

	③根据next数组进行匹配

	KMP的next 数组相当于告诉我们：当模式串中的某个字符跟文本串中的某个字符匹配失配时，模式串下一步应该跳到哪个位置。
		如模式串中在j 处的字符跟文本串在i 处的字符匹配失配时，下一步用next [j] 处的字符继续跟文本串i 处的字符匹配，相当于模式串向右移动 j - next[j] 位。

*/
func KmpNext(p string) []int {
	next := make([]int, len(p))
	next[0] = -1
	k := -1
	j := 0
	for j < len(p)-1 {
		//p[k]表示前缀，p[j]表示后缀
		if k == -1 || p[j] == p[k] {
			k++
			j++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}

func KmpSearch(str, substr string) (index int) {
	next := KmpNext(substr)
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
	next := KmpNext("ABCDABD")
	for _, v := range next {
		fmt.Println(v)
	}
}
