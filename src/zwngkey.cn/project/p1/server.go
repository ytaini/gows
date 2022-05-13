/*
 * @Author: zwngkey
 * @Date: 2022-05-14 02:39:03
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-14 02:54:42
 * @Description:
 */
package main

import (
	"log"
	"net"
	"strconv"
)

type Server struct {
	IP   string
	Port int
}

// Server构造函数
func NewServer(ip string, port int) *Server {
	return &Server{
		ip,
		port,
	}
}

func (s *Server) handler(conn net.Conn) {
	log.Println("连接建立成功!")
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.IP+":"+strconv.Itoa(s.Port))

	if err != nil {
		log.Fatalln("Listen err :", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Listen accept err:", err)
			continue
		}

		go s.handler(conn)

	}

}
