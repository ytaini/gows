/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 01:38:51
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 13:54:45
 */
package gee

import (
	"log"
	"net/http"
)

//  使用 Trie 树实现动态路由(dynamic route)解析。
//  支持两种模式:name和*filepath

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
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
