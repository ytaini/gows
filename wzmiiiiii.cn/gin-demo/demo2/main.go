// @author: wzmiiiiii
// @since: 2022/12/25 15:09:03
// @desc: TODO

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件。
	r.GET("/benchmark", myBenchLogger(), benchEndpoint)

	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	authorized := r.Group("/")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// 嵌套路由组
		testing := authorized.Group("testing")
		{
			testing.GET("/analytics", analyticsEndpoint)
		}
	}

}

func AuthRequired() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func myBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func analyticsEndpoint(c *gin.Context) {
}

func readEndpoint(c *gin.Context) {

}

func submitEndpoint(c *gin.Context) {

}

func loginEndpoint(c *gin.Context) {

}

func benchEndpoint(c *gin.Context) {

}
