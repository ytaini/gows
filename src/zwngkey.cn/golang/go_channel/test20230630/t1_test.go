/*
 * @Author: wzmiiiiii
 * @Date: 2023-06-30 00:31:37
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2023-06-30 00:40:32
 */
package test20230630

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

var (
	concurrent       int32
	concurrent_limit = make(chan struct{}, 10)
)

// 接口限流:
//
//	通过带容量的channel.
func readDB() string {
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("readDB()调用并发度 %d\n", atomic.LoadInt32(&concurrent))
	time.Sleep(200 * time.Millisecond)
	atomic.AddInt32(&concurrent, -1)
	return "OK"
}

func handler3() {
	concurrent_limit <- struct{}{}
	readDB()
	<-concurrent_limit
}

func Test01(t *testing.T) {
	for i := 0; i < 100; i++ {
		go handler3()
	}
	time.Sleep(3 * time.Second)
}
