package znet

import (
	"errors"
	"fmt"
	"sync"

	"wzmiiiiii.cn/zinx/ziface"
)

// 连接管理模块

type ConnManager struct {
	// 管理连接的集合
	conns map[uint32]ziface.IConnection
	mu    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{conns: make(map[uint32]ziface.IConnection)}
}

func (c *ConnManager) Add(connection ziface.IConnection) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.conns[connection.GetConnID()] = connection
	sugaredLogger.Infof("[ID:%d] Connection add to connection manager successfully...", connection.GetConnID())
}

func (c *ConnManager) Remove(connId uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.conns, connId)
	sugaredLogger.Infof("[ID:%d] Connection remove to connection manager successfully...", connId)
}

func (c *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	conn, ok := c.conns[connId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("connection[ID:%d] NOT FOUND...", connId))
	}
	return conn, nil
}

func (c *ConnManager) GetSize() int {
	return len(c.conns)
}

func (c *ConnManager) ClearAll() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, conn := range c.conns {
		_ = conn.Stop()
	}
}
