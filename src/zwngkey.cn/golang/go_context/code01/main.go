/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-27 11:39:11
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-27 12:42:51
 * @Description:
 */
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		Monitor(ctx)
	}()
	wg.Wait()
}

func Monitor(ctx context.Context) {
	// 单纯的透传并不会起作用
	// 即使context透传下去了，不使用也是不起任何作用的
	// for {
	// 	fmt.Println("Monitor")
	// 	time.Sleep(time.Second)
	// }

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("Monitor")
		}
		time.Sleep(time.Second)
	}
}
