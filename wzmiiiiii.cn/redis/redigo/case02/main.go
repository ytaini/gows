/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-25 16:06:08
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-25 16:10:25
 * @Description:
 */
package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 管道操作可以理解为并发操作，并通过Send()，Flush()，Receive()三个方法实现。
// 客户端可以使用send()方法一次性向服务器发送一个或多个命令，命令发送完毕时，
// 使用flush()方法将缓冲区的命令输入一次性发送到服务器，
// 客户端再使用Receive()方法依次按照先进先出的顺序读取所有命令操作结果。

// Send：发送命令至缓冲区
// Flush：清空缓冲区，将命令一次性发送至服务器
// Recevie：依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。
func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Send("hset", "stu", "name", "wd", "age", "22")
	conn.Send("hset", "stu", "score", "100")
	conn.Send("hget", "stu", "name")
	conn.Flush()

	res1, _ := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1) //2
	res2, _ := conn.Receive()
	fmt.Printf("Receive res2:%v\n", res2) //1
	res3, _ := conn.Receive()
	fmt.Printf("Receive res3:%s\n", res3) //22
}
