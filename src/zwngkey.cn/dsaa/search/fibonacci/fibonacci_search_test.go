/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-01 18:20:58
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-01 19:25:33
 */
package fibonacci

import "testing"

func Test(t *testing.T) {
	seq := []int{1, 8, 10, 89, 1000, 1234}
	target := 1235
	t.Log(FibonacciSearch(seq, target))
}
