/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-01 23:26:26
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-02 14:07:38
 * @Description:
	n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。即：任意两个皇后都不能处于同一行、同一列或同一斜线上，

	给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

*/
package main

import "fmt"

/*
	根据题目描述:
		将 N 个皇后放置在N×N的棋盘上，一定是每一行有且仅有一个皇后，每一列有且仅有一个皇后，且任何两个皇后都不能在同一条斜线上。

	方法一: 基于集合的

	回溯的具体做法是：
		依次在每一行放置一个皇后，每次新放置的皇后都不能和已经放置的皇后之间有攻击，
			即新放置的皇后不能和任何一个已经放置的皇后在同一列以及同一条斜线上。
			当 N 个皇后都放置完毕，则找到一个可能的解，将可能的解的数量加 1。

	快速判断每个位置是否可以放置皇后?
		为了判断一个位置所在的列和两条斜线上是否已经有皇后，使用三个数组分别记录每一列以及两个方向的每条斜线上是否有皇后。

		列的表示法很直观，一共有 N 列，每一列的下标范围从 0 到 N−1，使用列的下标即可明确表示每一列。

		如何表示两个方向的斜线呢？
			对于每个方向的斜线，需要找到斜线上的每个位置的行下标与列下标之间的关系。
				*从左上到右下方向，同一条斜线上的每个位置满足行下标与列下标之差相等
				*从右上到左下方向，同一条斜线上的每个位置满足行下标与列下标之和相等
*/

func totalNQueens(n int) (ans int) {
	columns := make([]bool, n)        // 通过下标记录某个列上是否有皇后
	diagonals1 := make([]bool, 2*n-1) // 通过下标记录某条左上到右下的斜线上是否有皇后
	diagonals2 := make([]bool, 2*n-1) // 通过下标记录某条右下到左上的斜线上是否有皇后

	var backtrack func(int)

	backtrack = func(row int) {
		if row == n {
			ans++
			return
		}
		// 使用回溯模拟依次在棋盘的每一行放置一个皇后
		for col := 0; col < n; col++ {
			d1 := row - col + n - 1 //row-col范围为[-(n-1),n+1],row-col+n-1范围为[0,2n-2]->2n-1个数.
			d2 := row + col         //row-col范围[0,2n-2]->2n-1个数.
			// 判断当前位置所在的列和两条斜线上是否已经有皇后
			if columns[col] || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			//如果没有,将当前位置所在列以及所在斜线标记为已有皇后
			columns[col] = true
			diagonals1[d1] = true
			diagonals2[d2] = true
			backtrack(row + 1) //递归,对下一行进行放置
			columns[col] = false
			diagonals1[d1] = false
			diagonals2[d2] = false
		}
	}
	backtrack(0)
	return
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Printf("totalNQueens(n): %v\n", totalNQueens(n))
}
