/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 00:56:18
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 21:37:20
 */

package main

import (
	"log"
	"net/http"
	"time"

	"zwngkey.cn/golang/go_http/code01/case07/gee"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.R.RequestURI, time.Since(t))

	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())

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
	v2.Use(onlyForV2())
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
