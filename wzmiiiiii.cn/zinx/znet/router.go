package znet

import (
	"wzmiiiiii.cn/zinx/ziface"
)

// BaseRouter 实现router时,先嵌入这个BaseRouter基类,然后根据需要对这个基类的方法进行重写就可以了
type BaseRouter struct{}

// 这里之所以BaseRoute的方法都为空,是因为有的Router不需要有PreHandle或PostHandle这两个方法.
// 所有router继承BaseRouter的好处就是,不需要一定要实现PreHandle或PostHandle.

func (br *BaseRouter) PreHandle(ziface.IRequest) {}

func (br *BaseRouter) Handle(ziface.IRequest) {}

func (br *BaseRouter) PostHandle(ziface.IRequest) {}
