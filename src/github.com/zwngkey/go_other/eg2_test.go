/*
 * @Author: zwngkey
 * @Date: 2022-05-04 19:33:22
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-04 19:41:29
 * @Description:
 	go哪些函数调用将在编译时刻被估值？
		如果一个函数调用在编译时刻被估值，则估值结果为一个常量。

		函数			返回类型				其调用是否总是在编译时刻估值？
	unsafe.Sizeof	uintptr						是
	unsafe.Alignof
	unsafe.Offsetof

	complex			默认类型为complex128	 表达式complex(sr, si)只有在sr和si都为常量表达式的时候才在编译时刻估值。
					（结果为类型不确定值）

	real			默认类型为float64		 表达式real(s)和imag(s)在s为一个复数常量表达式时才在编译时刻估值。
	imag			（结果为类型不确定值）

	len				int						如果表达式s表示一个字符串常量，则表达式len(s)将在编译时刻估值；
	cap										如果表达式s表示一个数组或者数组的指针，并且s中不含有数据接收操作
											和估值结果为非常量的函数调用，则表达式len(s)和cap(s)将在编译时刻估值。


*/
package goother
