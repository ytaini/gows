/*
 * @Author: zwngkey
 * @Date: 2022-04-24 04:43:28
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-28 10:08:46
 * @Description:
 */
package nt

import (
	"fmt"
	"testing"

	"zwngkey.cn/dsaa/util"
)

func TestGCD(t *testing.T) {
	fmt.Printf("util.GCD(): %v\n", util.GCD(25, 225))
}

func TestNtGcd(t *testing.T) {
	fmt.Printf("NTGCD([]int{3, 5, 10, 11, 15, 17, 19}): %v\n", NTGCD([]int{3, 5, 10, 11, 13, 17, 19}))
}

// [7,5,6,8,3]

func TestFindGCD(t *testing.T) {
	fmt.Printf("FindGCD([]int{7, 5, 6, 8, 3}): %v\n", FindGCD([]int{7, 5, 6, 8, 3}))
}

func Test1(t *testing.T) {
	fmt.Printf("HasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}): %v\n", HasGroupsSizeX([]int{2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 1, 1}))
}
