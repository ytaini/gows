/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-25 16:26:02
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-25 16:45:35
 * @Description:
 */

package case04

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Test(t *testing.T) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	_, err = c.Do("SET", "name", "duncanwang", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get name: %v \n", username)
	}
}

func Test2(t *testing.T) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	_, err = c.Do("MSET", "name", "duncanwang", "sex", "male")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	is_key_exit, err := redis.Bool(c.Do("EXISTS", "name"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("Is name key exists ? : %v \n", is_key_exit)
	}

	reply, err := redis.Values(c.Do("MGET", "name", "sex"))
	if err != nil {
		fmt.Printf("patch get name & sex error \n。")
	} else {
		var name string
		var sex string
		_, err := redis.Scan(reply, &name, &sex)
		if err != nil {
			fmt.Printf("Scan error \n。")
		} else {
			fmt.Printf("The name is %v, sex is %v \n", name, sex)
		}
	}
}

// 读写json到redis转换
func Test3(t *testing.T) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer conn.Close()
	imap := map[string]string{"name": "zhangsan", "age": "10", "tel": "10086"}
	data, _ := json.Marshal(imap)
	_, err = conn.Do("SETNX", "profile", data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("setnx suc")
}

func Test4(t *testing.T) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer conn.Close()
	var imap map[string]string
	data, _ := redis.Bytes(conn.Do("GET", "profile"))

	fmt.Println("get suc.")
	err = json.Unmarshal(data, &imap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("unmarshal suc.")
	fmt.Println("data:", imap)
}
