/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 00:56:18
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 02:24:06
 */

package main

import (
	"net/http"

	"zwngkey.cn/golang/go_http/code01/case04/gee"
)

// Handler的参数变成成了gee.Context，提供了查询Query/PostForm参数的功能。
// gee.Context封装了HTML/String/JSON函数，能够快速构造HTTP响应。
func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, `<h1>hello gee</h1>`)
	})

	r.Get("/hello", func(c *gee.Context) {
		// Path /hello?name=zhangsan
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"), c.Path)
	})

	r.Post("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9090")
}
