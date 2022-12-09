/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 04:49:41
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 05:35:59
 */
package strategy

type Cat struct {
	Height, Weight int
	Salary         float64
}

// 具体策略类
type CatWeightComparator struct{}

func (CatWeightComparator) Compare(t1, t2 Cat) int {
	return t1.Weight - t2.Height
}

// 具体策略类
type CatHeightComparator struct{}

func (CatHeightComparator) Compare(t1, t2 Cat) int {
	return t2.Height - t1.Height
}
