/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-24 23:51:41
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-25 00:01:11
 * @Description:
 */

package main

import (
	"fmt"
	"net"
)

func main() {
	// listener, err := net.ListenUDP("udp", &net.UDPAddr{
	// 	IP:   net.IPv4(0, 0, 0, 0),
	// 	Port: 9000,
	// })
	listener, err := net.ListenPacket("udp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("listen failed. err: ", err)
		return
	}
	defer listener.Close()
	for {
		var data [1024]byte
		n, addr, err := listener.ReadFrom(data[:])
		if err != nil {
			fmt.Println("read udp failed. err: ", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listener.WriteTo(data[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed. err: ", err)
			continue
		}
	}
}
