package znet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"wzmiiiiii.cn/zinx/zcfg"
	"wzmiiiiii.cn/zinx/ziface"
)

// 封包,拆包的具体模块

type DataPack struct{}

func NewDataPack() ziface.IDataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	// 消息头: 4个字节存消息长度,4个字节存消息ID
	return 8
}

func (d *DataPack) Pack(message ziface.IMessage) ([]byte, error) {
	buf := new(bytes.Buffer)
	// 将dataLen写入
	if err := binary.Write(buf, binary.LittleEndian, message.GetMsgLen()); err != nil {
		return nil, err
	}
	// 将MsgId写入
	if err := binary.Write(buf, binary.LittleEndian, message.GetMsgId()); err != nil {
		return nil, err
	}
	// 将data数据写入
	if err := binary.Write(buf, binary.LittleEndian, message.GetData()); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (d *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	reader := bytes.NewReader(data)

	msg := &Message{}

	// 先读取dataLen
	if err := binary.Read(reader, binary.LittleEndian, &msg.Len); err != nil {
		return nil, err
	}

	// 在读取dataId
	if err := binary.Read(reader, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	// 判断数据包长度是否已经超出了允许的最大包长度
	maxSize := zcfg.Config.MaxPackageSize
	if maxSize > 0 && msg.Len > maxSize {
		return nil, errors.New("too large msg data recv")
	}
	return msg, nil
}
