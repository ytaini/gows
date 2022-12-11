package znet

import (
	"errors"
	"io"
	"net"

	"wzmiiiiii.cn/zinx/ziface"
)

// Connection IConnection实现类
// 连接模块
type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 连接的id
	ConnID uint32
	// 当前的连接状态
	isClosed bool
	// 告知当前连接已经退出的channel
	ExitChan chan bool

	MsgHandles ziface.IMsgHandle
}

// NewConnection 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHandles ziface.IMsgHandle) *Connection {
	con := &Connection{
		Conn:       conn,
		ConnID:     connID,
		MsgHandles: msgHandles,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
	}
	return con
}

func (c *Connection) StartReader() {
	sugaredLogger.Infof("Reader Goroutine is running...")

	defer sugaredLogger.Infof("ConnID: %d, Connection is exit,remote addr is %v", c.ConnID, c.RemoteAddr())
	defer func(c *Connection) {
		err := c.Stop()
		if err != nil {
			sugaredLogger.Errorf("Conn stop err: %v", err)
		}
	}(c)

	dp := NewDataPack()
	for {
		// 读取数据包消息头字节数据
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			sugaredLogger.Errorf("Read msg head err: %v", err)
			break
		}
		// 解析数据包消息头字节数据
		msg, err := dp.Unpack(headData)
		if err != nil {
			sugaredLogger.Errorf("data unpack err: %v", err)
			break
		}

		// 读取数据包消息体数据
		data := make([]byte, msg.GetMsgLen())
		if _, err = io.ReadFull(c.Conn, data); err != nil {
			sugaredLogger.Errorf("Read msg data err: %v", err)
			break
		}
		// 存放数据包消息体数据
		msg.SetData(data)

		// 得到当前conn的Request请求数据对象
		req := &Request{
			msg:  msg,
			conn: c,
		}

		// 根据msgId找到对应处理函数并执行
		go c.MsgHandles.DoMsgHandler(req)

	}
}

// SendMsg 提供一个SendMsg方法,将要发送给客户端的数据,先进行封包,再发送.
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection closed when send msg")
	}

	// 将data进行封包
	dp := NewDataPack()

	sendData, err := dp.Pack(&Message{
		Id:   msgId,
		Len:  uint32(len(data)),
		Data: data,
	})
	if err != nil {
		sugaredLogger.Errorf("Pack msg err: %v", err)
		return err
	}
	if _, err := c.Conn.Write(sendData); err != nil {
		return err
	}
	return nil
}

func (c *Connection) Start() error {
	sugaredLogger.Infof("Conn: %d Starting ...", c.ConnID)
	// 启动从当前连接的读数据的业务.
	go c.StartReader()
	return nil
}

func (c *Connection) Stop() error {
	// 如果当前连接已经关闭
	if c.isClosed {
		return nil
	}
	sugaredLogger.Infof("Conn: %d is Stopped ...", c.ConnID)
	c.isClosed = true
	err := c.Conn.Close()
	close(c.ExitChan)
	return err
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
