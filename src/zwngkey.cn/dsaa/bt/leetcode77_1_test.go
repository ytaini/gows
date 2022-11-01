/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-31 17:35:29
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-31 22:09:13
 * @Description:
	给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。你可以按 任何顺序 返回答案。
*/
package bt

import (
	"fmt"
	"testing"
)

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。你可以按 任何顺序 返回答案。
func combine(n, k int) (result [][]int) {
	if k <= 0 || k > n {
		return result
	}

	// 定义回溯函数bt.
	var bt func(int)

	// 记录当前路径
	var path []int

	bt = func(cur int) {
		// 递归终止条件是：path 的长度等于 k
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			//不能直接将path加入到result.因为result中存的是path的地址.
			result = append(result, temp) // 记录合法答案
			return
		}
		// 分析搜索起点的上界，其实是在深度优先遍历的过程中剪枝，剪枝可以避免不必要的遍历，剪枝剪得好，可以大幅度节约算法的执行时间。

		// 事实上，如果 n = 7, k = 4，从 5 开始搜索就已经没有意义了，因为后面的数只有 6 和 7，一共就 3 个候选数，凑不出 4 个数的组合了。
		// 这题中:搜索起点的上界 + 接下来要选择的元素个数(K- len(path)) - 1 = n
		for i := cur; i <= n-(k-len(path))+1; i++ { // 遍历可能的搜索起点
			//增加路径
			path = append(path, i)
			fmt.Println("bt", i, "递归之前 => ", path)
			bt(i + 1) // 下一轮搜索，设置的搜索起点要加 1，因为组合数理不允许出现重复的元素

			// 重点理解这里：深度优先遍历有回头的过程，因此递归之前做了什么，递归之后需要做相同操作的逆向操作
			path = path[:len(path)-1]
			fmt.Println("bt", i, "回溯 => ", path)
		}

		// 剪枝前.
		// for i := cur; i <= n; i++ {
		// 	path = append(path, i)
		// 	bt(i + 1)
		// 	path = path[:len(path)-1]
		// }
	}

	bt(1)

	return result
}

func Test(t *testing.T) {
	res := combine(5, 4)
	fmt.Println(res)
}
