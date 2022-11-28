/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 00:37:17
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 00:50:38
 * @Description:
 */
package main

import (
	"fmt"
	"net/http"

	"zwngkey.cn/golang/go_http/code01/case03/gee"
)

// 使用New()创建 gee 的实例，使用 GET()方法添加路由，最后使用Run()启动Web服务。
// 这里的路由，只是静态路由，不支持/hello/:name这样的动态路由
func main() {
	r := gee.New()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello go")
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello go web")
	})

	r.Run(":9090")
}
