/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-25 15:30:59
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-25 16:06:00
 * @Description:
	使用Redigo 操作redis
*/

package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// go向redis写入数据
	// _, err = conn.Do("set", "s1", "hello")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("set suc")

	// data, err := conn.Do("get", "s1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("get suc. data:", data)

	// data, err := redis.String(conn.Do("get", "s1"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("get suc. data:", data)

	// 设置key过期时间
	// _, err = conn.Do("expire", "name", 10) //10秒过期
	// if err != nil {
	// 	fmt.Println("set expire error: ", err)
	// 	return
	// }

	// _, err = conn.Do("MSET", "name", "wd", "age", 22)
	// if err != nil {
	// 	fmt.Println("redis mset error:", err)
	// 	return
	// }
	// fmt.Println("mset suc.")

	// res, err := redis.Strings(conn.Do("MGET", "name", "age"))
	// if err != nil {
	// 	fmt.Println("redis get error:", err)
	// 	return
	// }
	// fmt.Println("mget suc. data: ", res)

	// _, err = conn.Do("HSET", "student", "name", "wd", "age", 22)
	// if err != nil {
	// 	fmt.Println("redis mset error:", err)
	// 	return
	// }
	// fmt.Println("hset suc.")

	// res, err := redis.String(conn.Do("HGET", "student", "age"))
	// res, err := redis.Int(conn.Do("HGET", "student", "age"))
	// res, err := redis.Int64(conn.Do("HGET", "student", "age"))
	// if err != nil {
	// 	fmt.Println("redis HGET error:", err)
	// 	return
	// }
	// fmt.Println("hget suc.data:", res)

	_, err = conn.Do("LPUSH", "list1", "ele1", "ele2", "ele3")
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.String(conn.Do("LPOP", "list1"))
	if err != nil {
		fmt.Println("redis POP error:", err)
	}
	fmt.Println("lpop suc.data:", res)
}
