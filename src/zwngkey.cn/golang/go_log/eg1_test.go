/*
 * @Author: zwngkey
 * @Date: 2022-05-13 02:44:37
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-20 18:25:10
 * @Description:
 */
package gopgklog

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

/*
	log包实现了简单的日志打印功能，支持日志输出到控制台或者日志文件。log包里核心的数据结构只有1个Logger，定义如下
		type Logger struct {
			mu     sync.Mutex // ensures atomic writes; protects the following fields
			prefix string     // prefix on each line to identify the logger (but see Lmsgprefix)
			flag   int        // properties
			out    io.Writer  // destination for output
			buf    []byte     // for accumulating text to write
		}

	Logger结构体里的字段，在使用上我们只需要关心prefix，flag和out这3个字段的含义：
		out：表示日志输出的地方。可以是标准输出os.Stdout，os.Stderr或者指定的本地文件

		flag：日志的属性设置。表示每行日志最开头打印的内容。取值如下：
			const (
				Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
				Ltime                         // the time in the local time zone: 01:23:23
				Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
				Llongfile                     // full file name and line number: /a/b/c/d.go:23
				Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
				LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
				Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
				LstdFlags     = Ldate | Ltime // initial values for the standard logger
			)

		prefix：每行日志最开头的日志前缀
			注意：如果flag开启了Lmsgprefix，那这个prefix前缀就不是放在每行日志的最开头了，而是在具体被打印的内容的前面。

		Logger结构体实现了若干指针接收者方法，包括设置日志属性、打印日志等。
*/
/*
	要使用log包打印日志，有2种方式，可以根据各自业务场景选择对应方法：
		方法1：使用log包里自带的std这个Logger指针。通常用于在控制台输出日志。
		方法2：自定义Logger。通常用于把日志输出到文件里。

	方法1和方法2相比，没有本质区别，只是使用场景上有一个偏好。
	当然方法1也可以实现输出日志到文件里，方法2也可以实现在控制台打印日志。

*/
/*
	方法1： log自带的标准错误输出

	下面的示例，使用了log包里自带的std标准输出，先通过SetFlags和SetPrefix这2个log包里的函数
		设置好std指向的Logger结构体对象里的flag和prefix属性，然后通过log包里定义的Println函数，把日志打印到控制台。
*/
func Test11(t *testing.T) {
	// 通过SetFlags设置Logger结构体里的flag属性
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	// 通过SetPrefix设置Logger结构体里的prefix属性
	log.SetPrefix("INFO:")

	// 调用辅助函数Println打印日志到标准错误输出
	log.Println("your message")
	log.Fatalln("123")
	log.Panicln(123)
}

/*
	方式1的使用流程如下：
		1.通过调用SetFlags，SetPrefix，SetOutput函数设置好日志属性。SetOutPut可以用于设置日志输出的地方，比如终端，文件等。
			如果省略这个步骤，会使用std创建时设置的默认属性。

			从源码中可以看出std是默认把日志输出到控制台，默认日志的prefix前缀为空串，默认flag属性是LstdFlags，
				也就是日志开头会打印日期和时间，比如：2009/01/23 01:23:23

		2.调用log包里的辅助函数Print[f|ln]，Fatal[f|ln]，Panic[f|ln]打印日志
			Fatal[f|ln]打印日志后会调用os.Exit(1)
			Panic[f|ln]打印日志后会调用panic
*/

/*
上面的例子是使用log包自带的std这个Logger指针把日志输出到控制台，我们也可以使用std把日志输出到指定文件，

	调用SetOutput设置日志输出的参数即可。如下代码示例：
*/
func Test12(t *testing.T) {
	filename := fmt.Sprintf("app_%s.log", time.Now().Format("2006-01-02-15-04-05"))

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("open file err: %s", err)
	}
	defer f.Close()
	log.SetOutput(f)
	// 调用SetOutput设置日志输出的地方
	// log.SetOutput(f)
	log.SetOutput(io.MultiWriter(os.Stdout, f))
	// 通过SetFlags设置Logger结构体里的flag属性
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	// 通过SetPrefix设置Logger结构体里的prefix属性
	log.SetPrefix("INFO:")
	// 调用辅助函数Println打印日志到指定文件
	log.Println("your message")
}

/*
方式2：自定义Logger

	方式1只建议打印到控制台的时候使用，对于打印到日志文件的场景，建议使用自定义Logger，参考如下代码：

	注意：New函数返回的是Logger指针，Logger结构体的方法都是指针接收者。

总结方式2的使用流程如下：

	1.通过log.New创建一个新的Logger指针，在New函数里指定好output, prefix和flag等日志属性
	2.调用log包里的辅助函数Print[f|ln]，Fatal[f|ln]，Panic[f|ln]打印日志

自定义Logger的方式，也可以实现打印日志到控制台，也可以实现同时打印日志到日志文件和控制台，

	只需要给New函数的第一个参数传递对应的io.Writer类型参数即可。

如果要打印到控制台，参数可以用os.Stdout或者os.Stderr

如果要同时打印到控制台和日志文件，参数可以用io.MultiWriter(os.Stdout, f)
*/
func Test13(t *testing.T) {
	// 打开文件
	fileName := fmt.Sprintf("app_%s.log", time.Now().Format("20060102"))
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	// 通过New方法自定义Logger，New的参数对应的是Logger结构体的output, prefix和flag字段
	logger := log.New(f, "[INFO] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
	// 调用Logger的方法Println打印日志到指定文件
	logger.Println("your message")
}

/*
	生产应用
		生产系统中打印日志就比上面的要复杂多了，需要考虑至少以下几个方面：
			日志路径设置：支持配置日志文件路径，将日志打印到指定路径的文件里。

			日志级别控制：支持Debug, Info, Warn, Error, Fatal等不同日志级别。

			日志切割：可以按照日期和日志大小进行自动切割。

			性能：在大量日志打印的时候不能对应用程序性能造成明显影响。

		Go生态中，目前比较流行的是Uber开发的zap，在GitHub上的开源地址：https://github.com/uber-go/zap
*/
/*
	LUTC属性：对于Logger结构体里的flag属性，如果开启了LUTC属性，那打印的日志里显示的时间就不是本地时间了，而是UTC标准时间。
		比如中国在东八区，中国时间减去8小时就是UTC时间。

	log打印的日志一定会换行。所以即使调用的是log包里的Print函数或方法，打印的日志也会换行。因此使用log包里的Print和Println没有区别了。

*/
