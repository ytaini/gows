package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func readFile1() {
	file, err := os.OpenFile("./main.go", os.O_RDONLY, 0)

	if err != nil {
		log.Fatalln("打开文件失败 err", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("读取文件失败")
			return
		}

		fmt.Print(line)

	}
}

func readFile2() {

	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	readFile2()
}
