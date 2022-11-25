/*
  - @Author: wzmiiiiii
  - @Date: 2022-11-25 17:16:00

* @LastEditors: wzmiiiiii
* @LastEditTime: 2022-11-25 23:31:57
  - @Description:
    连接 redis
*/
package code01

import (
	"context"
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v9"
)

// 普通连接模式

// 使用 redis.NewClient 函数连接 Redis 服务器
func Test(t *testing.T) {
	rct := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rct.Close()
	ctx := context.Background()
	status := rct.Ping(ctx)
	fmt.Println(status)
}

// 使用 redis.ParseURL 函数从表示数据源的字符串中解析得到 Redis 服务器的配置信息。
func Test1(t *testing.T) {
	// Tcp connection:
	// 	redis://<user>:<password>@<host>:<port>/<db_number>
	opt, err := redis.ParseURL(`redis://@localhost:6379/0`)
	if err != nil {
		fmt.Println(err)
		return
	}
	rct := redis.NewClient(opt)
	defer rct.Close()
	ctx := context.Background()
	status := rct.Ping(ctx)
	fmt.Println(status)
}

// TLS连接模式
// 如果使用的是 TLS 连接方式，则需要使用 tls.Config 配置。
func Test3(t *testing.T) {
	redis.NewClient(&redis.Options{
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			// Certificates: []tls.Certificate{cert},
			// ServerName: "your.domain.com",
		},
	})
}

// Redis Sentinel模式
// 使用下面的命令连接到由 Redis Sentinel 管理的 Redis 服务器。
func Test4(t *testing.T) {
	redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master-name",
		SentinelAddrs: []string{":9126", ":9127", ":9128"},
	})
}

// Redis Cluster模式
// 使用下面的命令连接到 Redis Cluster，go-redis 支持按延迟或随机路由命令。
func Test5(t *testing.T) {
	redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

		// 若要根据延迟或随机路由命令，请启用以下命令之一
		// RouteByLatency: true,
		// RouteRandomly: true,
	})
}
