/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-16 06:20:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 10:34:25
 * @Description:
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func readFile1() {
	file, err := os.OpenFile("logger.go", os.O_RDONLY, 0)

	if err != nil {
		log.Fatalln("打开文件失败 err", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%q\n", line)
		// fmt.Print(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("读取文件失败")
			return
		}
	}
}

func readFile2() {
	content, err := os.ReadFile("./logger.go")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	readFile1()
	fmt.Println("---------")
	readFile2()
}
