/*
 * @Author: zwngkey
 * @Date: 2021-11-11 01:30:36
 * @LastEditTime: 2022-10-28 12:31:26
 * @Description:
	稀疏数组1
*/
package sparsearray

import (
	"fmt"
	"testing"
)

type ValNode struct {
	row int
	col int
	val int
}

const baizi = 1
const heizi = 2
const defDate = 0

func Test(t *testing.T) {

	// 1.创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = baizi //白子
	chessMap[2][3] = heizi //黑子

	// 2.查看原始数组
	printSparseArray(chessMap)
	// 3.转成稀疏数组。
	/*
		思路：
		1.遍历chessMap,如果我们发现有一个元素的值不为0，则创建一个node结构体
		2.将其放入到对应的切片即可。
	*/

	//定义结构体切片
	var sparseArr []ValNode

	//记录元数组中的总行总列与默认值
	sparseArr = append(sparseArr, ValNode{
		row: 11,
		col: 11,
		val: 0,
	})

	//记录原数组关键信息
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != defDate {
				//创建值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	for i, vn := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, vn.col, vn.row, vn.val)
	}

	//将这个稀疏数组存盘

	//恢复原始数组
	var chessMap2 [11][11]int

	for _, v := range sparseArr {
		chessMap2[v.row][v.col] = v.val
	}

	printSparseArray(chessMap2)

}

func printSparseArray(chessMap [11][11]int) {
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
