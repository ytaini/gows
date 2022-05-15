/*
 * @Author: zwngkey
 * @Date: 2022-05-14 06:48:04
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-15 08:31:36
 * @Description:
	解析配置文件
*/
package main

import (
	"fmt"
)

func main() {
	var mc Config
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		panic(err)
	}
	fmt.Println(mc)
}
