package ziface

/*
	连接管理模块抽象层
*/

type IConnManager interface {
	// Add 添加连接
	Add(connection IConnection)
	// Remove 删除连接
	Remove(connId uint32)
	// Get 根据connID获取连接
	Get(connId uint32) (IConnection, error)
	// GetSize 得到当前连接总数
	GetSize() int
	// ClearAll 清除并终止所有连接
	ClearAll()
}
