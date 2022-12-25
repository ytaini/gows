// @author: wzmiiiiii
// @since: 2022/12/26 00:24:48
// @desc: TODO

package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) Router(r *gin.Engine) {
	r.GET("/hello", c.hello)
}

func (c *Controller) hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
