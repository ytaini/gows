/*
 * @Author: zwngkey
 * @Date: 2022-07-25 13:43:55
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-07-25 14:39:13
 * @Description:
 */
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// buf := "abc azc a7c aac 888 a9c  tac"
	// reg := regexp.MustCompile(`a.c`)
	// if reg == nil {
	// 	fmt.Println("regexp err")
	// 	return
	// }
	// // res := reg.Find([]byte(buf))
	// // res := reg.FindAll([]byte(buf), -1)
	// // res := reg.FindAllString(buf, -1)
	// // res := reg.FindAllStringIndex(buf, -1)
	// res := reg.FindAllStringSubmatch(buf, -1)
	// fmt.Println(res)

	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	fmt.Println(searchIn)
	pat := "\\d+[.]\\d+"

	if is, _ := regexp.Match(pat, []byte(searchIn)); is {
		fmt.Println("Match found")
	}
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
