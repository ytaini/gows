package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("服务器开始监听8080端口....")

	listener, err := net.Listen("tcp", "0.0.0.0:8080")

	if err != nil {
		log.Fatalln("listen end. err: ", err)
	}

	defer listener.Close()

	for {
		fmt.Println("等待客户端连接...")
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Accept err: ", err)
			continue
		}
		fmt.Println("Accept success,conn: ", conn.RemoteAddr())

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	ip := conn.RemoteAddr()

	for {
		data := make([]byte, 1024)
		fmt.Println(ip, "等待客服端发送信息...")
		//read方法会等待客户端发送信息,若客户端未发送信息read方法会阻塞.
		n, err := conn.Read(data)
		if err != nil {
			// log.Fatalln("err: ", err)
			fmt.Println("客户端", ip, "已断开连接")
			break
		}
		fmt.Println("客户端:", ip, ":", string(data[:n]))
		rand.Seed(time.Now().UnixNano())
		i := rand.Int63()
		s := strconv.Itoa(int(i))
		fmt.Println("发送给", ip, "的信息为", s)
		conn.Write([]byte(s))
	}
}
