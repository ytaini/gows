package znet

import (
	"fmt"

	"wzmiiiiii.cn/zinx/ziface"
)

// 消息处理模块的实现

type MsgHandle struct {
	// 存放每个MsgID所对应的处理方法
	routers map[uint32]ziface.IRouter
}

func NewMsgHandle() ziface.IMsgHandle {
	return &MsgHandle{
		routers: make(map[uint32]ziface.IRouter),
	}
}

func (m *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	router := m.routers[request.GetMsgId()]
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
