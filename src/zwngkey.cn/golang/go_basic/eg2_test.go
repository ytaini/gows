/*
 * @Author: zwngkey
 * @Date: 2022-05-11 20:59:24
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:35:30
 * @Description:
	import本地的module需要借助replace指令来实现。
*/
package gobasic

import (
	"fmt"
	"testing"

	"zwngkey.cn/dsaa/util"
)

func TestEg1(t *testing.T) {
	fmt.Printf("util.Max(1, 2): %v\n", util.Max(1, 2))

}
