/*
 * @Author: zwngkey
 * @Date: 2021-12-09 17:16:13
 * @LastEditTime: 2022-11-21 10:55:54
 * @Description:
 */
package main

import (
	"fmt"
	"log"

	"zwngkey.cn/gobasic/day02/04copy_file_lianxi/copyfile"
)

func main() {
	ok, err := copyfile.CopyFile("main.go", "maincopy.go")
	if !ok || err != nil {
		log.Fatalln("拷贝失败!! err:", err)
	}
	fmt.Println("copy done")
}
