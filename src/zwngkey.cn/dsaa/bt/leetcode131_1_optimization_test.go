/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-01 00:05:20
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-01 00:31:30
 * @Description:
	优化leetcode131问题.
*/
package bt

import (
	"fmt"
	"testing"
)

/*
	还存在什么问题？
	每次递归都调用isPali判断，是否有必要？
		我们做了重复的计算，有的子串已经判断过是否回文了，就别再判断了。
		做法是用一个memo二维数组，将计算过的子问题结果存下来，下次再遇到就直接拿出来用。
*/

func partition1(s string) (result [][]string) {

	var path []string

	var bt func(curIndex int)

	// 记录判断过的子字符串 0:未计算,1:回文,2:不是回文
	var memo = make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
	}

	bt = func(curIndex int) {
		if curIndex == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := curIndex; i < len(s); i++ {

			if memo[curIndex][i] == 2 { //不回文，直接跳过
				continue
			}
			if memo[curIndex][i] == 1 || !isPali1(s, curIndex, i, memo) { //记录为回文 或没有记录但isPali调用为真
				path = append(path, s[curIndex:i+1])
				bt(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	bt(0)
	return
}

func isPali1(s string, l, r int, memo [][]int) bool {
	for l < r {
		if s[l] != s[r] {
			memo[l][r] = 2
			return false
		}
		l++
		r--
	}
	memo[l][r] = 1
	return true
}

func Test3(t *testing.T) {
	// 0.069s
	fmt.Println(partition1("aabbcddee"))
}
