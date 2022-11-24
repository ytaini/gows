/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-20 10:03:47
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-22 09:46:11
 * @Description:
 */
package code

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	var arr [3]int
	// fmt.Println(arr == nil) 报错.数组为值类型.
	fmt.Println(arr)
}

func Test2(t *testing.T) {
	fmt.Printf("%08b\n", 1)
}

func Test3(t *testing.T) {
	// 这样解析出来的时间是UTC时间.
	// tm, err := time.Parse("2006-01-02 15:04:05", "2022-10-12 10:12:13")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(tm)

	// 这样解析出来的时间才是北京时间.
	loc, err := time.LoadLocation("Asis/Shanghai")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	tm, err1 := time.ParseInLocation("2006-01-02 15:04:05", "2022-10-12 10:12:13", loc)
	if err1 != nil {
		fmt.Printf("err: %v\n", err1)
		return
	}
	fmt.Printf("fmDate: %v\n", tm)
}

func Test4(t *testing.T) {
	arr := [...]int{1, 2, 2}
	s := arr[:]
	arr1 := (*[4]int)(s)
	fmt.Println(arr1)
}
