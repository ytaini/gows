/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 11:14:08
 * @Description:
 */
package main

import (
	"fmt"
	"log"
	"os"
)

// insertContentInText 在文件某个位置添加指定内容
// src: 源文件路径
// insertSrc: 要添加的内容
// localtion: 从指定位置开始插入
func InsertContentInText(src string, insertSrc string, localtion int64) {
	file, err := os.Open(src)
	if err != nil {
		log.Fatalf("open %s file failed,err: %v\n", src, err)
		return
	}

	file1, err1 := os.OpenFile("temp", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err1 != nil {
		log.Fatalf("err: %v\n", err)
	}

	tem := make([]byte, localtion-1)

	_, err2 := file.Read(tem)

	if err2 != nil {
		log.Fatalf("err: %v\n", err2)
	}
	tem = append(tem, []byte(insertSrc)...)

	_, err3 := file1.Write(tem[:])

	if err3 != nil {
		log.Fatalf("err: %v\n", err3)
	}

	n, err4 := file.Read(tem)

	if err4 != nil {
		log.Fatalf("err: %v\n", err3)
	}

	_, err5 := file1.Write(tem[:n])

	if err5 != nil {
		log.Fatalf("err: %v\n", err3)
	}
	err6 := os.Remove(src)
	if err6 != nil {
		log.Fatalf("err: %v\n", err3)
	}
	os.Rename("./temp", src)
	fmt.Println("文件添加内容成功")
}

func InsertContentInText2(src string, insertSrc string, localtion int64) {
	file, err := os.OpenFile(src, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("open %s file failed,err: %v\n", src, err)
	}

	tem := make([]byte, localtion-1)

	_, err2 := file.Read(tem)
	if err2 != nil {
		log.Fatalf("err: %v\n", err2)
	}
	tem = append(tem, []byte(insertSrc)...)

	tem2 := make([]byte, 1024)
	n, err3 := file.Read(tem2)
	if err3 != nil {
		log.Fatalf("err: %v\n", err3)
	}
	tem = append(tem, tem2[:n]...)

	// file.Truncate(0)
	file.Seek(0, 0)
	_, err4 := file.Write(tem[:])
	if err4 != nil {
		log.Fatalf("err: %v\n", err4)
	}

}

func main() {
	// InsertContentInText("./test", "123", 5)
	InsertContentInText2("./test", "123", 5)
}
