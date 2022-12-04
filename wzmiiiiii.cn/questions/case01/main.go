/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-03 18:05:33
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-03 18:19:13
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	// a, b := 1, 0
	// x := a / b // panic: runtime error: integer divide by zero
	// fmt.Println(x)
	a, b, c := 1.0, 1.0, 0.0
	x, y := a/c, b/c
	fmt.Println(x, y)   // +Inf
	fmt.Println(x == y) // true

	n := math.NaN()
	// Special cases are:
	// 	Sqrt(+Inf) = +Inf
	// 	Sqrt(±0) = ±0
	// 	Sqrt(x < 0) = NaN
	// 	Sqrt(NaN) = NaN
	m := math.Sqrt(-1)
	fmt.Println(n == m) //false
}
