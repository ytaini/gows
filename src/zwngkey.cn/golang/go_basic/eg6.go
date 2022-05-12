package gobasic

import "fmt"

func Eg61() {
	//go中数组是值类型
	x := [3]int{
		1, 2, 3,
	}
	test(x)        //4,2,3
	fmt.Println(x) //1,2,3

}
func test(arr [3]int) {
	arr[0] = 4
	fmt.Println(arr)
}
