/*
 * @Author: zwngkey
 * @Date: 2022-05-14 02:39:03
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-16 03:34:53
 * @Description:
 */
package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	IP        string
	Port      int
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	Message   chan string
}

// Server构造函数
func NewServer(ip string, port int) *Server {
	return &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}
func (s *Server) Broadcast(u *User, msg string) {
	sendMsg := "[" + u.Addr + "]" + u.Name + ": " + msg
	s.Message <- sendMsg
}

func (s *Server) handler(conn net.Conn) {
	// log.Println("连接建立成功!")

	user := NewUser(conn, s)

	user.Online()

	isLive := make(chan struct{})

	go func() {
		buf := make([]byte, 10*1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				break
			}
			if err != nil && err != io.EOF {
				log.Println("读取消息失败")
				return
			}
			user.DealMsg(string(buf[:n-1]))
			isLive <- struct{}{}
		}
	}()

	for {
		//超时控制.
		select {
		//这个case放在上面可以重置计时器
		case <-isLive:
		case <-time.After(1 * time.Minute):
			user.SendMsg("提示:你被踢了...")
			close(user.Ch)
			conn.Close()
			return
		}
	}
}

func (s *Server) ListenMsg() {
	for msg := range s.Message {
		s.mapLock.RLock()
		for _, user := range s.OnlineMap {
			user.Ch <- msg
		}
		s.mapLock.RUnlock()
	}
}

func (s *Server) Run() {
	listener, err := net.Listen("tcp", s.IP+":"+strconv.Itoa(s.Port))

	if err != nil {
		log.Fatalln("Listen err :", err)
	}

	defer listener.Close()

	go s.ListenMsg()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Listen accept err:", err)
			continue
		}
		go s.handler(conn)
	}

}
