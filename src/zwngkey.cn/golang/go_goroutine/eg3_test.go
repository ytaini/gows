/*
 * @Author: zwngkey
 * @Date: 2022-05-12 17:23:24
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-12 17:27:55
 * @Description:
 */
package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func TestEg31(t *testing.T) {

	var wg sync.WaitGroup
	/* wg跟踪10个goroutine */
	size := 10
	wg.Add(size)
	/* 开启10个goroutine并发执行 */
	for i := 0; i < size; i++ {
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	/* Wait会一直阻塞，直到wg的计数器为0*/
	wg.Wait()
	fmt.Println("end")
}
