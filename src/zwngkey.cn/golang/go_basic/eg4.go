package gobasic

import "fmt"

func Testeg41() {

	//var a []int = make([]int,0)
	var a []int

	//上面两种声明方式,下面都error
	//a[0] = 1 //panic

	//上面两种声明方式,都可以调用append函数
	a = append(a, 1) //ac
	fmt.Println(a)

	//var m = map[int]int{}
	// var m map[int]int

	//不能给值为nil的map添加元素,必须初始化
	// m[1] = 1
}
