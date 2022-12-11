package ziface

// 将请求的消息封装到一个message中,定义抽象的接口

type IMessage interface {
	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte
	SetMsgId(uint32)
	SetData([]byte)
	SetMsgLen(uint32)
}
