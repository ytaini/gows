package ziface

/*
	路由抽象接口
	路由里的数据都是IRequest
*/

type IRouter interface {
	// PreHandle 处理conn之前的钩子方法hook
	PreHandle(IRequest)
	// Handle 处理conn的主方法
	Handle(IRequest)
	// PostHandle 处理conn之后的钩子方法hook
	PostHandle(IRequest)
}
