/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 01:38:51
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 02:23:43
 */
package gee

import "net/http"

// 将路由(router)独立出去，方便之后增强。
// 设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。

type HandlerFunc func(*Context)

// 将router相关的代码独立后，gee.go简单了不少。
// 最重要的还是通过实现了 ServeHTTP 接口，接管了所有的 HTTP 请求。
// 相比之前的代码，这个方法也有细微的调整，在调用 router.handle 之前，构造了一个 Context 对象。
// 这个对象目前还非常简单，仅仅是包装了原来的两个参数
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
