/*
 * @Author: zwngkey
 * @Date: 2021-11-21 20:47:24
 * @LastEditTime: 2022-05-02 16:36:32
 * @Description: 函数类型的值 与 string.Map函数.
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	f := s
	fmt.Printf("f: %T\n", f) //f: func(int) int
	fmt.Printf("f(2): %v\n", f(2))

	f = n
	fmt.Printf("f: %T\n", f) //f: func(int) int 函数s与函数n类型相同.
	fmt.Printf("f(2): %v\n", f(2))

	// f = p  编译错误.	can't assign func(int, int) int to func(int) int

	// var a func(int) int
	// a(3)  此处a的值为nil, 会引起panic错误

	fmt.Println("----------------------")

	fmt.Println(strings.Map(
		func(r rune) rune {
			fmt.Println(r)
			return r + 1
		}, "1234")) //2345
	fmt.Println(strings.Map(
		func(r rune) rune {
			fmt.Println(r)
			return r + 1
		}, "abcd")) //bcde

}

func s(n int) int { return n * n }

func n(m int) int { return -m }

func p(n, m int) int { return m * n }
