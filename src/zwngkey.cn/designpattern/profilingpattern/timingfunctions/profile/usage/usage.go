/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 18:45:22
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-06 18:57:04
 */
package usage

import (
	"math/big"
	"time"

	"zwngkey.cn/designpattern/profilingpattern/timingfunctions/profile"
)

// Usage : 求x!
func BigIntFactorial(x *big.Int) *big.Int {
	// Arguments to a defer statement is immediately evaluated and stored.
	// 延迟语句的参数会立即被求值并存储。
	// The deferred function receives the pre-evaluated values when its invoked.
	// 延迟函数在调用时会收到预计算的值。
	defer profile.Duration(time.Now(), "IntFactorial") //记录执行时间
	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
		y.Mul(y, x)
	}
	return x.Set(y)
}
