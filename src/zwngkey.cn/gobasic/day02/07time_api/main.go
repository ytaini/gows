/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 11:36:14
 * @Description:
 */
package main

import (
	"fmt"
	"time"
)

func f1() {
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
	// fmDate, err := time.Parse(formatStr, dateString)

	z, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

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

func main() {
	f1()

	time.Sleep(5 * time.Second)

	timer := time.Tick(time.Second)
	for i := range timer {
		fmt.Println(i)
	}

}
