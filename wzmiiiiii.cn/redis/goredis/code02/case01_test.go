/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-25 17:40:30
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-25 21:00:40
 * @Description:	goredis基本使用
 */
package code02

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
)

// 基本使用1
func Test(t *testing.T) {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := conn.Get(ctx, "key").Result()
	fmt.Println(val)
	fmt.Println(err)              // redis: nil
	fmt.Println(err == redis.Nil) // true

	statusCmd := conn.Set(ctx, "key", "hello", time.Second)
	fmt.Println(statusCmd.Err())      //<nil>
	fmt.Println(statusCmd.Name())     //set
	fmt.Println(statusCmd.FullName()) //set
	fmt.Println(statusCmd.Val())      //OK
	fmt.Println(statusCmd.Result())   //OK <nil>
	fmt.Println(statusCmd.Args()...)  //set key hello ex 1

	strCmd := conn.Get(ctx, "key")
	fmt.Println(strCmd.Val()) //hello
	fmt.Println(strCmd.Err()) //<nil>
}

// go-redis 还提供了一个执行任意命令或自定义命令的 Do 方法，特别是一些 go-redis 库暂时不支持的命令都可以使用该方法执行。
func Test1(t *testing.T) {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// cmd := conn.Do(ctx, "set", "language", "golang")
	// if cmd.Err() == redis.Nil {
	// 	fmt.Println(cmd.Err())
	// 	return
	// }
	// fmt.Println("set suc.")

	val, err := conn.Do(ctx, "get", "asd").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println("key不存在")
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("val: ", val)
}

// redis.Nil
// go-redis 库提供了一个 redis.Nil 错误来表示 Key 不存在的错误。因此在使用 go-redis 时需要注意对返回错误的判断。
// 在某些场景下我们应该区别处理 redis.Nil 和其他不为 nil 的错误。
