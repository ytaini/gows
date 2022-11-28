/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 01:02:18
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 02:18:48
 */
package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 构建JSON数据时，显得更简洁。
type H map[string]any

// Context目前只包含了http.ResponseWriter和*http.Request，另外提供了对 Method 和 Path 这两个常用属性的直接访问。
// 提供了访问Query和PostForm参数的方法。
// 提供了快速构造String/Data/JSON/HTML响应的方法。
type Context struct {
	W http.ResponseWriter
	R *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// 只能获取到 Content-Type: application/x-www-form-urlencoded 和
// Content-Type: multipart/form-data; 的数据
func (c *Context) PostForm(key string) string {
	return c.R.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.W.Header().Set(key, value)
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}

func (c *Context) String(code int, format string, values ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.Data(code, []byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj any) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W)

	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Data(code, []byte(html))
}
