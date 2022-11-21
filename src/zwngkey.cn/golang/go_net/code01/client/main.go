package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.103:8080")
	if err != nil {
		log.Fatalln("conn err: ", err)
	}
	defer conn.Close()

	ip := conn.RemoteAddr()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("请输入发生给服务器的内容: ")
		line, err := reader.ReadString('\n')
		if strings.TrimSpace(line) == "exit" {
			break
		}

		if err != nil {
			log.Fatalln("read err: ", err)
		}
		conn.Write([]byte(line))

		data := make([]byte, 1024)

		n, err := conn.Read(data)
		if err != nil {
			log.Fatalln("read err: ", err)
		}

		fmt.Println("服务端:", ip, ":", string(data[:n]))
	}
}
