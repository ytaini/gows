/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 22:43:25
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-06 23:43:46
 */
package test

import (
	"fmt"
	"sync"
	"testing"

	"zwngkey.cn/designpattern/singleton/type2"
)

func Test2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			instance := type2.GetInstance1()
			fmt.Printf("%p\n", instance)
		}()
	}
	wg.Wait()
}
