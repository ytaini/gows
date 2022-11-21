package gobasic

import "fmt"

/*
  使用简式声明重复声明变量
		你不能在一个单独的声明中重复声明一个变量，但在多变量声明中这是允许的，其中至少要有一个新的声明变量。
		重复变量需要在相同的代码块内，否则你将得到一个隐藏变量。
*/
/*
func Test1() {
	a := 1
	// a := 1 不能重复声明 error
	a, b := 0, 1 //可以,但其中至少要有一个新的声明变量。

}
*/

/*
偶然的变量隐藏

	如果你在一个新的代码块中犯了这个错误，将不会出现编译错误，但你的应用将不会做你所期望的事情。
*/
func Test2() {
	x := 1
	fmt.Println(x) //prints 1
	{
		fmt.Println(x) //prints 1
		x := 2         //在当前作用域下,重新声明了一个与上一级作用域同名的变量.
		fmt.Println(x) //prints 2
	}
	fmt.Println(x) //prints 1 (bad if you need 2)
}
