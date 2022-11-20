package main

import (
	"fmt"
	"strings"
	"testing"
	// "time"
)

/**
实现一个 fibonacci 函数，它返回一个函数（闭包）
*/
// 返回一个“返回int的函数”
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b - a
	}

}

func Test1(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// func main() {
// 	fmt.Println(fn1())
// }

// func fn1() (i int) {
// 	defer func() {
// 		i++
// 	}()
// 	return 1
// }

// func main() {
// 	startAt := time.Now()
// 	defer func() {
// 		fmt.Println(time.Since(startAt))
// 	}()
// 	time.Sleep(time.Second)
// }

func add(x int) func(int) int {
	return func(y int) int {
		y += x
		return y
	}
}
func Test(t *testing.T) {
	res := add(100)(200) // 100 + 200
	fmt.Printf("res: %v\n", res)
}

func makeSuffixFunc(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func Test2(t *testing.T) {
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
}

func calc(base int) (func(x int) int, func(y int) int) {
	// 对于add,del函数来说,base 相当于一个全局变量
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

func Test3(t *testing.T) {
	f1, f2 := calc(10)
	fmt.Printf("f1(1): %v\n", f1(1)) // 11
	fmt.Printf("f2(2): %v\n", f2(2)) // 9
	fmt.Printf("f1(3): %v\n", f1(3)) // 12
	fmt.Printf("f2(4): %v\n", f2(4)) // 8
	fmt.Printf("f1(5): %v\n", f1(5)) // 13
	fmt.Printf("f2(6): %v\n", f2(6)) // 7
}
