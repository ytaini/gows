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
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	log.Println(r.Run())
}
