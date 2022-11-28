/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 00:22:06
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 00:34:05
 * @Description:
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

// http.ListenAndServe()的第二个参数是一个http.Handler接口,
// 需要实现方法 ServeHTTP ，也就是说，只要传入任何实现了 ServerHTTP 接口的实例，
// 所有的HTTP请求，就都会交给了该实例处理了

// 我们定义了一个空的结构体Engine，实现了方法ServeHTTP。
type Engine struct{}

// 这个方法有2个参数，第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，
// 比如请求地址、Header和Body等信息；
// 第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应。
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "hello world")
	case "/hello":
		fmt.Fprintln(w, "hello go web")
	default:
		http.NotFound(w, r)
	}
}

// 在 main 函数中，我们给 ListenAndServe 方法的第二个参数传入了刚才创建的engine实例。
// 至此，我们走出了实现Web框架的第一步，即，将所有的HTTP请求转向了我们自己的处理逻辑。
// 还记得吗，在实现Engine之前，我们调用 http.HandleFunc 实现了路由和Handler的映射，
// 也就是只能针对具体的路由写处理逻辑。比如/hello。但是在实现Engine之后，我们拦截了所有的HTTP请求，
// 拥有了统一的控制入口。在这里我们可以自由定义路由映射的规则，也可以统一添加一些处理逻辑，
// 例如日志、异常处理等。
func main() {
	engine := new(Engine)
	log.Fatalln(http.ListenAndServe(":9090", engine))
}
