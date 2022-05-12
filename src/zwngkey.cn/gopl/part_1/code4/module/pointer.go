/*
 * @Author: zwngkey
 * @Date: 2021-11-14 01:01:54
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:12:11
 * @Description:
 */
package module

import "fmt"

func TestPointer() {
	x := 1
	p := &x
	fmt.Println(*p)

	*p = 2

	fmt.Println(x)

}
