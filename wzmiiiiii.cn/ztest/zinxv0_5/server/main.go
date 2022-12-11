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

// PingRouter ping test 自定义路由
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
	if err := request.GetConnection().SendMsg(request.GetMsgId()+10000, []byte("ping...ping...ping...")); err != nil {
		logger.Errorln(err)
	}
}

func main() {
	// 1.创建一个server句柄,使用zinx API
	server := znet.NewServer()

	// 给当前zinx框架添加一个自定义的router
	server.AddRouter(&PingRouter{})

	// 2.启动server
	server.Server()
}
