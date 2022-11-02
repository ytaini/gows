/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-01 16:41:14
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-02 12:59:24
 * @Description:
 	方法一: 基于数组的回溯.

	在n × n格的国际象棋上摆放n个皇后，使其不能互相攻击，即：任意两个皇后都不能处于同一行、同一列或同一斜线上，问有多少种摆法。

	如何判断在同一斜线上? 两点连成的线斜率为1.
		即如果点A(x1,y1)与点B(x2,y2)满足 |x1-x2|÷|y1-y2| = 1.那么我们就说A,B两点在同一斜线上.

*/
package main

import (
	"fmt"
	"math"
)

func main() {
	// 记录棋盘行列
	var size int

	fmt.Scanf("%d", &size)

	// **这种解法的核心**
	// 通过数组下标(代表所在行),和这个下标存储的值(代表列) 来代表 一个皇后正确摆放的位置
	// 通过整个数组来记录一种解法.
	var arr = make([]int, size)

	// 记录所有解法
	var result [][]int

	var bt func(int)

	bt = func(row int) {
		if row == size { //探测完了最后一行,确定了一种解法,记录这种解法.
			temp := make([]int, size)
			copy(temp, arr)
			result = append(result, temp)
			return //开始回溯到前面的行.在前面的行中寻找另一种解法.
		}
		// 递归深度就是row控制棋盘的行，每一层里for循环的col控制棋盘的列

		// 函数一开始进入for循环,row不变,即行不变,通过col变化即列不断变化.来回溯试探皇后可以放置的正确位置.
		for col := 0; col < size; col++ {
			arr[row] = col       // 探测第row行的不同列是否存在正确位置.
			if judge(row, arr) { // 第row行第col列可以放置皇后
				// n不断变化,行不断变化.
				bt(row + 1) //递归试探row+1行的皇后放置位置
			}
			// 若当前for自然结束,说明第k行不同列没有正确的摆放位置,就开始回溯.
			// 这里回溯的结果就是回退到第k-1行,第k-1行就会i++,即探测第k-1行的下一列.
		}

	}
	// 第一个皇后放在第一行第一列的位置.
	bt(0)

	// 输出每一种解法
	for _, res := range result {
		for i, v := range res {
			fmt.Printf("{%d,%d} ", i+1, v+1)
		}
		fmt.Println()
	}

}

// 判断第n个皇后,是否与前面n-1个皇后冲突.
func judge(n int, arr []int) bool {
	for i := 0; i < n; i++ {
		if arr[i] == arr[n] || math.Abs(float64(i-n)) == math.Abs(float64(arr[i]-arr[n])) {
			return false
		}
	}
	return true
}
