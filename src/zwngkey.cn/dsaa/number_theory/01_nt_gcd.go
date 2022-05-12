/*
 * @Author: zwngkey
 * @Date: 2022-04-24 04:01:34
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:08:02
 * @Description:
 */
package nt

import "zwngkey.cn/dsaa/util"

/*
问题描述:
给定一个长度为n的整型数组A,你可以对数组进行操作,每次操作选定一个下标i(1≤i≤n),再确定一个整数k(0≤k≤10^9),令A[i] = A[i] + k,
最少需要多少次操作,才能使数组中任意相邻两数的最大公约数为1. 即对于任意i(2≤i≤n),gcd(A[i-1],A[i]) = 1
如果无论怎么操作都无法让数组满足条件,返回-1

其中1≤A[i]≤10^9,2≤n≤10^5
*/

/*
问题分析:
	对于任意两数i,j 如果gcd(i,j) = m , 无论使i或j 乘以 一个多大的整数k,gcd(i,j) 的值并不会变小,反而更大.
	而对于任意两数,其最小公约数也为1,因此只有当原数组本身就满足条件(也就是不需要操作),才能满足题意.

*/
func NTGCD(A []int) int {
	for i := 0; i < len(A)-1; i++ {
		if util.GCD(A[i], A[i+1]) != 1 {
			return -1
		}
	}
	return 0
}
