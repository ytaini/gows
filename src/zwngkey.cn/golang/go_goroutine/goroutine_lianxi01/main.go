/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-21 19:42:12
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-22 18:58:47
 * @Description:
 */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// fmt.Println(runtime.GOMAXPROCS(0)) // 8
	// fmt.Println(runtime.NumCPU())      // 8
	// fmt.Println(runtime.Version())     // go1.19.3
	// fmt.Println(runtime.GOARCH)        // arm64
	// fmt.Println(runtime.GOOS)          // darwin
	// fmt.Println(runtime.Compiler)      // gc
	// fmt.Println(runtime.GOROOT())      // /usr/local/go
	// runtime.GOMAXPROCS(1)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

func a() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Println("B:", i)
	}
	// fmt.Println(runtime.NumGoroutine()) //3
}

/*
	B: 0
	B: 1
	B: 2
	B: 3
	B: 4
	B: 5
	B: 6
	A: 0
	A: 1
	A: 2
	A: 3
	A: 4
	A: 5
	B: 7
	B: 8
	B: 9
	A: 6
	A: 7
	A: 8
	A: 9
*/
