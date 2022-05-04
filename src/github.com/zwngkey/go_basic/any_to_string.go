package gobasic

import (
	"fmt"
	//"strconv"
)

func Test() {

	//Sprintf格式化字符串,把指定的数据类型转成我要的字符串
	//Spirntf()会返转换后的字符串
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string // 空的str

	// 使用第一种方式转换 fmt.Sprintf方法
	//把int整数,转为string
	str = fmt.Sprintf("%d\n", num1)
	fmt.Printf("str type %T str=%v", str, str)

	//把小数转为string
	str = fmt.Sprintf("%f\n", num2)
	fmt.Printf("str type %T str=%v", str, str)

	//把bool转为string
	str = fmt.Sprintf("%t\n", b)
	fmt.Printf("str type %T str=%v", str, str)

	//把字符类型byte转为string
	str = fmt.Sprintf("%c\n", mychar)
	fmt.Printf("str type %T str=%v", str, str)
}
