/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-25 00:01:20
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-25 00:30:01
 * @Description:
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", ":9000")
	if err != nil {
		fmt.Println("dial failed.err:", err)
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte("hello server"))
	if err != nil {
		fmt.Println("send data failed. err:", err)
		return
	}
	data := make([]byte, 10240)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("recv data failed. err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), conn.RemoteAddr(), n)

	// na := &net.UDPAddr{
	// 	IP:   net.IPv4(0, 0, 0, 0),
	// 	Port: 9000,
	// }
	// udpConn, err := net.DialUDP("udp", nil, na)
	// if err != nil {
	// 	fmt.Println("dial failed.err:", err)
	// 	return
	// }
	// defer udpConn.Close()

	// _, err = udpConn.Write([]byte("udp,udp,udp,udp"))
	// if err != nil {
	// 	fmt.Println("send data failed. err:", err)
	// 	return
	// }
	// data := make([]byte, 10240)
	// n, addr, err := udpConn.ReadFrom(data)
	// // n, addr, err := udpConn.ReadFromUDP(data)
	// // n, addr, err := udpConn.ReadFromUDPAddrPort(data)

	// if err != nil {
	// 	fmt.Println("recv data failed. err:", err)
	// 	return
	// }
	// fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), addr, n)
}
