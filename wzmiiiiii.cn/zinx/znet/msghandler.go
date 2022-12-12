package znet

import (
	"fmt"

	"wzmiiiiii.cn/zinx/zcfg"

	"wzmiiiiii.cn/zinx/ziface"
)

// 消息处理模块的实现

type MsgHandle struct {
	// 存放每个MsgID所对应的处理方法
	routers map[uint32]ziface.IRouter

	// 负责Worker取任务的消息队列
	TaskQueue []chan ziface.IRequest
	// 业务工作Worker池的goroutine数量
	WorkerPoolSize uint32
}

func NewMsgHandle() ziface.IMsgHandle {
	return &MsgHandle{
		routers:        make(map[uint32]ziface.IRouter),
		WorkerPoolSize: zcfg.Config.WorkerPoolSize, // 从zcfg.GlobalConfig中获取
		TaskQueue:      make([]chan ziface.IRequest, zcfg.Config.WorkerPoolSize),
	}
}

func (m *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	router, ok := m.routers[request.GetMsgId()]
	if !ok {
		sugaredLogger.Errorf("MsgId: %d is NOT FOUND! Need Register!!!", request.GetMsgId())
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (m *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	// 当前msg绑定的处理方法是否已经存在
	if _, ok := m.routers[msgId]; ok {
		panic(fmt.Sprintf("repeat router, msgId:%d", msgId))
	}
	m.routers[msgId] = router
}

// StartWorkerPool 启动一个Worker工作池(只能启动一次)
func (m *MsgHandle) StartWorkerPool() {
	// 根据workerPoolSize 依次开启Worker
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		// 给当前worker对应的消息队列开辟空间.
		m.TaskQueue[i] = make(chan ziface.IRequest, zcfg.Config.MaxWorkerTaskLen)
		// 启动Worker
		go m.StartWorker(i)
	}
}

// StartWorker StartOneWorker 启动一个Worker
func (m *MsgHandle) StartWorker(workerId int) {
	sugaredLogger.Infof("Worker%v is starting...", workerId)
	for {
		select {
		// 队列中有消息
		case request := <-m.TaskQueue[workerId]:
			sugaredLogger.Infof("TaskQueue%d recv request,msgId: %d", workerId, request.GetMsgId())
			m.DoMsgHandler(request)
		}
	}
}

// SendMsgToTaskQueue 将消息发送到某个消息队列
func (m *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// 将消息平均分配给不同的消息队列. 根据connId来进行分配
	workerId := request.GetConnection().GetConnID() % m.WorkerPoolSize
	sugaredLogger.Infof("Send data[msgId:%d] to TaskQueue%d", request.GetMsgId(), workerId)
	// 发送
	m.TaskQueue[workerId] <- request
}
