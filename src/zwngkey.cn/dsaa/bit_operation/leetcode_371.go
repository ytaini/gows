/*
 * @Author: zwngkey
 * @Date: 2022-04-22 07:39:13
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 05:50:28
 * @Description:
 */
package bitoper

func GetSum(a, b int) (c int) {
	if b == 0 {
		return a
	} else {
		c = GetSum(a^b, (a&b)<<1)
		return c
	}
}
