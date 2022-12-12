package ziface

// IServer 定义一个服务器接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Server 运行服务器
	Server()
	// AddRouter 路由功能: 给当前的服务注册一个路由方法,供客户端的连接处理使用
	AddRouter(uint32, IRouter)
	// GetConnMgr 获取当前Server的连接管理器
	GetConnMgr() IConnManager
	// RegisterOnCreateConnAfter 注册OnCreateConnAfter Hook函数
	RegisterOnCreateConnAfter(func(conn IConnection))
	// RegisterOnDestroyConnBefore 注册OnDestroyConnBefore Hook函数
	RegisterOnDestroyConnBefore(func(conn IConnection))
	// CallOnCreateConnAfter 调用OnCreateConnAfter Hook函数
	CallOnCreateConnAfter(conn IConnection)
	// CallOnDestroyConnBefore 调用注册OnDestroyConnBefore Hook函数
	CallOnDestroyConnBefore(conn IConnection)
}
