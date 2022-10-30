/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-30 17:55:01
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 18:27:37
 * @Description:

 */
package problem

import (
	"fmt"
	"testing"
)

func Test4(t *testing.T) {
	// exp := "100+((22+31)*44)-53"
	exp := "(3+4)*5+6*1"
	// exp := "(12+5)*(8-1)-6*6"
	// exp := "2+3*5"
	// exp := "2+3"
	// exp := "10-10"
	fmt.Println("中缀表达式字符串为: ", exp)
	// [100,+,(,(,22,+,31,),*,44,),-,53]

	res := Split(exp)
	fmt.Println("读取得到的中缀表达式为: ", res)

	// [100,22,31,+,44,*,+,53,-]
	res = InfixToSuffix(res)
	fmt.Println("转换得到的后缀表达式为: ", res)
}
