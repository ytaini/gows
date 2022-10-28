/*
 * @Author: zwngkey
 * @Date: 2022-04-21 23:16:37
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 18:26:22
 * @Description:
 */
package dp

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	fmt.Printf("CoinChange([]int{2, 5, 7}, 27): %v\n", CoinChange([]int{2, 5, 7}, 27))
}

func TestUniquePaths(t *testing.T) {
	fmt.Printf("UniquePaths(3, 3): %v\n", UniquePaths(3, 3))
}

func TestClimbStairs(t *testing.T) {
	fmt.Printf("ClimbStairs(10): %v\n", ClimbStairs(10))
}

func TestDeleteAndEarn(t *testing.T) {
	DeleteAndEarn([]int{2, 2, 3, 3, 3, 4})
}

func TestDPEraseOverlapIntervals1(t *testing.T) {
	fmt.Printf("greedy.EraseOverlapIntervals(): %v\n", EraseOverlapIntervals([][]int{{-52, 31}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}))
}
