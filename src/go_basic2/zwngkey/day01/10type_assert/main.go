package main

import "fmt"

func main() {
	typeAssert(12)
	typeAssert("123")
	typeAssert([...]int{1, 2, 3})
	typeAssert([]int{1, 2, 3})
	typeAssert(false)
	typeAssert(map[string]int{})
	typeAssert(func() {})
	typeAssert('1')
	typeAssert(1.23)
	typeAssert(int8(8))
}

func typeAssert(a any) {
	fmt.Println()
	switch v := a.(type) {
	case int:
		fmt.Printf("a的类型:%T,a的值:%v", a, v)
	case string:
		fmt.Printf("a的类型:%T,a的值:%v", a, v)
	case []int:
		fmt.Printf("a的类型:%T,a的值:%v", a, v)
	case func():
		fmt.Printf("a的类型:%T", a)
	case int32:
		fmt.Printf("a的类型:%T,a的值:%v", a, v)
	default:
		fmt.Printf("a的类型:%T,a的值:%v", a, v)
	}
}
