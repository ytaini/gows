/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 23:36:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 00:41:01
 */
package type4

import (
	"fmt"
	"sync"
)

// Go语言惯用的单例模式(懒汉式)

// sync.Once实现并发安全的单例
type singleton struct{}

func (s *singleton) Print() {
	fmt.Println("asd")
}

var once sync.Once
var instance *singleton

func GetInstance() *singleton {
	// 使用sync.Once类型来保证某个函数只被调用一次
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
