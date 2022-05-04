package main

import "fmt"

var name string = "zhangsan"

// 简短变量声明语句只有对已经在同级词法域声明过的变量才和赋值操作语句等价，
// 如果变量是在外部词法域声明的，那么简短变量声明语句将会在当前词法域重新声明一个新的变量。
func main() {
	name := "lisi"
	fmt.Println(name)
}
