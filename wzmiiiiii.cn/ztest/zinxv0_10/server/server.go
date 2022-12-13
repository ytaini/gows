package main

import (
	"wzmiiiiii.cn/zinx/ziface"
	"wzmiiiiii.cn/zinx/zlog"
	"wzmiiiiii.cn/zinx/znet"
)

/*
	基于zinx框架来开发的 服务器端应用程序
*/

var logger = zlog.GetSugaredLogger()

// PingRouter test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Handle Test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {

	// 先读取客户端的数据.
	logger.Infof("recv client data:")
	logger.Infof("msgID: %d", request.GetMsgId())
	logger.Infof("data: %s", string(request.GetData()))

	// 再回写数据
	if err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping...")); err != nil {
		logger.Errorln(err)
	}
}

// HelloRouter test 自定义路由
type HelloRouter struct {
	znet.BaseRouter
}

// Handle Test Handle
func (pr *HelloRouter) Handle(request ziface.IRequest) {

	// 先读取客户端的数据.
	logger.Infof("recv client data:")
	logger.Infof("\tmsgID: %d", request.GetMsgId())
	logger.Infof("\tdata: %s", string(request.GetData()))

	// 再回写数据
	if err := request.GetConnection().SendMsg(201, []byte("Hello...Hello...Hello...")); err != nil {
		logger.Errorln(err)
	}
}

func main() {
	// 1.创建一个server句柄,使用zinx API
	server := znet.NewServer()

	// 注册Hook函数
	server.RegisterOnCreateConnAfter(func(conn ziface.IConnection) {
		// 给当前连接设置一些属性.
		conn.SetProperty("name", "张三")
		logger.Infof("Set Property TO Connection: [name: %v]", "张三")
	})
	server.RegisterOnDestroyConnBefore(func(conn ziface.IConnection) {
		// 获取当前连接的name属性
		value, _ := conn.GetProperty("name")
		logger.Infof("Get Property: [name: %v]", value)
	})

	// 给当前zinx框架添加一些自定义的router
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})

	// 2.启动server
	server.Server()
}
