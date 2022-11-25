/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-24 22:27:57
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 23:33:56
 * @Description:
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"zwngkey.cn/golang/go_net/case01/code03/proto"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("listen failed. err: ", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connect failed. err: ", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据:", msg)
	}
}
