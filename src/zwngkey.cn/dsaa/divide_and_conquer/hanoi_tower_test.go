/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 08:20:34
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-19 05:52:33
 * @Description:
	分治算法经典应用:汉诺塔问题.

	汉诺塔问题可以分为2步:
	 1. 如果只有一个盘,A(起点)->C(终点)
	 2. 如果有≥2盘,我们总是可以将其看做是两个盘
	 	1.最下边的盘
		2.上面的所有盘(看做一个盘)
	  1.先把最上面的盘 A(起点)->B(终点)
	  2.把最下边的盘A(起点)->C(终点)
	  3.最后B(起点)->C(终点)
*/
package divideandconquer

import (
	"fmt"
	"testing"
)

// n: 盘的总数.
// start(起点),guodu(过度),c(终点) 表示3根柱子
func HanoiTower(n int, start, guodu, end rune) {
	// 如果只有一个盘,A->C
	if n == 1 {
		fmt.Printf("从%c塔-->%c塔\n", start, end)
	} else {
		// 1.先把A柱子最上面的所有盘(n-1) 移动到 B柱子上.
		HanoiTower(n-1, start, end, guodu) //将B当做终点,C当做过度
		// 2.把A柱子最下边的盘 移动到 C柱子上
		fmt.Printf("从%c塔-->%c塔\n", start, end)
		// 3.再把B柱子上所有盘 移动到 C柱子上.
		HanoiTower(n-1, guodu, start, end) //将B当做起点,A当做过度
	}
}
func Test(t *testing.T) {
	HanoiTower(5, 'A', 'B', 'C') //目标: 将'A'柱子上所有的盘子移动到'C'柱子上.
}
