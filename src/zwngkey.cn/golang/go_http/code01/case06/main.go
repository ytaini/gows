/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 00:56:18
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-30 13:59:15
 */

package main

import (
	"net/http"

	"zwngkey.cn/golang/go_http/code01/case06/gee"
)

func main() {
	r := gee.New()
	r.Get("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, `<h1>Index page</h1>`)
	})
	v1 := r.Group("/v1")
	{
		v1.Get("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, `<h1>Hello gee/h1>`)
		})
		v1.Get("/hello", func(c *gee.Context) {
			// Path /hello?name=zhangsan
			c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.Get("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s,you are at %s\n", c.Param("name"), c.Path)
		})

		v2.Get("/assets/*filepath", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"filepath": c.Param("filepath"),
			})
		})

		v2.Post("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9090")
}
