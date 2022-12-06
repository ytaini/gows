/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 23:28:47
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-06 23:41:39
 */
package type3

import (
	"sync"
	"sync/atomic"
)

// 懒汉式 (并发安全)
type singleton struct{}

var instance *singleton

// 通过使用sync/atomic这个包，我们可以原子化加载并设置一个标志，该标志表明我们是否已初始化实例。
var initialized uint32

var mu sync.Mutex

func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 { //原子操作
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = &singleton{}
		atomic.AddUint32(&initialized, 1)
	}
	return instance
}
