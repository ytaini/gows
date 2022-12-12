package znet

import (
	"fmt"
	"net"

	"wzmiiiiii.cn/zinx/zcfg"
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
	// Server的消息管理模块,用来绑定MsgId和对应的处理函数
	MsgHandles ziface.IMsgHandle
	// 连接管理器
	ConnMgr ziface.IConnManager
	// 创建连接之后自动调用的Hook函数
	OnCreateConnAfter func(conn ziface.IConnection)
	// 销毁连接之前自动调用的Hook函数
	OnDestroyConnBefore func(conn ziface.IConnection)
}

const IPVersion = "tcp4"

func NewServer() ziface.IServer {
	server := &Server{
		Name:       zcfg.Config.Name,
		IPVersion:  IPVersion,
		IP:         zcfg.Config.Host,
		Port:       zcfg.Config.TcpPort,
		MsgHandles: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return server
}

func (s *Server) Start() {
	sugaredLogger.Infof("[Start] Server Listener at IP: %s, Port: %d, is starting", s.IP, s.Port)

	go func() {
		// 开启worker pool.
		s.MsgHandles.StartWorkerPool()

		// 1. 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			sugaredLogger.DPanicln("Resolve tcp addr error: ", err)
			return
		}
		// 2. 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)

		if err != nil {
			sugaredLogger.DPanicln("err: %v", err)
			return
		} else {
			defer listener.Close()
		}

		sugaredLogger.Infof("[Success] Start Zinx server success,serverName: %s, Listening...", s.Name)

		var cid uint32 = 0

		// 3. 阻塞的等待客户端连接,处理客户端连接业务(读写)
		for {
			// 如果有客户端连接过来,阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				sugaredLogger.Errorf("Accept err: %v", err)
				continue
			}

			// 是否超过最大连接数.
			if s.ConnMgr.GetSize() >= zcfg.Config.MaxConn {
				// TODO 给客户端响应一个超出最大连接的错误
				sugaredLogger.Errorln("超出最大连接...")
				conn.Close()
				continue
			}

			sugaredLogger.Infof("[%04d] Client Addr:%s connect...", cid, conn.RemoteAddr())

			// 将处理连接的业务方法和conn进行绑定 得到我们的连接模块.
			myConn := NewConnection(s, conn, cid, s.MsgHandles)
			cid++

			// 启动当前的连接业务处理.
			go myConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	// TODO 将一些服务器的资源,状态或者一些已经开辟的连接信息 进行停止或者回收

	s.ConnMgr.ClearAll()
}

func (s *Server) Server() {
	// 启动server 的服务功能
	s.Start()

	// TODO 做一些启动服务器之后的额外业务.

	// 阻塞状态
	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandles.AddRouter(msgId, router)
	sugaredLogger.Infoln("Add Router Success...")
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) RegisterOnCreateConnAfter(f func(conn ziface.IConnection)) {
	s.OnCreateConnAfter = f
}

func (s *Server) RegisterOnDestroyConnBefore(f func(conn ziface.IConnection)) {
	s.OnDestroyConnBefore = f
}

func (s *Server) CallOnCreateConnAfter(conn ziface.IConnection) {
	if s.OnCreateConnAfter != nil {
		sugaredLogger.Infof("OnCreateConnAfter exec...")
		s.OnCreateConnAfter(conn)
	}
}

func (s *Server) CallOnDestroyConnBefore(conn ziface.IConnection) {
	if s.OnDestroyConnBefore != nil {
		sugaredLogger.Infof("OnDestroyConnBefore exec...")
		s.OnDestroyConnBefore(conn)
	}
}
