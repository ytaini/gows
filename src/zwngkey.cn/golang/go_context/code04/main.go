/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-27 12:47:12
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-27 15:07:59
 * @Description:
 */

package main

import (
	"context"
	"fmt"
	"time"
)

// 取消控制

// 日常业务开发中我们往往为了完成一个复杂的需求会开多个gouroutine去做一些事情，
// 这就导致我们会在一次请求中开了多个goroutine确无法控制他们，
// 这时我们就可以使用withCancel来衍生一个context传递到不同的goroutine中，
// 当我想让这些goroutine停止运行，就可以调用cancel来进行取消。
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		Speak(ctx)
		done <- struct{}{}
	}()
	time.Sleep(5 * time.Second)
	cancel()
	<-done
}

// 我们使用withCancel创建一个基于Background的ctx，
// 然后启动一个讲话程序，每隔1s说一下话，main函数在5s后执行cancel，那么speak检测到取消信号就会退出。
func Speak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("我要闭嘴了")
			return
		default:
			fmt.Println("balabalabalabala")
		}
	}
}
