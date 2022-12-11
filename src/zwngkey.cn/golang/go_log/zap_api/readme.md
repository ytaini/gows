# go.uber.org/zap 日志库
https://blog.csdn.net/weixin_52690231/article/details/124879087

Zap是非常快的、结构化的，分日志级别的Go日志库。

- 它同时提供了结构化日志记录和printf风格的日志记录
- 它非常的快

> 配置Zap Logger

Zap提供了两种类型的日志记录器—`Sugared Logger`和`Logger`。

在性能很好但不是很关键的上下文中，使用`SugaredLogger`。它比其他结构化日志记录包快4-10倍，并且支持结构化和printf风格的日志记录。

在每一微秒和每一次内存分配都很重要的上下文中，使用`Logger`。它甚至比`SugaredLogger`更快，内存分配次数也更少，但它`只支持强类型的结构化日志记录`。

> Logger

- 通过调用`zap.NewProduction()`,`zap.NewDevelopment()`或者`zap.Example()`创建一个`Logger`。
- 上面的每一个函数都将创建一个logger。
  - 唯一的区别在于它将记录的信息不同。例如production logger默认记录调用函数信息、日期和时间等。
- 通过Logger调用Info/Error等。
- 默认情况下日志都会打印到应用程序的console界面。

```go
var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error fetching url..", zap.String("url", url), zap.Error(err))
	} else {
		logger.Info("success...", zap.String("statusCode", resp.Status), zap.String("url", url))
		_ = resp.Body.Close()
	}
}
```

在上面的代码中，我们首先创建了一个`Logger`，然后使用`Info/Error`等Logger方法记录消息。


日志记录器方法的语法是这样的：
```go
func (log *Logger) MethodXXX(msg string, fields ...Field) 
```
其中`MethodXXX`是一个可变参数函数，可以是`Info / Error/ Debug / Panic`等。每个方法都接受一个消息字符串和任意数量的zapcore.Field场参数。

每个`zapcore.Field`其实就是一组键值对参数。

执行上面的代码会得到如下输出结果：
```log
{"level":"error","ts":1670708238.44241,"caller":"main/main.go:24","msg":"Error fetching url..","url":"www.baidu.com","error":"Get \"www.baidu.com\": unsupported protocol scheme \"\"","stacktrace":"main.simpleHttpGet\n\t/Users/imzw/gows/src/zwngkey.cn/golang/go_log/zap_api/test1/main/main.go:24\nmain.main\n\t/Users/imzw/gows/src/zwngkey.cn/golang/go_log/zap_api/test1/main/main.go:13\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:250"}
{"level":"info","ts":1670708238.5097191,"caller":"main/main.go:26","msg":"success...","statusCode":"200 OK","url":"http://www.baidu.com"}
```

<br>


> Sugared Logger

使用Sugared Logger来实现相同的功能。

- 大部分的实现基本都相同。
- 唯一的区别是，通过调用`主logger的Sugar()`方法来获取一个`SugaredLogger`。
- 然后使用SugaredLogger以printf格式记录语句

下面是修改过后使用SugaredLogger代替Logger的代码：
```go

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}
func InitLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
```

执行上面的代码会得到如下输出：
```log
{"level":"error","ts":1670709293.98158,"caller":"test2/main.go:25","msg":"Error fetching URL www.baidu.com : Error = Get \"www.baidu.com\": unsupported protocol scheme \"\"","stacktrace":"main.simpleHttpGet\n\t/Users/imzw/gows/src/zwngkey.cn/golang/go_log/zap_api/test2/main.go:25\nmain.main\n\t/Users/imzw/gows/src/zwngkey.cn/golang/go_log/zap_api/test2/main.go:13\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:250"}
{"level":"info","ts":1670709294.0655339,"caller":"test2/main.go:27","msg":"Success! statusCode = 200 OK for URL http://www.baidu.com"}
```

<br>

> 这两个logger都打印输出JSON结构格式。

<br>

> 定制logger

将日志写入文件而不是终端

- 将使用zap.New(…)方法来手动传递所有配置，而不是使用像zap.NewProduction()这样的预置方法来创建logger。

```go
func New(core zapcore.Core, options ...Option) *Logger
```
`zapcore.Core`需要三个配置——`Encoder，WriteSyncer，LogLevel`。

1. `Encoder:编码器`(如何写入日志)。我们将使用开箱即用的NewJSONEncoder()，并使用预先设置的ProductionEncoderConfig()。
```go
 zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
```
2. `WriterSyncer` ：指定日志将写到哪里去。我们使用zapcore.AddSync()函数并且将打开的文件句柄传进去。
```go
file, _ := os.Create("./test.log")
writeSyncer := zapcore.AddSync(file)
```
3. `Log Level`：哪种级别的日志将被写入。

将修改上述部分中的Logger代码，并重写InitLogger()方法。其余方法—main() /SimpleHttpGet()保持不变。

```go
func InitLogger() {
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}
```

当使用这些修改过的logger配置调用上述部分的main()函数时，以下输出将打印在文件——test.log中。
```log
{"level":"debug","ts":1670710197.821301,"msg":"Trying to hit GET request for www.baidu.com"}
{"level":"error","ts":1670710197.8215501,"msg":"Error fetching URL www.baidu.com : Error = Get \"www.baidu.com\": unsupported protocol scheme \"\""}
{"level":"debug","ts":1670710197.821561,"msg":"Trying to hit GET request for http://www.baidu.com"}
{"level":"info","ts":1670710197.8819,"msg":"Success! statusCode = 200 OK for URL http://www.baidu.com"}
```
<br>

> 将JSON Encoder更改为普通的Log Encoder

现在，我们希望将编码器从JSON Encoder更改为普通Encoder。为此，我们需要将NewJSONEncoder()更改为NewConsoleEncoder()。
```go
return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
```

当使用这些修改过的logger配置调用上述部分的main()函数时，以下输出将打印在文件——test.log中。
```log
1.670710213494401e+09	debug	Trying to hit GET request for www.baidu.com
1.67071021349453e+09	error	Error fetching URL www.baidu.com : Error = Get "www.baidu.com": unsupported protocol scheme ""
1.670710213494535e+09	debug	Trying to hit GET request for http://www.baidu.com
1.670710213533246e+09	info	Success! statusCode = 200 OK for URL http://www.baidu.com
```


> 更改时间编码并添加调用者详细信息

鉴于我们对配置所做的更改，有下面两个问题：
- 时间是以非人类可读的方式展示，例如1.572161051846623e+09
- 调用方函数的详细信息没有显示在日志中

我们要做的第一件事是覆盖默认的ProductionConfig()，并进行以下更改:
- 修改时间编码器
- 在日志文件中使用大写字母记录日志级别
```go
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
```
接下来，我们将修改zap logger代码，添加将调用函数信息记录到日志中的功能。为此，我们将在zap.New(..)函数中添加一个Option。
```go
logger := zap.New(core, zap.AddCaller())
```

当使用这些修改过的logger配置调用上述部分的main()函数时，以下输出将打印在文件——test.log中。
```log
2022-12-11T06:17:03.241+0800	DEBUG	test3/main.go:42	Trying to hit GET request for www.baidu.com
2022-12-11T06:17:03.241+0800	ERROR	test3/main.go:45	Error fetching URL www.baidu.com : Error = Get "www.baidu.com": unsupported protocol scheme ""
2022-12-11T06:17:03.241+0800	DEBUG	test3/main.go:42	Trying to hit GET request for http://www.baidu.com
2022-12-11T06:17:03.295+0800	INFO	test3/main.go:47	Success! statusCode = 200 OK for URL http://www.baidu.com
```

<br>

> AddCallerSkip

当我们不是直接使用初始化好的logger实例记录日志，而是将其包装成一个函数等，此时日录日志的函数调用链会增加，想要获得准确的调用信息就需要通过AddCallerSkip函数来跳过。

```go
logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
```

> 将日志输出到多个位置

我们可以将日志同时输出到文件和终端。

```go
func getLogWriter() zapcore.WriteSyncer {
    file, _ := os.Create("./test.log")
    //return zapcore.AddSync(file)
    // 将日志输出到多个位置
    //return zapcore.NewMultiWriteSyncer(zapcore.AddSync(file), zapcore.AddSync(os.Stderr))
    // 利用io.MultiWriter 将日志输出到多个位置
    mw := io.MultiWriter(file, os.Stderr)
    return zapcore.AddSync(mw)
}
```

> 将err日志单独输出到文件

有时候我们除了将全量日志输出到xx.log文件中之外，还希望将ERROR级别的日志单独输出到一个名为xx.err.log的日志文件中。我们可以通过以下方式实现。

```go
func InitLogger() {
	encoder := getEncoder()
	// test.log记录全量日志
	logF, _ := os.Create("./test.log")
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
	// test.err.log记录ERROR级别的日志
	errF, _ := os.Create("./test.err.log")
	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
	// 使用NewTee将c1和c2合并到core
	core := zapcore.NewTee(c1, c2)
	logger = zap.New(core, zap.AddCaller())
}
```

<br>

# 使用Lumberjack进行日志切割归档

zap日志程序中唯一缺少的就是日志切割归档功能。

> Zap本身不支持切割归档日志文件

官方的说法是为了添加日志切割归档功能，我们将使用第三方库Lumberjack来实现。

目前只支持按文件大小切割，原因是按时间切割效率低且不能保证日志数据不被破坏。详情戳https://github.com/natefinch/lumberjack/issues/54。

想按日期切割可以使用github.com/lestrrat-go/file-rotatelogs这个库，虽然目前不维护了，但也够用了。

```go
// 使用file-rotatelogs按天切割日志

import rotatelogs "github.com/lestrrat-go/file-rotatelogs"

l, _ := rotatelogs.New(
	filename+".%Y%m%d%H%M",
	rotatelogs.WithMaxAge(30*24*time.Hour),    // 最长保存30天
	rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
)
zapcore.AddSync(l)
```

<br>

> 安装

执行下面的命令安装 Lumberjack v2 版本。

```shell
go get gopkg.in/natefinch/lumberjack.v2
```

<br>

> zap logger中加入Lumberjack


要在zap中加入Lumberjack支持，我们需要修改WriteSyncer代码。我们将按照下面的代码修改getLogWriter()函数：

```go
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
```

Lumberjack Logger采用以下属性作为输入:
- Filename: 日志文件的位置
- MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
- MaxBackups：保留旧文件的最大个数
- MaxAges：保留旧文件的最大天数
- Compress：是否压缩/归档旧文件


