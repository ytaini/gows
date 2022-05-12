package main

import (
	"fmt"
	"strings"
)

func add(x int) func(int) int {
	return func(y int) int {
		y += x
		return y
	}
}

func makeSuffixFunc(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(x int) int, func(y int) int) {
	add := func(x int) int {
		base += x
		return base
	}

	del := func(y int) int {
		base -= y
		return base
	}

	return add, del

}

func ctest(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	res := add(100)(200) // 100 + 200
	fmt.Printf("res: %v\n", res)

	fmt.Println("----------------------------------------------------")

	jpgFunc := makeSuffixFunc(".jpg") //得到判断给定的name的后缀是否为.jpg的函数
	pngFunc := makeSuffixFunc(".png") //得到判断给定的name的后缀是否为.png的函数

	res1 := jpgFunc("test")
	res3 := jpgFunc("test.jpg")
	res2 := pngFunc("test")
	res4 := pngFunc("test.png")

	fmt.Printf("res1: %v\n", res1)
	fmt.Printf("res2: %v\n", res2)
	fmt.Printf("res3: %v\n", res3)
	fmt.Printf("res4: %v\n", res4)

	fmt.Println("----------------------------------------------------")

	f1, f2 := calc(10)

	fmt.Printf("f1(1): %v\n", f1(1)) // 11
	fmt.Printf("f2(2): %v\n", f2(2)) // 9
	fmt.Printf("f1(3): %v\n", f1(3)) // 12
	fmt.Printf("f2(4): %v\n", f2(4)) // 8
	fmt.Printf("f1(5): %v\n", f1(5)) // 13
	fmt.Printf("f2(6): %v\n", f2(6)) // 7

	fmt.Println("----------------------------------------------------")

	a := 1
	b := 2
	defer ctest("1", a, ctest("10", a, b))
	a = 0
	defer ctest("2", a, ctest("20", a, b))
	b = 1
	defer ctest("3", a, ctest("30", a, b))
}
