## Gee 的错误处理机制

对一个 Web 框架而言，错误处理机制是非常必要的。可能是框架本身没有完备的测试，导致在某些情况下出现空指针异常等情况。也有可能用户不正确的参数，触发了某些异常，例如数组越界，空指针等。如果因为这些原因导致系统宕机，必然是不可接受的。

如果框架并没有加入异常处理机制，如果代码中存在会触发 panic 的 BUG，很容易宕掉。

例如下面的代码:
```go
func main() {
	r := gee.New()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
```

在上面的代码中，我们为 gee 注册了路由 /panic，而这个路由的处理函数内部存在数组越界 names[100]，如果访问 localhost:9999/panic，Web 服务就会宕掉。

今天，我们将在 gee 中添加一个非常简单的错误处理机制，即在此类错误发生时，向用户返回 Internal Server Error，并且在日志中打印必要的错误信息，方便进行错误定位。

我们之前实现了中间件机制，错误处理也可以作为一个中间件，增强 gee 框架的能力。

新增文件 `gee/recovery.go`，在这个文件中实现中间件 `Recovery`。
```go
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
```

Recovery 的实现非常简单，使用 defer 挂载上错误恢复的函数，在这个函数中调用 *recover()*，捕获 panic，并且将堆栈信息打印在日志中，向用户返回 Internal Server Error。

trace() 函数，这个函数是用来获取触发 panic 的堆栈信息，完整代码如下：
```go
// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
```

在 trace() 中，调用了 runtime.Callers(3, pcs[:])，Callers 用来返回调用栈的程序计数器, 第 0 个 Caller 是 Callers 本身，第 1 个是上一层 trace，第 2 个是再上一层的 defer func。因此，为了日志简洁一点，我们跳过了前 3 个 Caller。

接下来，通过 runtime.FuncForPC(pc) 获取对应的函数，在通过 fn.FileLine(pc) 获取到调用该函数的文件名和行号，打印在日志中。

至此，gee 框架的错误处理机制就完成了。

## 使用Demo
```go
func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
```

接下来进行测试，先访问主页，访问一个有BUG的 /panic，服务正常返回。接下来我们再一次成功访问了主页，说明服务完全运转正常。

```cmd
$ curl "http://localhost:9999"
Hello Geektutu
$ curl "http://localhost:9999/panic"
{"message":"Internal Server Error"}
$ curl "http://localhost:9999"
Hello Geektutu
```

我们可以在后台日志中看到如下内容，引发错误的原因和堆栈信息都被打印了出来，通过日志，我们可以很容易地知道，在day7-panic-recover/main.go:47 的地方出现了 index out of range 错误。

```log
2022/11/29 21:48:00 Route  GET - /
2022/11/29 21:48:00 Route  GET - /panic
2022/11/29 21:48:11 [200] / in 9.791µs
2022/11/29 21:48:16 runtime error: index out of range [100] with length 1
TraceBack:
        /usr/local/go/src/runtime/panic.go:885
        /usr/local/go/src/runtime/panic.go:113
        /Users/imzw/gows/src/zwngkey.cn/golang/go_http/code01/case08/main.go:25
        /Users/imzw/gows/src/zwngkey.cn/golang/go_http/code01/case08/gee/context.go:44
        /Users/imzw/gows/src/zwngkey.cn/golang/go_http/code01/case08/gee/recovery.go:27
        /Users/imzw/gows/src/zwngkey.cn/golang/go_http/code01/case08/gee/context.go:44
        /Users/imzw/gows/src/zwngkey.cn/golang/go_http/code01/case08/gee/logger.go:18
        /Users/imzw/gows/src/zwngkey.cn/golang/go_http/code01/case08/gee/context.go:44
        /Users/imzw/gows/src/zwngkey.cn/golang/go_http/code01/case08/gee/router.go:106
        /Users/imzw/gows/src/zwngkey.cn/golang/go_http/code01/case08/gee/gee.go:46
        /usr/local/go/src/net/http/server.go:2948
        /usr/local/go/src/net/http/server.go:1992
        /usr/local/go/src/runtime/asm_arm64.s:1173

2022/11/29 21:48:16 [500] /panic in 504.833µs
2022/11/29 21:48:31 [200] / in 32.75µs
```