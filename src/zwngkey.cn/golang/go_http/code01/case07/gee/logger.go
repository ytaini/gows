/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 20:10:38
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 20:11:58
 */
package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.R.RequestURI, time.Since(t))
	}
}
