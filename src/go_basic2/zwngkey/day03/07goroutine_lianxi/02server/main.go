package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		// handleConn(conn) //此时服务器同一时刻只能处理一个客户端连接.
		go handleConn(conn)

	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)

	}
}
