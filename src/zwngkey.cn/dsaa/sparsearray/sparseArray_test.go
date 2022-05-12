/*
 * @Author: zwngkey
 * @Date: 2021-11-11 01:30:36
 * @LastEditTime: 2022-04-30 23:36:22
 * @Description:
	稀疏数组的应用
		稀疏数组可以看做是普通数组的压缩，但是这里说的普通数组是值无效数据量远大于有效数据量的数组

     	稀疏数组可以简单的看作为是压缩，在开发中也会使用到。比如将数据序列化到磁盘上，减少数据量，在IO过程中提高效率等等。

		 为什么要进行压缩？
         - 由于稀疏矩阵中存在大量的“空”值，占据了大量的存储空间，而真正有用的数据却少之又少，
         - 且在计算时浪费资源，所以要进行压缩存储以节省存储空间和计算方便。

	概念:
		当一个数组中大部分元素为0，或者为同一值的数组时，可以使用稀疏数组来保存该数组。

		稀疏数组的处理方式是:
			记录数组一共有几行几列，有多少个不同值；
			把具有不同值的元素和行列及值记录在一个小规模的数组中，从而缩小程序的规模



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
