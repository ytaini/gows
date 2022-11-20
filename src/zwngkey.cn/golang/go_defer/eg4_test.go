/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-20 14:54:36
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-20 15:08:14
 * @Description:
 */
package godefer

import (
	"fmt"
	"testing"
)

func ctest(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func Test3(t *testing.T) {
	a := 1
	b := 2
	//defer ctest("1",1,3). 此时会先执行ctest("10",1,2)函数.因为对于一个延迟函数调用，它的实参是在此调用被推入延迟调用堆栈的时候被估值的。
	defer ctest("1", a, ctest("10", a, b)) //👇🏻的也是如此
	a = 0
	defer ctest("2", a, ctest("20", a, b)) // defer ctest("2",0,2)
	b = 1
	defer ctest("3", a, ctest("30", a, b)) // defer ctest("3",0,1)
}
