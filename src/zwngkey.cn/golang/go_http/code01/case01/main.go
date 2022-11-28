/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-28 18:44:28
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 00:20:53
 * @Description:
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.Handle("/", http.HandlerFunc(indexHandler))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// 设置了2个路由，/和/hello，分别绑定 indexHandler 和 helloHandler ，
	// 根据不同的HTTP请求会调用不同的处理函数。
	log.Fatalln(http.ListenAndServe(":9090", nil))
	// 用来启动 Web 服务的，第一个参数是地址，:9090 表示在 本地 9090 端口监听。
	// 第二个参数是一个http.Handler接口:
	// 	The handler is typically nil, in which case the DefaultServeMux is used.
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
