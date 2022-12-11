package znet

// Message 将请求的消息封装到一个message中
type Message struct {
	// 消息的Id
	Id uint32
	// 消息的长度
	Len uint32
	// 消息的内容
	Data []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{Id: id, Data: data, Len: uint32(len(data))}
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.Len
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetMsgLen(len uint32) {
	m.Len = len
}
