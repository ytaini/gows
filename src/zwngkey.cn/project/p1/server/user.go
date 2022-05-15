/*
 * @Author: zwngkey
 * @Date: 2022-05-15 21:52:56
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-16 03:36:08
 * @Description:
 */
package main

import (
	"log"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	Ch   chan string
	Conn net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		Ch:     make(chan string),
		Conn:   conn,
		server: server,
	}

	go user.ListenMsg()

	return user
}

//监听当前User channel方法,channel 一有消息,就发送
func (u *User) ListenMsg() {
	for msg := range u.Ch {
		_, err := u.Conn.Write([]byte(msg + "\n"))
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func (u *User) Online() {
	//将上线用户加入到OnlineMap中
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()

	//广播当前用户上线信息
	u.server.Broadcast(u, "已上线")
}

func (u *User) Offline() {
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()

	u.server.Broadcast(u, "下线了")
}

func (u *User) DealMsg(msg string) {
	if msg == "who" {
		u.server.mapLock.RLock()
		for _, userInfo := range u.server.OnlineMap {
			userMsg := "[" + userInfo.Addr + "]" + userInfo.Name + ": " + "在线..."
			u.SendMsg(userMsg)
		}
		u.server.mapLock.RUnlock()
	} else if len(msg) > 7 && msg[:7] == "rename:" {
		ss := strings.Split(msg, ":")[1]
		_, ok := u.server.OnlineMap[ss]
		if ok {
			u.SendMsg("提示:该用户名已存在...")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.server.OnlineMap[ss] = u
			u.server.mapLock.Unlock()
			u.Name = ss
			u.SendMsg("提示:修改完成")
		}
	} else if len(msg) > 3 && msg[:3] == "to-" {
		ss := msg[3:]
		if !strings.Contains(ss, ":") {
			u.SendMsg(`提示:格式错误!!,请使用"to-username:msg"格式.`)
		} else {
			msgCtx := strings.Split(ss, ":")
			user, ok := u.server.OnlineMap[msgCtx[0]]
			if !ok {
				u.SendMsg("提示:该用户不在线...")
			} else {
				sendMsg := "[" + u.Addr + "]" + u.Name + ": " + msgCtx[1]
				user.Ch <- sendMsg
			}
		}
	} else if len(msg) > 4 && msg[:4] == "all:" {
		u.server.Broadcast(u, msg[4:])
	} else {
		u.SendMsg("提示:输入错误!!")
	}
}

func (u *User) SendMsg(msg string) {
	u.Conn.Write([]byte(msg + "\n"))
}
