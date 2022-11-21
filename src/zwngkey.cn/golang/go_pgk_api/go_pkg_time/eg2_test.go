/*
 * @Author: zwngkey
 * @Date: 2022-05-05 18:02:42
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-05 18:48:24
 * @Description: time包的api
 */
package gopkgtime

import (
	"fmt"
	"testing"
	"time"
)

/*
Since(t)函数:

	返回从t开始经过的时间。它是time.Now().Sub(t)的简写。
*/
func TestEg21(t *testing.T) {
	startTime := time.Now()
	_ = time.Since(startTime)
}

func TestEg22(t *testing.T) {
	now := time.Now()
	y, m, d := now.Date()
	fmt.Printf("y: %v\n", y)
	fmt.Printf("m: %v\n", m)
	fmt.Printf("d: %v\n", d)

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Printf("timeStamp1: %v\n", timeStamp1)
	fmt.Printf("timeStamp2: %v\n", timeStamp2)

	formatStr := "2006-01-02 15:04:05"
	formatDate := now.Format(formatStr)
	fmt.Printf("formatDate: %v\n", formatDate)

	dateString := "2021-12-12 12:03:22"

	// fmDate, _ := time.Parse(formatStr, dateString)
	// fmt.Println(fmDate)

	z, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(z, "-----")

	fmDate, err1 := time.ParseInLocation(formatStr, dateString, z)
	if err1 != nil {
		fmt.Printf("err: %v\n", err1)
		return
	}
	fmt.Printf("fmDate: %v\n", fmDate)

	afterOneHour := now.Add(1 * time.Hour)
	fmt.Printf("afterOneHour: %v\n", afterOneHour)

	dateDur := fmDate.Sub(now)
	fmt.Printf("dateDur: %v\n", dateDur)
}
