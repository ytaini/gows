/*
 * @Author: zwngkey
 * @Date: 2022-04-21 23:00:40
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:07:24
 * @Description:
 */
package dp

import (
	"math"
	"zwngkey.cn/dsaa/util"
)

/**
 * @param coins: a list of integer
 * @param amount: a total amount of money amount
 * @return: the fewest number of coins that you need to make up
 */
func CoinChange(coins []int, amount int) int {
	//0...n:[n+1]
	//f[i]表示 i金额可以换取的最少的硬币数量.
	f := make([]int, amount+1)
	//初始化
	//0金额可以换取0枚硬币
	f[0] = 0

	// f[1],f[2],f[3]...f[amount]
	// 依次计算 1,2,3.....n 的金额可以换取的最少硬币数量.
	for i := 1; i <= amount; i++ {
		//表示不能进行换取.
		f[i] = math.MaxInt
		//last coin coins[j]
		for j := range coins {
			//当 i 金额少于 所有硬币面值时,不能进行换取.
			//当f[i]=math.MaxInt时,
			if i >= coins[j] && f[i-coins[j]] != math.MaxInt {
				//假设有2,5,7面值的硬币.总金额为27.
				//则f[27] = min {f[27-2]+1,f[27-5]+1,f[27-7]+1}.
				//即总金额为27时,能换取的最少硬币数量
				//f[27-2]+1 :表示当最后一个硬币为2时,这时总金额为25,数量+1
				//此时求f[27]等于求f[25]+1
				//f[27-5]+1 :表示当最后一个硬币为5时,这时总金额为22,数量+1
				//此时求f[27]等于求f[22]+1
				//其他同理.
				//然后在这些情况中选择硬币数量最少的情况.
				f[i] = util.Min(f[i-coins[j]]+1, f[i])
			}
		}
	}
	if f[amount] == math.MaxInt {
		f[amount] = -1
	}
	return f[amount]
}
