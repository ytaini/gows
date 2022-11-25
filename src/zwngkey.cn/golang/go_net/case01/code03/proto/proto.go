/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-24 23:17:46
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 23:44:09
 * @Description:
 */
package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// 封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾”内容)。
// 包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。

// 我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。

// Encode 将消息编码
func Encode(massage string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(massage))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(massage))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lenByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lenBuf := bytes.NewBuffer(lenByte)
	var len int32
	err := binary.Read(lenBuf, binary.LittleEndian, &len)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < len+4 {
		return "", err
	}
	pack := make([]byte, int(4+len))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
