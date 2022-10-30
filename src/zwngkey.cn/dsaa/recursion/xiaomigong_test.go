/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-31 00:20:13
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-31 03:08:03
 * @Description:
	递归解决迷宫问题.
	说明:
		小人得到的路径，和程序员设置的找路策略有关即：找路的上下左右的顺序相关
		再得到小人路径时，可以先使用(下右上左)，再改成(上右下左)，看看路径是不是有变化
		测试回溯现象

	策略决定了算法.
*/
package recursion

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	// 定义二维切片,模拟迷宫
	map1 := make([][]int, 11)
	for i := range map1 {
		map1[i] = make([]int, 9)
	}
	// 初始化迷宫
	Init1(map1)
	// Init2(map1)

	// 显示迷宫
	Print(map1)

	// 人初始位置(1,1)
	// 开始找路
	findWay(map1, 1, 1)
	fmt.Println("------------")

	Print(map1)
}

// 使用递归回溯来给小人找路
// * 设终点为右下角[9,7]位置
// * 当map[i][j]为0时，表示该点可以走
// * 当map[i][j]为1时，表示墙.
// * 当map[i][j]为2时，表示该点走过了.
// * 当map[i][j]为3时，表示该点走过,但走不通.
// 走迷宫时，需要确定一个策略（先走哪个方向,再走哪个反向） 下 -> 右 -> 上 -> 左 ， 如果该点走不通，再回溯
func findWay(map1 [][]int /* 当前的迷宫状态 */, i, j int /*  i,j 人当前位置 */) bool /* 通路是否形成 */ {

	if map1[9][7] == 2 {
		return true
	}

	// 每走一步判断(下右上左)4个方向能不能走,每走一步就将前一步置为2(表示走过了)
	if map1[i][j] == 0 {
		map1[i][j] = 2
		if findWay(map1, i+1, j) {
			return true
		} else if findWay(map1, i, j+1) {
			return true
		} else if findWay(map1, i-1, j) {
			return true
		} else if findWay(map1, i, j-1) {
			return true
		} else {
			map1[i][j] = 3
			return false
		}
	} else {
		return false
	}
}

// 显示迷宫
func Print(map1 [][]int) {
	for _, v := range map1 {
		fmt.Println(v)
	}
}

// 初始化地图
func Init1(map1 [][]int) {
	for i := 0; i < 9; i++ {
		map1[0][i] = 1
		map1[1][i] = 1
		map1[10][i] = 1
	}
	for i := 5; i < 8; i++ {
		map1[1][i] = 0
	}
	for i := 0; i < 11; i++ {
		map1[i][0] = 1
		map1[i][8] = 1
	}
	for i := 0; i < 8; i++ {
		map1[6][i] = 1
	}
	for i := 0; i < 5; i++ {
		map1[i][4] = 1
	}
	for i := 2; i < 6; i++ {
		map1[i][6] = 1
	}
	for i := 2; i < 9; i++ {
		map1[8][i] = 1
	}
	map1[3][1] = 1
	map1[3][2] = 1
	map1[6][7] = 0
	map1[1][1] = 0
}

// 初始化地图
func Init2(map1 [][]int) {
	for i := 0; i < 9; i++ {
		map1[0][i] = 1
		map1[10][i] = 1
	}
	for i := 0; i < 11; i++ {
		map1[i][0] = 1
		map1[i][8] = 1
	}
	map1[3][1] = 1
	map1[3][2] = 1
	map1[3][3] = 1
	map1[3][4] = 1
}
