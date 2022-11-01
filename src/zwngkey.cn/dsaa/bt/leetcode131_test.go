/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-31 22:21:44
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-01 00:30:42
 * @Description:
	给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
*/
package bt

import (
	"fmt"
	"testing"
)

// 从所有的切割方案中,找出符合要求的方案.
func partition(s string) (result [][]string) {
	// 记录切割路径
	var path []string

	var bt func(curIndex int)

	// 用指针 curIndex 试着去切，切出一个回文串，基于新的 curIndex，继续往下切，直到 curIndex 越界
	// 每次基于当前的 curIndex，可以选择不同的 i，切出 curIndex 到 i 的子串，我们枚举出这些选项 i：
	// 	切出的子串满足回文，将它加入部分解 path 数组，并继续往下切（递归）
	// 	切出的子串不是回文，跳过该选择，不落入递归，继续下一轮迭代
	bt = func(curIndex int) {
		// 将符合要求的方案记录下来
		if curIndex == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp) // 将path的拷贝 加入解集result
			return                        // 结束掉当前递归（分支）
		}

		// 下面两种情况，是结束当前递归的两种情形：
		// 	1.指针越界了，没有可以切分的子串了，递归到这一步，说明一直在切出回文串，现在生成了一个合法的部分解，return
		// 	2.走完了当前递归的 for 循环，考察了基于当前 curIndex 的所有的切分可能，当前递归自然地结束

		//  枚举出当前的所有选项，从索引curIndex到末尾索引
		for i := curIndex; i < len(s); i++ {
			// 对每一次切割出来的子串进行回文判断.
			if !isPali(s[curIndex : i+1]) { //过滤不符合要求的方案
				continue
			}
			// 当前选择i，如果 curIndex 到 i 是回文串，就切它
			path = append(path, s[curIndex:i+1]) // 切出来，加入到部分解temp
			bt(i + 1)                            // 基于这个选择，继续往下递归，继续切
			path = path[:len(path)-1]            // 上面递归结束了，撤销当前选择i，去下一轮迭代
		}
	}

	bt(0) //从索引0开始，往后切回文串
	return
}

func isPali(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func Test2(t *testing.T) {
	// 0.284s
	fmt.Println(partition("aabbcddee"))
	// partition("aabbcddee")
}
