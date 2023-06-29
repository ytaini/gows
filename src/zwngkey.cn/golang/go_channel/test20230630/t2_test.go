/*
 * @Author: wzmiiiiii
 * @Date: 2023-06-30 00:42:53
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2023-06-30 00:54:21
 */
package test20230630

import (
	"net/http"
	"testing"
	"time"
)

func timeoutMsg() string {
	return "timeout"
}

// 接口超时控制: 通过channel的阻塞机制.
func home(w http.ResponseWriter, req *http.Request) {
	var resp string
	// 控制readDB操作的操作时间:
	// 创建done channel 用于协程间同步
	done := make(chan struct{}, 1)
	go func() {
		resp = readDB()
		done <- struct{}{}
	}()
	// 通过Select多路复用来进行超时控制.
	select {
	case <-done:
	case <-time.After(300 * time.Millisecond):
		resp = timeoutMsg()
	}
	w.Write([]byte(resp))
}

func Test02(t *testing.T) {
	http.HandleFunc("/", home)
	http.ListenAndServe(":5678", nil)
}
