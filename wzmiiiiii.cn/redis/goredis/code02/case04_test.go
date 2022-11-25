/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-25 23:12:26
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-26 00:39:53
 * @Description:
 */
package code02

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
)

// 事务

// Redis 是单线程执行命令的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行，
// 例如在它们之间交替执行。但是，Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。

// 在这种场景我们需要使用 TxPipeline 或 TxPipelined 方法将 pipeline 命令使用 MULTI 和EXEC包裹起来。
func Test10(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	pipe := rct.TxPipeline()
	incr := pipe.Incr(ctx, "tx_pipeline_counter")
	pipe.Expire(ctx, "tx_pipeline_counter", time.Minute)
	_, err := pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)

	// 上面代码相当于在一个RTT下执行了下面的redis命令
	// MULTI
	// INCR tx_pipeline_counter
	// EXPIRE tx_pipeline_counter 60
	// EXEC
	var incr2 *redis.IntCmd
	_, err = rct.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		incr2 = pipe.Incr(ctx, "tx_pipeline_counter")
		pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
		return nil
	})
	fmt.Println(incr2.Val(), err)
}

// Watch

// 我们通常搭配 WATCH命令来执行事务操作。
// 从使用WATCH命令监视某个 key 开始，直到执行EXEC命令的这段时间里，
// 如果有其他用户抢先对被监视的 key 进行了替换、更新、删除等操作，
// 那么当用户尝试执行EXEC的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。

// Watch方法接收一个函数和一个或多个key作为参数
// Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error

// 下面的代码片段演示了 Watch 方法搭配 TxPipelined 的使用示例。
func Test11(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	key := "age"
	err := rct.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 假设操作耗时5秒
		// 5秒内我们通过其他的客户端修改key，当前事务就会失败
		time.Sleep(5 * time.Second)
		_, err = tx.TxPipelined(ctx, func(p redis.Pipeliner) error {
			return p.Set(ctx, key, n+1, 0).Err()
		})
		return err
	}, key)

	if err != nil {
		fmt.Println(err)
		return
	}
	// 将上面的函数执行并打印其返回值，如果我们在程序运行后的5秒内修改了被 watch 的 key 的值，那么该事务操作失败，返回redis: transaction failed错误。
}

// go-redis 官方文档中使用 GET 、SET和WATCH命令实现一个 INCR 命令的完整示例。
func Test12(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Redis transactions use optimistic locking.
	const maxRetries = 100

	// Increment transactionally increments the key using GET and SET commands.
	increment := func(key string) error {
		// Transactional function.
		txf := func(tx *redis.Tx) error {
			// 获得当前值或零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}
			// 实际操作（乐观锁定中的本地操作）
			n++

			// 仅在监视的Key保持不变的情况下运行
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				// pipe 处理错误情况
				pipe.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		// Retry if the key has been changed.
		for i := 0; i < maxRetries; i++ {
			err := rdb.Watch(ctx, txf, key)
			if err == nil {
				//  success
				return nil
			}
			if err == redis.TxFailedErr {
				// Optimistic lock lost. Retry.
				continue
			}
			return err
		}
		return errors.New("increment reached maximum number of retries")
	}

	var wg sync.WaitGroup
	wg.Add(maxRetries)
	for i := 0; i < maxRetries; i++ {
		go func() {
			defer wg.Done()
			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get(ctx, "counter3").Int()
	fmt.Println(n, err)
	// 在这个示例中使用了 redis.TxFailedErr 来检查事务是否失败。
}
