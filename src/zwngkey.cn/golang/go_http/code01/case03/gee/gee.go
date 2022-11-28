/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 00:36:59
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 00:53:49
 * @Description:
 */
package gee

import "net/http"

// 首先定义了类型HandlerFunc，这是提供给框架用户的，用来定义路由映射的处理方法。
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 我们在Engine中，添加了一张路由映射表router，key 由请求方法和静态路由地址构成，
// 例如GET-/、GET-/hello、POST-/hello，这样针对相同的路由，
// 如果请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法。
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// 当用户调用(*Engine).GET()方法时，会将路由和处理方法注册到映射表 router 中
func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// Engine实现的 ServeHTTP 方法的作用就是，解析请求的路径，查找路由映射表，
// 如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND 。
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}

// (*Engine).Run()方法，是 ListenAndServe 的包装。
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
