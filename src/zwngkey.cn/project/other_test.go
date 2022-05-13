/*
 * @Author: zwngkey
 * @Date: 2022-05-14 04:42:35
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-14 05:00:51
 * @Description:
 */
package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	s := fmt.Sprintf("%04d", 12)
	fmt.Println(s)
}
