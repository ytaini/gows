/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 22:43:25
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-06 23:02:30
 */
package test

import (
	"fmt"
	"sync"
	"testing"

	"zwngkey.cn/designpattern/singleton/type1"
)

func Test(t *testing.T) {
	instance := type1.New()
	instance["hello"] = "world"

	instance1 := type1.New()
	fmt.Println(instance1)
}

func Test1(t *testing.T) {
	instance := type1.New1()
	instance1 := type1.New1()
	fmt.Println(instance == instance1) //true
}

func Test2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			instance := type1.New1()
			fmt.Printf("%p\n", instance)
		}()
	}
	wg.Wait()
}
