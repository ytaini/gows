// @author: wzmiiiiii
// @since: 2022/12/26 00:21:57
// @desc: TODO

package controller

import (
	"github.com/gin-gonic/gin"
	"wzmiiiiii.cn/gind/demo3/controller/hello"
)

// RegisterRouter 路由设置
func RegisterRouter(r *gin.Engine) {
	new(hello.Controller).Router(r)

	api := apiGroup(r)
	new(MemberController).Router(api)
}

func apiGroup(r *gin.Engine) *gin.RouterGroup {
	return r.Group("/api")
}
