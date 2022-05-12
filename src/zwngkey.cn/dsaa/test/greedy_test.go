/*
 * @Author: zwngkey
 * @Date: 2022-04-24 04:42:48
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 05:19:40
 * @Description:
 */
package test

import (
	"fmt"
	"testing"

	"zwngkey.cn/dsaa/greedy"
)

func TestFindContentChildren(t *testing.T) {
	greedy.FindContentChildren([]int{3, 15, 62, 8, 27, 94}, []int{1, 4, 5, 6, 161})
}
func TestEraseOverlapIntervals(t *testing.T) {
	fmt.Printf("greedy.EraseOverlapIntervals(): %v\n", greedy.EraseOverlapIntervals2([][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}))
}
