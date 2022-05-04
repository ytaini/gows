package main

import (
	"fmt"
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

func main() {
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
