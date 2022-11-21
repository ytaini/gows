/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 11:23:33
 * @Description:
 */
package main

import (
	"log"
	"os"
)

func init() {
	file, _ := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetPrefix("[zwngkey] ")
	log.SetOutput(file)
}

func main() {

	log.Println("普通日志")
	log.Printf("putong%s\n", "日志")
	log.Fatalln("123")
}
