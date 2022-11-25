/*
 * @Author: zwngkey
 * @Date: 2022-05-16 03:36:32
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-16 05:38:59
 * @Description:
 */
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	IP   string
	Port int
	Name string
	conn net.Conn
	flag int
}

var ip string
var port int

func init() {
	flag.StringVar(&ip, "i", "127.0.0.1", "ip address")
	flag.IntVar(&port, "p", 8080, "port")
}

func (c *Client) menu() {
	fmt.Println(`
	1.公聊模式
	2.私聊模式
	3.修改用户名
	4.查看用户名
	0.退出
	`)
	var flag int
	for {
		_, err := fmt.Scanln(&flag)
		if err != nil {
			fmt.Println("非法,重新输入")
			continue
		}
		if flag >= 0 && flag <= 4 {
			c.flag = flag
			break
		} else {
			fmt.Println("非法,重新输入")
		}
	}
}
func (c *Client) updateName() bool {
	fmt.Print("请输入用户名:")
	fmt.Scanln(&c.Name)
	sendMsg := "rename:" + c.Name + "\n"
	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		log.Println("conn write err:", err)
		return false
	}
	return true
}
func (c *Client) PublicChat() {
	var msg string
	for strings.TrimSpace(msg) != "exit" {
		time.Sleep(time.Microsecond * 500)
		msg = ""
		fmt.Println("请输入聊天内容:")
		fmt.Scan(&msg)
		if len(strings.TrimSpace(msg)) != 0 && strings.TrimSpace(msg) != "exit" {
			sendMsg := "all:" + msg + "\n"
			_, err := c.conn.Write([]byte(sendMsg))
			if err != nil {
				log.Println("发送信息失败")
				break
			}
		}
	}
}

func (c *Client) PrivateChat() {
	var username string
	for {
		fmt.Println("请输入用户名[输入exit退出][输入show查看当前在线用户]")
		fmt.Scanln(&username)
		if username == "exit" {
			return
		} else if username == "show" {
			time.Sleep(time.Microsecond * 500)
			c.showAllUser()
		} else {
			var msg string
			for {
				time.Sleep(time.Microsecond * 100)
				fmt.Println("输入消息内容,输入exit退出:")
				fmt.Scan(&msg)
				msg = strings.TrimSpace(msg)
				if msg == "exit" {
					return
				}
				if len(msg) != 0 {
					sendMsg := "to-" + username + ":" + msg + "\n"
					_, err := c.conn.Write([]byte(sendMsg))
					if err != nil {
						log.Println("发送信息失败")
						return
					}
				}

			}
		}
	}

}

func (c *Client) showAllUser() {
	sendMsg := "who\n"
	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		log.Println(err)
	}
}

func (c *Client) Run() {
	log.Println("服务器连接成功...")
	for c.flag != 0 {
		time.Sleep(time.Microsecond * 500)
		c.menu()
		switch c.flag {
		case 1:
			c.PublicChat()
		case 2:
			c.PrivateChat()
		case 3:
			c.updateName()
		case 4:
			fmt.Println("用户名:", c.Name)
		}
	}
}
func (c *Client) DealResp() {
	//这里copy函数会一直阻塞监听c.conn连接.
	io.Copy(os.Stdout, c.conn)
}

func NewClint(ip string, port int) *Client {
	client := &Client{
		IP:   ip,
		Port: port,
		flag: 999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Panicln("连接服务器失败:", err)
	}
	client.conn = conn
	client.Name = conn.LocalAddr().String()
	return client
}

func main() {
	flag.Parse()

	client := NewClint(ip, port)

	go client.DealResp()

	client.Run()
}
