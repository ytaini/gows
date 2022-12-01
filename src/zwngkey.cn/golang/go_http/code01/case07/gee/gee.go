/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 01:38:51
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 21:37:38
 */
package gee

import (
	"net/http"
	"strings"
)

// 设计并实现 Web 框架的中间件(Middlewares)机制。
// 实现通用的Logger中间件，能够记录请求到响应所花费的时间，

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup //store all groups
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// ServeHTTP 函数也有变化，当我们接收到一个具体请求时，要判断该请求适用于哪些中间件，
// 在这里我们简单通过 URL 的前缀来判断。得到中间件列表后，赋值给 c.handlers
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
