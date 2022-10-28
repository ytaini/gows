/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-28 10:41:00
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 12:31:51
 * @Description:
	稀疏数组的应用
		稀疏数组可以看做是普通数组的压缩，但是这里说的普通数组是值无效数据量远大于有效数据量的数组

     	稀疏数组可以简单的看作为是压缩，在开发中也会使用到。比如将数据序列化到磁盘上，减少数据量，在IO过程中提高效率等等。

		 为什么要进行压缩？
         - 由于稀疏矩阵中存在大量的“空”值，占据了大量的存储空间，而真正有用的数据却少之又少，
         - 且在计算时浪费资源，所以要进行压缩存储以节省存储空间和计算方便。

	概念:
		当一个数组中大部分元素为同一值时，可以使用稀疏数组来保存该数组。

		稀疏数组的处理方式是:
			记录数组一共有几行几列，有多少个不同值；
			把具有不同值的元素和行列及值记录在一个小规模的数组中，从而缩小程序的规模
*/
package sparsearray

import (
	"fmt"
	"testing"
)

// 稀疏数组元素定义
type dataNode struct {
	row  int //在稀疏矩阵中行
	col  int //在稀疏矩阵中列
	data any //在稀疏矩阵中数据
}

func newNode(row, col int, data any) *dataNode {
	return &dataNode{row, col, data}
}

// 定义稀疏数组
// 数组的第一个元素存储矩阵的行,列以及不同的值的个数
type SArray struct {
	Arr []*dataNode
}

// 创建一个稀疏数组
func NewSArray() *SArray {
	return new(SArray)
}

// 向稀疏数组添加元素
func (s *SArray) Append(data *dataNode) {
	s.Arr = append(s.Arr, data)
}

// 将稀疏矩阵转换为稀疏数组
func (s *SArray) Convert(arr [][]any /*稀疏矩阵*/, zeorValue any /*稀疏矩阵中的“空”值*/) {
	totalRow := len(arr)
	totalCol := len(arr[0])
	var diffTotal int

	temp := newNode(totalRow, totalCol, diffTotal)
	s.Append(temp)

	for i, v := range arr {
		for j, val := range v {
			if val != zeorValue {
				diffTotal++
				tempNode := newNode(i, j, val)
				s.Append(tempNode)
			}
		}
	}
	s.Arr[0].data = diffTotal

}

// 将稀疏数组转换为稀疏矩阵
func (s *SArray) ConvertTo() [][]any {
	arr := s.Arr
	data := arr[0]
	matrix := NewMatrix(data.row, data.col)
	for i := 1; i < len(arr); i++ {
		matrix[arr[i].row][arr[i].col] = arr[i].data
	}
	return matrix
}

func (s *SArray) Show() {
	for _, v := range s.Arr {
		fmt.Printf("%d\t%d\t%v\n", v.row, v.col, v.data)
	}
}

// 创建一个稀疏矩阵
func NewMatrix(row, col int) [][]any {
	matrix := make([][]any, row)
	for i := range matrix {
		matrix[i] = make([]any, col)
	}
	return matrix
}

// 获取稀疏矩阵数据
func GetData() [][]any {
	matrix := NewMatrix(11, 11)
	matrix[1][2] = 1
	matrix[2][3] = 2
	matrix[8][9] = 3
	matrix[3][2] = 4
	return matrix
}

// 遍历一个稀疏矩阵
func Print(matrix [][]any) {
	for _, v := range matrix {
		for _, v1 := range v {
			fmt.Printf("%v\t", v1)
		}
		fmt.Println()
	}
}

func Test1(t *testing.T) {
	matrix := GetData()
	fmt.Println("原数组:")
	Print(matrix)

	sa := NewSArray()
	sa.Convert(matrix, nil)
	fmt.Println("转换后的稀疏数组:")
	sa.Show()

	fmt.Println("将稀疏数组转换为矩阵:")
	v := sa.ConvertTo()
	Print(v)
}
