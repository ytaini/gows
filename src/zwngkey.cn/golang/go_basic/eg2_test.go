/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-27 23:07:55
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
