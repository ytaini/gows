/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 18:58:47
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-06 19:01:28
 */
package usage

import (
	"fmt"
	"math/big"
	"testing"
)

func Test(t *testing.T) {
	ret := BigIntFactorial(big.NewInt(31))
	fmt.Println(ret)
}
