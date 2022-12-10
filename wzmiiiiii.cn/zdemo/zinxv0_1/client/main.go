package main

import (
	"log"
	"net"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", ":8999")
	if err != nil {
		log.Println("connect server err:", err)
		return
	}
	defer conn.Close()

	for {
		_, err := conn.Write([]byte("hello, zinxV0.1!"))
		if err != nil {
			log.Println("client send data err:", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Println("Client recv server data err:", err)
			continue
		}
		log.Printf("Client recv server data: %s,cnt:%d", string(buf[:cnt]), cnt)

		time.Sleep(1 * time.Second)
	}

}
