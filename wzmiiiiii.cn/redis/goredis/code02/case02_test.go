/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-25 21:02:00
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-25 22:45:27
 * @Description:
 */
package code02

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
)

// zset(有序集合)示例
func Test3(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// key
	zsetkey := "language_rank"

	// value
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}

	// ZADD 指令
	intCmd := rct.ZAdd(ctx, zsetkey, languages...)
	err := intCmd.Err()
	if err != nil {
		fmt.Println("zadd failed. err: ", err)
		return
	}
	fmt.Println("zdd suc")

	// 把Golang的分数加10
	newScore, err := rct.ZIncrBy(ctx, zsetkey, 10, "Golang").Result()
	if err != nil {
		fmt.Println("zincrby failed. err: ", err)
		return
	}
	fmt.Printf("Golang's score is %.2f now.\n", newScore)

	// 取分数最高的3个
	ret := rct.ZRevRange(ctx, zsetkey, 0, 2).Val()
	for _, item := range ret {
		fmt.Println(item)
	}
	// 取分数最高的3个
	// NewZSliceCmd(ctx, "zrevrange", key, start, stop, "withscores")
	ret1, err := rct.ZRevRangeWithScores(ctx, zsetkey, 0, 2).Result()
	if err != nil {
		fmt.Println("zrevrange failed.err:", err)
		return
	}
	for _, z := range ret1 {
		fmt.Println(z.Score, z.Member)
	}

	// 取95~100分的
	ret2, err := rct.ZRangeByScoreWithScores(ctx, zsetkey, &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}).Result()

	if err != nil {
		fmt.Println("zrangebyscore failed. err:", err)
		return
	}
	for _, z := range ret2 {
		fmt.Println(z.Member, z.Score)
	}

}

// 扫描或遍历所有key
// 你可以使用KEYS prefix* 命令按前缀获取所有 key。
// vals, err := rdb.Keys(ctx, "prefix*").Result()

// 但是如果需要扫描数百万的 key ，那速度就会比较慢。这种场景下你可以使用Scan 命令来遍历所有符合要求的 key。
func Test4(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err := rct.Scan(ctx, cursor, "*", 0).Result()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, key := range keys {
			fmt.Println("key:", key)
		}
		if cursor == 0 {
			break
		}
	}
}

// Go-redis 允许将上面的代码简化为如下示例。
func Test5(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	iter := rct.Scan(ctx, 0, "s*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys:", iter.Val())
	}
	if err := iter.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

// 一个将所有匹配指定模式的 key 删除的示例。
func Test6(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	iter := rct.Scan(ctx, 0, "s*", 0).Iterator()

	for iter.Next(ctx) {
		err := rct.Del(ctx, iter.Val()).Err()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if err := iter.Err(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("del suc.")
}

// 此外，对于 Redis 中的 set、hash、zset 数据类型，go-redis 也支持类似的遍历方法。
// iter := rdb.SScan(ctx, "set-key", 0, "prefix:*", 0).Iterator()
// iter := rdb.HScan(ctx, "hash-key", 0, "prefix:*", 0).Iterator()
// iter := rdb.ZScan(ctx, "sorted-hash-key", 0, "prefix:*", 0).Iterator()
