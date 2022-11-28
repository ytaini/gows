/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 00:56:18
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 02:24:08
 */

package main

import (
	"net/http"

	"zwngkey.cn/golang/go_http/code01/case04/gee"
)

func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, `<h1>hello gee</h1>`)
	})

	r.Get("/hello", func(c *gee.Context) {
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
