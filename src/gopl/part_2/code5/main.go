package main

import (
	"fmt"
	"strconv"
)

func main() {
	// s := "你好吗?"
	// a := "世界"
	// b := "\xe4\xb8\x96\xe7\x95\x8c"
	// c := "\u4e16\u754c"
	// d := "\U00004e16\U0000754c"
	// e := '\x41'
	// f := '\u4e16'
	// g := "\xe4\xb8\x96"
	// fmt.Println(len(s))     //10
	// fmt.Println(s[1], s[2]) //189 160
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(string(e))
	// fmt.Println(string(f))
	// fmt.Println(string(g))

	for i, r := range "hello,时间" {
		fmt.Printf("%d,%q,%[2]d\n", i, r) //%q 带双引号的字符串"abc"或带单引号的字符'c'
		/*输出:
		0,'h',104
		1,'e',101
		2,'l',108
		3,'l',108
		4,'o',111
		5,',',44
		6,'时',26102
		9,'间',38388
		*/
	}

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	x, _ = strconv.Atoi("123")
	z, _ := strconv.ParseInt("123123131123", 10, 8)
	fmt.Printf("%d,%[1]T\n", x)
	fmt.Printf("%d,%[1]b,%[1]T", z)
}
