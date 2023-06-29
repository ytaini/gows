/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2023-06-30 00:19:19
 */
package main

import (
	"embed"
	"log"
)

// 这是一个指令,而不是普通的注释
//
//go:embed test/*
var f embed.FS

func main() {
	data, err := f.ReadFile("test/1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	print(string(data))
	// dirs, _ := f.ReadDir("test")
	// for _, dir := range dirs {
	// 	fmt.Printf("dir.Name(): %v\n", dir.Name())
	// }
}
