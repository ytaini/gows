package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func writeFileByWrite() {
	file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		log.Fatalln(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("你好嘻嘻嘻嘻go\n"))
	_, err = file.WriteString("xixxi")
	// err = file.Truncate(0)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("写入成功")
}

func writeFileBybufio() {
	file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("123") //将数据写入缓存
	writer.Flush()                     //将缓存写入文件
}

func writeFileByioutil() {
	err := ioutil.WriteFile("./test.txt", []byte("ahahahah"), 0644)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
func main() {
	// writeFileByWrite()
	// writeFileBybufio()
	writeFileByioutil()
}
