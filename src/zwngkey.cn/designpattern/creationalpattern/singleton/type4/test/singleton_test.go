/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 22:43:25
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 17:11:16
 */
package test

import (
	"fmt"
	"sync"
	"testing"

	"zwngkey.cn/designpattern/creationalpattern/singleton/type4"
)

func Test2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			instance := type4.GetInstance()
			fmt.Printf("%p\n", instance)
		}()
	}
	wg.Wait()
}
