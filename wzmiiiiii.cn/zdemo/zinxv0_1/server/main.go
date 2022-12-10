package main

import "wzmiiiiii.cn/zinx/znet"

/*
	基于zinx框架来开发的 服务器端应用程序
*/

func main() {
	// 1.创建一个server句柄,使用zinx API
	server := znet.NewServer("[zinx V0.1]")
	// 2.启动server
	server.Server()
}
