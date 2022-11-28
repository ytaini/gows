/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-27 11:50:23
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-27 12:14:43
 * @Description:
 */
package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

var wg sync.WaitGroup

type str string

const KEY str = "trace_id"

// WithValue 携带数据.

// 我们日常在业务开发中都希望能有一个trace_id能串联所有的日志，这就需要我们打印日志时能够获取到这个trace_id，
// 在python中我们可以用gevent.local来传递，
// 在java中我们可以用ThreadLocal来传递，
// 在Go语言中我们就可以使用Context来传递，通过使用WithValue来创建一个携带trace_id的context，
// 然后不断透传下去，打印日志时输出即可，来看使用例子：
func main() {
	ProcessEnter(NewContextWithTraceID())
	wg.Wait()
}

func NewRequestID() string {
	return strings.Replace(uuid.NewString(), "-", "", -1)
}

func NewContextWithTraceID() context.Context {
	return context.WithValue(context.Background(), KEY, NewRequestID())
}

func ProcessEnter(ctx context.Context) {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			defer wg.Done()
			printLog(ctx, "go"+strconv.Itoa(i))
		}()
	}
}

func printLog(ctx context.Context, msg string) {
	fmt.Printf("%s|info|trace_id=%s|%s\n", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), msg)
}

func GetContextValue(ctx context.Context, k str) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

// 我们基于context.Background创建一个携带trace_id的ctx，然后通过context树一起传递，
// 从中派生的任何context都可以获取此值，我们最后打印日志的时候就可以从ctx中取值输出到日志中。
// 目前一些RPC框架都是支持了Context，所以trace_id的向下传递就更方便了。
