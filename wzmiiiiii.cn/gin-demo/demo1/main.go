// @author: wzmiiiiii
// @since: 2022/12/25 15:09:03
// @desc: TODO

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 记录到文件。
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "pong",
		//})

		// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON。
		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		//c.AsciiJSON(http.StatusOK, gin.H{
		//	"lang": "go语言",
		//	"tag":  "<br>",
		//})

		// 提供 unicode 实体
		// {"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})

		// 提供字面字符
		// {"html":"<b>Hello, world!</b>"}
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})

		names := []string{"lena", "austin", "foo"}
		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// JSONP stands for JSON with Padding.
	// JSONP 是一种无需担心跨域问题的发送 JSON 数据的方法。
	// JSONP 不使用 XMLHttpRequest 对象。
	// JSONP 使用 <script> 标签代替。

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// /JSONP?callback=x
		// 将输出：x({"foo":"bar"})
		c.JSONP(http.StatusOK, data)
	})

	// 从 reader 读取数据
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		log.Println(err, response)
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	log.Println(r.Run())
}
