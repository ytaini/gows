package main

import (
	"log"
	"wzmiiiiii.cn/zinx/ziface"
	"wzmiiiiii.cn/zinx/znet"
)

/*
	基于zinx框架来开发的 服务器端应用程序
*/

// PingRouter ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// PreHandle Test PerHandle
func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	log.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ...\n"))
	if err != nil {
		log.Println("call back before ping error:", err)
	}
}

// Handle Test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {
	log.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ...ping ...ping ...\n"))
	if err != nil {
		log.Println("call back ping ... error:", err)
	}
}

// PostHandle Test PostHandle
func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	log.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping ...\n"))
	if err != nil {
		log.Println("call back after ping error:", err)
	}
}

func main() {
	// 1.创建一个server句柄,使用zinx API
	server := znet.NewServer("[zinx V0.1]")

	// 给当前zinx框架添加一个自定义的router
	server.AddRouter(&PingRouter{})

	// 2.启动server
	server.Server()
}
