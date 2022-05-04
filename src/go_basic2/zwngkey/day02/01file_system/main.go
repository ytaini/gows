package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./main.go")

	if err != nil {
		log.Fatalln("文件打开失败")
		return
	}

	defer file.Close()

	temp := make([]byte, 128) //指定每次读取数据的大小
	for {
		n, err := file.Read(temp)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("文件读取失败2")
			os.Exit(1)
		}
		fmt.Printf("读取了: %d个字符\n", n)

		fmt.Println(string(temp[:n]))

	}

}
