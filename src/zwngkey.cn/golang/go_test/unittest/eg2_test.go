/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-08 02:13:02
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-08 04:36:28
 */
package unit_test

import (
	"testing"
)

func Test(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}
	t.Log(123)
}
