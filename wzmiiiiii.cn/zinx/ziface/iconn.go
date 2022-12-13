package ziface

import "net"

// IConnection 定义连接模块的抽象层
type IConnection interface {
	// Start 启动连接
	Start() error

	// Stop 停止连接
	Stop() error

	// GetTCPConnection 获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前连接模块的连接id
	GetConnID() uint32

	// RemoteAddr 获取远程客户端的TCP状态 IP Port
	RemoteAddr() net.Addr

	// SendMsg 发送数据 将数据发送给远程的客户端.
	SendMsg(uint32, []byte) error

	// SetProperty 设置连接属性
	SetProperty(key string, val any)

	// GetProperty 获取连接属性
	GetProperty(key string) (any, error)

	// RemoveProperty 移除连接属性
	RemoveProperty(key string)
}
