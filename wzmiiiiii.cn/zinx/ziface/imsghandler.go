package ziface

// 消息管理抽象层

type IMsgHandle interface {
	// DoMsgHandler 调度对应的Router消息出来方法
	DoMsgHandler(IRequest)
	// AddRouter 为消息添加具体的处理逻辑
	AddRouter(uint32, IRouter)
	// StartWorkerPool 启动Worker工作池
	StartWorkerPool()
	// SendMsgToTaskQueue 将消息发送到某个消息队列
	SendMsgToTaskQueue(IRequest)
}
