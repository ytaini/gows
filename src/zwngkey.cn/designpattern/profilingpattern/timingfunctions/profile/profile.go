/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 18:36:16
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-06 18:58:05
 */
package profile

import (
	"log"
	"time"
)

// When optimizing code, sometimes a quick and dirty time measurement
// is required as opposed to utilizing profiler tools/frameworks to validate assumptions.

// Time measurements can be performed by utilizing time package and defer statements.

// 记录某个函数的执行时间.
func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	log.Printf("%s lasted %s", name, elapsed)
}
