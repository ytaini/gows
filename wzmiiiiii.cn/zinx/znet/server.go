package znet

import (
	"fmt"
	"io"
	"log"
	"net"
	"wzmiiiiii.cn/zinx/ziface"
)

// Server IServer的接口实现,定义一个Server的服务器模块
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的Ip
	IP string
	// 服务器监听的Port
	Port int
}

func NewServer(name string) ziface.IServer {
	server := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return server
}

func (s *Server) Start() {
	log.Printf("[Start] Server Listener at IP: %s, Port: %d, is starting\n", s.IP, s.Port)

	go func() {
		// 1. 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Fatalln("Resolve tcp addr error: ", err)
		}
		// 2. 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)

		if err != nil {
			log.Fatalf("Listen %s,err: %v", s.IPVersion, err)
		} else {
			defer func(listener *net.TCPListener) {
				err := listener.Close()
				if err != nil {
					log.Println("listener close err:", err)
				}
			}(listener)
		}

		log.Printf("[Success] Start Zinx server success,serverName: %s, Listening...", s.Name)

		// 3. 阻塞的等待客户端连接,处理客户端连接业务(读写)
		for {
			// 如果有客户端连接过来,阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Printf("Accept err: %v", err)
				continue
			}

			log.Printf("Client Addr:%s connect...", conn.RemoteAddr())

			// 已经与客户端建立连接,处理业务.
			go process(conn)
		}
	}()
}

func process(conn *net.TCPConn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("Client Addr:%s is leaving.. ", conn.RemoteAddr())
			} else {
				log.Println("recv buf err:", err)
			}
			return
		}

		log.Printf("Server read client info: %s,cnt:%d", string(buf[:cnt]), cnt)

		if _, err := conn.Write(buf[0:cnt]); err != nil {
			log.Printf("Server Write back buf err: %v", err)
			continue
		}
	}
}

func (s *Server) Stop() {
	// TODO 将一些服务器的资源,状态或者一些已经开辟的连接信息 进行停止或者回收
}

func (s *Server) Server() {
	// 启动server 的服务功能
	s.Start()

	// TODO 做一些启动服务器之后的额外业务.

	// 阻塞状态
	select {}
}
