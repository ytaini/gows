/*
 * @Author: zwngkey
 * @Date: 2022-05-13 07:33:55
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 07:41:43
 * @Description:
 */
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
