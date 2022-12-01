/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 18:32:07
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 20:03:03
 */
package gee

import "log"

// Group对象，还需要有访问Router的能力，为了方便，我们可以在Group中，保存一个指针，指向Engine，
// 整个框架的所有资源都是由Engine统一协调的，那么就可以通过Engine间接地访问各种接口了。

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share a Engine instance
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 定义Use函数，将中间件应用到某个 Group 。
func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middleware...)
}

// 那我们就可以将和路由有关的函数，都交给RouterGroup实现了。
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
