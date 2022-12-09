/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 07:06:14
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 07:19:18
 */
package main

import (
	"fmt"

	. "zwngkey.cn/designpattern/behavioral/proxy/static/code01"
)

func main() {
	// NewLogProxyCar(&Car{}).Move()
	NewTimerProxy(NewLogProxy(&Car{})).Move()

	fmt.Println("---------------")

	NewLogProxy(NewTimerProxy(&Car{})).Move()
}
