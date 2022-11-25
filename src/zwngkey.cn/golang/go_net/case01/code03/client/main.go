/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-24 23:14:19
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 23:33:58
 * @Description:
 */
package main

import (
	"fmt"
	"net"

	"zwngkey.cn/golang/go_net/case01/code03/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("dial failed.err: ", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed. err:", err)
			return
		}
		conn.Write(data)
	}
}
