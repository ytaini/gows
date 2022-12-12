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
	// 当前Conn属于哪个Server
	Server ziface.IServer
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 连接的id
	ConnID uint32
	// 当前的连接状态
	isClosed bool
	// 告知当前连接已经退出的channel(由Reader告知Writer退出)
	ExitChan chan bool
	// 管理消息与其对应的处理函数
	MsgHandles ziface.IMsgHandle
	// 无缓冲的管道,用于读写Goroutine之间的消息通信.
	msgChan chan []byte
}

// NewConnection 初始化链接模块的方法
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandles ziface.IMsgHandle) *Connection {
	con := &Connection{
		Server:     server,
		Conn:       conn,
		ConnID:     connID,
		MsgHandles: msgHandles,
		isClosed:   false,
		msgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
	}
	// 将当前conn加入Connection Manager
	con.Server.GetConnMgr().Add(con)
	return con
}

// StartReader 收消息goroutine,专门接收客户端消息的模块
func (c *Connection) StartReader() {
	sugaredLogger.Infoln("[Reader Goroutine is running...]")
	defer sugaredLogger.Infof("[Reader Goroutine is exit,ConnID: %d, remote addr is %v ...]", c.ConnID, c.RemoteAddr())
	defer c.Stop()

	dp := NewDataPack()
	for {
		// 读取数据包消息头字节数据
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			if err != io.EOF {
				sugaredLogger.Errorf("Read msg head err: %v", err)
			}
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

		// 将消息发送到某个消息队列
		c.MsgHandles.SendMsgToTaskQueue(req)
	}
}

// StartWriter 写消息goroutine,专门给客户端发送消息的模块
func (c *Connection) StartWriter() {
	sugaredLogger.Infof("[Writer Goroutine is running...]")
	defer sugaredLogger.Infof("[Writer Goroutine is exit,ConnID: %d, remote addr is %v ...]", c.ConnID, c.RemoteAddr())
	// 不断阻塞的等待msgChan的消息,并响应客户端.
	for {
		select {
		// 有数据要写给客户端
		case msgData := <-c.msgChan:
			// 将数据发送给客户端
			if _, err := c.Conn.Write(msgData); err != nil {
				sugaredLogger.Errorf("Send data err: %v", err)
				return
			}
		case <-c.ExitChan:
			// 代表Reader goroutine 已退出.此时Writer goroutine也要退出
			return
		}
	}
}

func (c *Connection) Start() error {

	// 创建连接对象后,执行响应Hook函数.
	c.Server.CallOnCreateConnAfter(c)

	// 启动对当前连接的读数据的goroutine.
	go c.StartReader()
	// 启动对当前连接的写数据的goroutine.
	go c.StartWriter()

	return nil
}

func (c *Connection) Stop() error {
	// 如果当前连接已经关闭
	if c.isClosed {
		return nil
	}

	// 销毁连接之前,执行相应的Hook函数
	c.Server.CallOnDestroyConnBefore(c)

	err := c.Conn.Close()
	c.isClosed = true

	// 将当前连接从Connection Manager 中移除
	c.Server.GetConnMgr().Remove(c.GetConnID())

	// 告知Writer goroutine 关闭
	close(c.ExitChan)

	close(c.msgChan)
	return err
}

// SendMsg 提供一个SendMsg方法,将要发送给客户端的数据,先进行封包,再发送.
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection closed when send msg")
	}

	dp := NewDataPack()

	// 将data进行封包
	sendData, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		sugaredLogger.Errorf("Pack msg err: %v", err)
		return err
	}
	// 将数据发送给 Writer goroutine
	c.msgChan <- sendData
	return nil
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
