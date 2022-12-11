package ziface

// 封包,拆包,模块
// 直接面向TCP连接中的数据流,用于处理TCP粘包问题.

type IDataPack interface {
	// Pack 封包
	Pack(IMessage) ([]byte, error)
	// Unpack 拆包
	Unpack([]byte) (IMessage, error)
	// GetHeadLen 获取消息头的字节数
	GetHeadLen() uint32
}
