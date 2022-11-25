/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-25 22:14:57
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-25 23:11:38
 * @Description:
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

// Redis Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能。
// 区别于一个接一个地执行100个命令，你可以将这些命令放入 pipeline 中，然后使用1次读写操作像执行单个命令一样执行它们。
// 这样做的好处是节省了执行命令的网络往返时间（RTT）。

// 在下面的示例代码中演示了使用 pipeline 通过一个 write + read 操作来执行多个命令。
func Test7(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	pipe := rct.Pipeline()

	incr := pipe.Incr(ctx, "counter")
	pipe.Expire(ctx, "counter", time.Minute)
	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 在执行pipe.Exec之后才能获取到结果
	fmt.Println(incr.Val())
}

// 上面的代码相当于将以下两个命令一次发给 Redis Server 端执行，与不使用 Pipeline 相比能减少一次RTT。
// INCR counter
// EXPIRE counter 60

// 你也可以使用Pipelined 方法，它会在函数退出时调用 Exec。
func Test8(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var incr *redis.IntCmd
	cmds, err := rct.Pipelined(ctx, func(p redis.Pipeliner) error {
		incr = p.Incr(ctx, "counter")
		p.Expire(ctx, "counter", time.Minute)
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cmd := range cmds {
		fmt.Println(cmd)
		// incr counter: 1
		// expire counter 60: true
	}

	// 在执行pipe.Exec之后才能获取到结果
	fmt.Println(incr.Val())
}

// 我们可以遍历 pipeline 命令的返回值依次获取每个命令的结果。
// 下方的示例代码中使用pipiline一次执行了100个 Get 命令，在pipeline 执行后遍历取出100个命令的执行结果。
func Test9(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cmds, err := rct.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			pipe.Get(ctx, fmt.Sprintf("key%d", i))
		}
		return nil
	})
	if err != nil && !errors.Is(err, redis.Nil) {
		fmt.Println(err)
		return
	}

	for _, cmd := range cmds {
		// fmt.Println(cmd.(*redis.StringCmd).Val())
		fmt.Println(cmd)
	}
}

// 在那些我们需要一次性执行多个命令的场景下，就可以考虑使用 pipeline 来优化。
