/*
 * @Author: zwngkey
 * @Date: 2022-04-22 07:40:37
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 05:50:39
 * @Description:
 */
package bitoper

func InsertBits(N int, M int, i int, j int) int {
	for k := i; k <= j; k++ {
		N = N &^ (1 << k)
	}
	return N | (M << i)
}
func InsertBits2(N int, M int, i int, j int) int {
	left := N >> (j + 1)
	left = left << (j + 1)
	mid := M << i
	right := N & (1<<i - 1)
	return left | mid | right
}
