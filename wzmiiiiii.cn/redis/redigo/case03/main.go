/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-25 16:18:09
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-25 16:23:22
 * @Description:
 */
package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 新建一个连接池
	pool := &redis.Pool{
		MaxIdle:     10,  //最初的连接数量
		MaxActive:   0,   //连接池最大连接数量,（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", ":6379")
		},
	}

	conn := pool.Get()
	defer conn.Close()

	// 0. ping正常返回pong， 异常res is nil, err not nil
	res, err := redis.String(conn.Do("ping"))
	if err != nil {
		fmt.Printf("ping err=%v\n", err.Error())
	}
	fmt.Printf("ping res=%v\n", res)
}
